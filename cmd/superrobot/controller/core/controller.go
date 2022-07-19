package core

import (
	"errors"
	"github.com/gangcheng1030/game_script/cmd/superrobot/controller/config"
	"github.com/gangcheng1030/game_script/cmd/superrobot/model"
	"log"
	"strconv"
	"sync"
	"sync/atomic"
)

type Controller interface {
	AddEvent(event *model.Event) error
	DeleteEvent(eventId int64, nodeId, groupId string, eventType model.EventType) error
	GetAllEvents() []model.Event
	PickEvent(group string, id string) (*model.Event, error)
	FinishEvent(event *model.Event) error

	// node
	GetNodeConfig(id string) *model.NodeConfig
	OnlineNode(group string, id string, addr string)
	OfflineNode(group string, id string) error
	UpdateNodeStatus(group string, id string, oldStatus model.NodeStatus, newStatus model.NodeStatus) error
}

func NewController() *controller {
	c := controller{
		events:           make([]*model.Event, 0, 8),
		fixedIdEvents:    make(map[string][]*model.Event, 8),
		fixedGroupEvents: make(map[string][]*model.Event, 8),

		nodeManagerGroup: make(map[string]*NodeManager, 8),
		accountManager:   NewAccountManager(),
		meiriManager:     &MeiriManager{},
	}

	return &c
}

type controller struct {
	events           []*model.Event
	fixedIdEvents    map[string][]*model.Event
	fixedGroupEvents map[string][]*model.Event

	nodeManagerGroup map[string]*NodeManager
	accountManager   *AccountManager
	meiriManager     *MeiriManager

	idGen int64
	sync.RWMutex
}

func (c *controller) GetAllEvents() []model.Event {
	c.RLock()
	defer c.RUnlock()
	events := make([]model.Event, 0, len(c.events))
	for _, tevents := range c.fixedIdEvents {
		for _, event := range tevents {
			events = append(events, *event)
		}
	}
	for _, tevents := range c.fixedGroupEvents {
		for _, event := range tevents {
			events = append(events, *event)
		}
	}
	for _, event := range c.events {
		events = append(events, *event)
	}
	return events
}

func (c *controller) AddEvent(event *model.Event) error {
	c.Lock()
	defer c.Unlock()
	event.Id = atomic.AddInt64(&c.idGen, 1)
	event.Status = model.EVENT_STATUS_INIT
	if len(event.NodeId) > 0 {
		event.Type = model.EVENT_TYPE_FIXEDID
		c.fixedIdEvents[event.NodeId] = append(c.fixedIdEvents[event.NodeId], event)
	} else if len(event.GroupId) > 0 {
		event.Type = model.EVENT_TYPE_FIXEDGROUP
		c.fixedGroupEvents[event.GroupId] = append(c.fixedGroupEvents[event.GroupId], event)
	} else {
		event.Type = model.EVENT_TYPE_COMMON
		c.events = append(c.events, event)
	}

	return nil
}

func (c *controller) DeleteEvent(eventId int64, nodeId, groupId string, eventType model.EventType) error {
	c.Lock()
	defer c.Unlock()

	return c.unsafeDeleteEvent(eventId, nodeId, groupId, eventType)
}

func (c *controller) PickEvent(group string, id string) (*model.Event, error) {
	c.Lock()
	defer c.Unlock()

	nodeManager, ok := c.nodeManagerGroup[group]
	if !ok {
		return nil, errors.New("group " + group + " 不存在")
	}
	nodeManager.Lock()
	defer nodeManager.Unlock()

	node, ok := nodeManager.nodes[id]
	if !ok {
		return nil, errors.New("node " + id + " 不存在")
	}
	if node.Status != model.NODE_STATUS_IDLE {
		return nil, errors.New("node " + id + " 状态不符")
	}

	// fix id event
	if len(c.fixedIdEvents[id]) > 0 {
		log.Printf("start pickFixedIdEvent %s", id)
		events := c.fixedIdEvents[id]
		for i := 0; i < len(events); i++ {
			err := c.checkStatusForPickEvent(events[i], nodeManager)
			if err != nil {
				log.Printf("pickFixedIdEvent %d failed: %v", events[i].Id, err)
				continue
			}

			event := c.supplementEvent(events[i], nodeManager)
			c.updateStatusForPickEvent(events[i], nodeManager)
			return event, nil
		}
	}

	//// fix group event
	//if len(c.fixedGroupEvents[group]) > 0 {
	//
	//}

	return nil, nil
}

