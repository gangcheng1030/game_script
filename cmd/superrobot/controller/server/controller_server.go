package server

import (
	"encoding/json"
	"github.com/gangcheng1030/game_script/cmd/superrobot/controller/core"
	"github.com/gangcheng1030/game_script/cmd/superrobot/model"
	"log"
	"net/http"
)

type ControllerHandler struct {
	Controller core.Controller
}

func (ch *ControllerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	eventType := r.Form.Get("eventType")
	if eventType == "addEvent" {
		data := r.Form.Get("data")
		event := model.Event{}
		json.Unmarshal([]byte(data), &event)
		err := ch.Controller.AddEvent(&event)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if eventType == "getAllEvents" {
		events := ch.Controller.GetAllEvents()
		if len(events) > 0 {
			resp, _ := json.Marshal(events)
			w.Write(resp)
		}
	} else if eventType == "deleteEvent" {
		data := r.Form.Get("data")
		event := model.Event{}
		json.Unmarshal([]byte(data), &event)
		err := ch.Controller.DeleteEvent(event.Id, event.NodeId, event.GroupId, event.Type)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if eventType == "onlineNode" {
		data := r.Form.Get("data")
		node := model.Node{}
		json.Unmarshal([]byte(data), &node)
		ch.Controller.OnlineNode(node.Config.Group, node.Id, node.Addr)
	} else if eventType == "offlineNode" {
		data := r.Form.Get("data")
		node := model.Node{}
		json.Unmarshal([]byte(data), &node)
		err := ch.Controller.OfflineNode(node.Config.Group, node.Id)
		if err != nil {
			log.Printf("offline node %s err: %v", node.Id, err)
		}
	} else if eventType == "pickEvent" {
		data := r.Form.Get("data")
		node := model.Node{}
		json.Unmarshal([]byte(data), &node)
		event, err := ch.Controller.PickEvent(node.Config.Group, node.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if event != nil {
			resp, _ := json.Marshal(event)
			w.Write(resp)
		}
	} else if eventType == "finishEvent" {
		data := r.Form.Get("data")
		event := model.Event{}
		json.Unmarshal([]byte(data), &event)
		err := ch.Controller.FinishEvent(&event)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if eventType == "updateNodeStatus" {
		data := r.Form.Get("data")
		req := model.UpdateNodeStatusReq{}
		json.Unmarshal([]byte(data), &req)
		err := ch.Controller.UpdateNodeStatus(req.Group, req.Id, req.OldStatus, req.NewStatus)
		if err != nil {
			log.Printf("updateNodeStatus err, %s", data)
		}
	} else if eventType == "getNodeConfig" {
		data := r.Form.Get("data")
		node := model.Node{}
		json.Unmarshal([]byte(data), &node)
		cfg := ch.Controller.GetNodeConfig(node.Id)
		resp, _ := json.Marshal(cfg)
		w.Write(resp)
	}

	w.WriteHeader(http.StatusOK)
}
