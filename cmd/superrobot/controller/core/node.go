package core

import (
	"github.com/gangcheng1030/game_script/cmd/superrobot/controller/config"
	"github.com/gangcheng1030/game_script/cmd/superrobot/model"
	"log"
	"sync"
	"time"
)

type NodeManager struct {
	nodes map[string]*model.Node

	sync.RWMutex
}

func NewNodeManager() *NodeManager {
	m := NodeManager{
		nodes: make(map[string]*model.Node, 8),
	}
	return &m
}

func (m *NodeManager) unsafeOnlineNode(id string, addr string) {
	if _, ok := m.nodes[id]; ok {
		log.Printf("node %s has exists.", id)
		return
	}

	node := model.Node{
		Id:              id,
		Addr:            addr,
		Status:          model.NODE_STATUS_IDLE,
		LastUpdatedTime: time.Now(),
		Config:          *config.NodeConfigCache[id],
	}
	m.nodes[id] = &node
}

func (m *NodeManager) unsafeOfflineNode(id string) {
	delete(m.nodes, id)
}

func (m *NodeManager) unsafeHeartbeat(id string) {
	node, ok := m.nodes[id]
	if !ok {
		return
	}
	node.LastUpdatedTime = time.Now()
}