func (c *controller) FinishEvent(event *model.Event) error {
	c.Lock()
	defer c.Unlock()

	eventType := event.Type
	nodeId := event.NodeId
	groupId := event.GroupId
	eventId := event.Id
	var idx int
	if eventType == model.EVENT_TYPE_FIXEDID {
		idx = model.FindEventIndex(c.fixedIdEvents[nodeId], eventId)
		if idx == -1 {
			return errors.New("event不存在")
		}
		if c.fixedIdEvents[nodeId][idx].Status != model.EVENT_STATUS_PROCESSING {
			return errors.New("event " + strconv.FormatInt(eventId, 10) + " 状态不符")
		}
		c.fixedIdEvents[nodeId] = append(c.fixedIdEvents[nodeId][:idx], c.fixedIdEvents[nodeId][(idx+1):]...)
	} else if eventType == model.EVENT_TYPE_FIXEDGROUP {
		idx = model.FindEventIndex(c.fixedGroupEvents[groupId], eventId)
		if idx == -1 {
			return errors.New("event不存在")
		}
		if c.fixedGroupEvents[groupId][idx].Status != model.EVENT_STATUS_PROCESSING {
			return errors.New("event " + strconv.FormatInt(eventId, 10) + " 状态不符")
		}
		c.fixedGroupEvents[groupId] = append(c.fixedGroupEvents[groupId][:idx], c.fixedGroupEvents[groupId][(idx+1):]...)
	} else {
		idx = model.FindEventIndex(c.events, eventId)
		if idx == -1 {
			return errors.New("event不存在")
		}
		if c.events[idx].Status != model.EVENT_STATUS_PROCESSING {
			return errors.New("event " + strconv.FormatInt(eventId, 10) + " 状态不符")
		}
		c.events = append(c.events[:idx], c.events[(idx+1):]...)
	}

	nodeManager, ok := c.nodeManagerGroup[groupId]
	if !ok {
		return errors.New("group " + groupId + " 不存在")
	}
	nodeManager.Lock()
	defer nodeManager.Unlock()

	node, ok := nodeManager.nodes[nodeId]
	if ok && node.Status == model.NODE_STATUS_BUSY {
		node.Status = model.NODE_STATUS_IDLE
	}

	for _, tmpNodeId := range event.FollowerNodeIds {
		followerNode, ok := nodeManager.nodes[tmpNodeId]
		if ok && followerNode.Status == model.NODE_STATUS_BUSY {
			followerNode.Status = model.NODE_STATUS_IDLE
		}
	}

	account := event.Account
	leaderAccount, ok := c.accountManager.accounts[account.AccountName]
	if !ok {
		return errors.New("account " + account.AccountName + " 不存在")
	}
	if leaderAccount.Status != model.ACCOUNT_STATUS_BUSY {
		return errors.New("account " + account.AccountName + " 状态不符")
	}
	leaderAccount.Status = model.ACCOUNT_STATUS_IDLE

	for _, accountName := range account.FollowerAccountNames {
		followerAccount, ok := c.accountManager.accounts[accountName]
		if !ok {
			return errors.New("account " + accountName + " 不存在")
		}
		if followerAccount.Status != model.ACCOUNT_STATUS_BUSY {
			return errors.New("account " + accountName + " 状态不符")
		}
		followerAccount.Status = model.ACCOUNT_STATUS_IDLE
	}

	return nil
}

func (c *controller) GetNodeConfig(id string) *model.NodeConfig {
	cfg, ok := config.NodeConfigCache[id]
	if ok {
		return cfg
	} else {
		return model.NewDefaultNodeConfig()
	}
}

func (c *controller) OnlineNode(group string, id string, addr string) {
	c.Lock()
	defer c.Unlock()

	nodeManager, ok := c.nodeManagerGroup[group]
	if !ok {
		nodeManager = NewNodeManager()
		c.nodeManagerGroup[group] = nodeManager
	}
	nodeManager.Lock()
	defer nodeManager.Unlock()

	nodeManager.unsafeOnlineNode(id, addr)
}

func (c *controller) OfflineNode(group string, id string) error {
	c.Lock()
	defer c.Unlock()

	nodeManager, ok := c.nodeManagerGroup[group]
	if !ok {
		return errors.New("group " + group + " 不存在")
	}

	nodeManager.Lock()
	defer nodeManager.Unlock()

	_, ok = nodeManager.nodes[id]
	if !ok {
		return errors.New("node " + id + " 不存在")
	}

	nodeManager.unsafeOfflineNode(id)
	return nil
}

func (c *controller) UpdateNodeStatus(group string, id string, oldStatus model.NodeStatus, newStatus model.NodeStatus) error {
	c.Lock()
	defer c.Unlock()

	nodeManager, ok := c.nodeManagerGroup[group]
	if !ok {
		return errors.New("group " + group + " 不存在")
	}

	nodeManager.Lock()
	defer nodeManager.Unlock()

	node, ok := nodeManager.nodes[id]
	if !ok {
		return errors.New("node " + id + " 不存在")
	}
	if node.Status != oldStatus {
		return errors.New("node " + id + " 状态不符")
	}

	node.Status = newStatus
	return nil
}

