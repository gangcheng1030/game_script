package core

import (
	"encoding/json"
	"github.com/gangcheng1030/game_script/cmd/superrobot/model"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var event1 = model.Event{
	MeiRi:   "shzk",
	NodeId:  "node2",
	GroupId: "1",
	Account: model.Account{
		AccountName: "gangcheng11223@163.com",
		Roles: []model.Role{
			{
				Id: 6,
			},
		},
	},
}

var event2 = model.Event{
	MeiRi:           "shzk",
	NodeId:          "node2",
	FollowerNodeIds: []string{"node1", "node3"},
	GroupId:         "1",
	Account: model.Account{
		AccountName:          "gangcheng11223@163.com",
		FollowerAccountNames: []string{"gangcheng11224@163.com", "gangcheng11225@163.com"},
		Roles: []model.Role{
			{
				Id:              6,
				FollowerRoleIds: []int{1, 1},
			},
		},
	},
}

func TestController_Event(t *testing.T) {
	ctrl := NewController()
	ctrl.AddEvent(&event1)
	ctrl.AddEvent(&event2)
	//allEvents := ctrl.GetAllEvents()
	//log.Println(allEvents)
	ctrl.OnlineNode("1", "node2", "192.168.1.9:6688")
	ctrl.OnlineNode("1", "node1", "192.168.1.6:6688")
	ctrl.OnlineNode("1", "node3", "192.168.1.8:6688")

	var resEvent *model.Event
	var err error
	var resEventStr []byte
	resEvent, err = ctrl.PickEvent("1", "node2")
	assert.NoError(t, err, "pick event err")
	resEventStr, _ = json.Marshal(resEvent)
	log.Println(string(resEventStr))

	ctrl.UpdateNodeStatus("1", "node2", model.NODE_STATUS_BUSY, model.NODE_STATUS_IDLE)

	resEvent, err = ctrl.PickEvent("1", "node2")
	assert.NoError(t, err, "pick event err")
	assert.NotNil(t, resEvent)
	resEventStr, _ = json.Marshal(resEvent)
	log.Println(string(resEventStr))
}