func (c *controller) checkStatusForPickEvent(event *model.Event, nm *NodeManager) error {
	if event.Status == model.EVENT_STATUS_PROCESSING {
		return errors.New("event " + strconv.FormatInt(event.Id, 10) + " 状态不符")
	}

	leaderNode, ok := nm.nodes[event.NodeId]
	if !ok {
		return errors.New("node " + event.NodeId + " 不存在")
	}
	if leaderNode.Status != model.NODE_STATUS_IDLE {
		return errors.New("node " + event.NodeId + " 状态不符")
	}

	for _, nodeId := range event.FollowerNodeIds {
		node, ok := nm.nodes[nodeId]
		if !ok {
			return errors.New("node " + nodeId + " 不存在")
		}
		if node.Status != model.NODE_STATUS_IDLE {
			return errors.New("node " + nodeId + " 状态不符")
		}
	}

	account := event.Account
	leaderAccount, ok := c.accountManager.accounts[account.AccountName]
	if !ok {
		return errors.New("account " + account.AccountName + " 不存在")
	}
	if leaderAccount.Status != model.ACCOUNT_STATUS_IDLE {
		return errors.New("account " + account.AccountName + " 状态不符")
	}

	for _, accountName := range account.FollowerAccountNames {
		followerAccount, ok := c.accountManager.accounts[accountName]
		if !ok {
			return errors.New("account " + accountName + " 不存在")
		}
		if followerAccount.Status != model.ACCOUNT_STATUS_IDLE {
			return errors.New("account " + accountName + " 状态不符")
		}
	}

	return nil
}

// 假定node一定存在
func (c *controller) updateStatusForPickEvent(event *model.Event, nm *NodeManager) {
	event.Status = model.EVENT_STATUS_PROCESSING

	leaderNode := nm.nodes[event.NodeId]
	leaderNode.Status = model.NODE_STATUS_BUSY

	for _, nodeId := range event.FollowerNodeIds {
		node := nm.nodes[nodeId]
		node.Status = model.NODE_STATUS_BUSY
	}

	account := event.Account
	leaderAccount := c.accountManager.accounts[account.AccountName]
	leaderAccount.Status = model.NODE_STATUS_BUSY

	for _, accountName := range account.FollowerAccountNames {
		followerAccount := c.accountManager.accounts[accountName]
		followerAccount.Status = model.NODE_STATUS_BUSY
	}
}

func (c *controller) supplementEvent(event *model.Event, nm *NodeManager) *model.Event {
	acount := &event.Account
	acount.Password = config.AccountCache[acount.AccountName].Password
	acount.FollowerAddrs = acount.FollowerAddrs[:0]
	acount.FollowerPasswords = acount.FollowerPasswords[:0]
	for i, followerNodeId := range event.FollowerNodeIds {
		acount.FollowerAddrs = append(acount.FollowerAddrs, nm.nodes[followerNodeId].Addr)
		acount.FollowerPasswords = append(acount.FollowerPasswords, config.AccountCache[acount.FollowerAccountNames[i]].Password)
	}

	for i := range acount.Roles {
		role := &acount.Roles[i]
		role.FollowerJunTuanNames = role.FollowerJunTuanNames[:0]
		for _, followAccountName := range acount.FollowerAccountNames {
			role.FollowerJunTuanNames = append(role.FollowerJunTuanNames, config.AccountCache[followAccountName].JunTuanName)
		}
	}

	return event
}

func (c *controller) unsafeDeleteEvent(eventId int64, nodeId, groupId string, eventType model.EventType) error {
	var idx int
	if eventType == model.EVENT_TYPE_FIXEDID {
		idx = model.FindEventIndex(c.fixedIdEvents[nodeId], eventId)
		if idx == -1 {
			return errors.New("event不存在")
		}
		if c.fixedIdEvents[nodeId][idx].Status != model.EVENT_STATUS_INIT {
			return errors.New("event " + strconv.FormatInt(eventId, 10) + " 状态不符")
		}
		c.fixedIdEvents[nodeId] = append(c.fixedIdEvents[nodeId][:idx], c.fixedIdEvents[nodeId][(idx+1):]...)
		return nil
	} else if eventType == model.EVENT_TYPE_FIXEDGROUP {
		idx = model.FindEventIndex(c.fixedGroupEvents[groupId], eventId)
		if idx == -1 {
			return errors.New("event不存在")
		}
		if c.fixedGroupEvents[groupId][idx].Status != model.EVENT_STATUS_INIT {
			return errors.New("event " + strconv.FormatInt(eventId, 10) + " 状态不符")
		}
		c.fixedGroupEvents[groupId] = append(c.fixedGroupEvents[groupId][:idx], c.fixedGroupEvents[groupId][(idx+1):]...)
		return nil
	} else {
		idx = model.FindEventIndex(c.events, eventId)
		if idx == -1 {
			return errors.New("event不存在")
		}
		if c.events[idx].Status != model.EVENT_STATUS_INIT {
			return errors.New("event " + strconv.FormatInt(eventId, 10) + " 状态不符")
		}

		c.events = append(c.events[:idx], c.events[(idx+1):]...)
		return nil
	}
}
