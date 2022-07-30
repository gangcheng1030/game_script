package client

import (
	"encoding/json"
	"errors"
	"github.com/gangcheng1030/game_script/cmd/superrobot/model"
	"github.com/gangcheng1030/game_script/utils/iputil"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func GetNodeConfig(addr string, id string) (*model.NodeConfig, error) {
	url1 := "http://" + addr + "/controller"
	urlValues := url.Values{}
	node := model.Node{
		Id: id,
	}
	dataBytes, _ := json.Marshal(node)
	urlValues.Set("eventType", "getNodeConfig")
	urlValues.Set("data", string(dataBytes))

	resp, err := http.PostForm(url1, urlValues)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Printf("getNodeConfig return: %s", string(body))

	if resp.StatusCode != 200 {
		return nil, errors.New(string(body))
	}

	cfg := model.NodeConfig{}
	err = json.Unmarshal(body, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func OnlineNode(addr string, id string, port string, config *model.NodeConfig) error {
	url1 := "http://" + addr + "/controller"
	urlValues := url.Values{}
	localIp, err := iputil.GetLocalIP()
	if err != nil {
		return err
	}
	node := model.Node{
		Id:     id,
		Addr:   localIp + ":" + port,
		Config: *config,
	}
	dataBytes, _ := json.Marshal(node)
	urlValues.Set("eventType", "onlineNode")
	urlValues.Set("data", string(dataBytes))

	resp, err := http.PostForm(url1, urlValues)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New(string(body))
	}

	return nil
}

func OfflineNode(addr string, id string, config *model.NodeConfig) error {
	url1 := "http://" + addr + "/controller"
	urlValues := url.Values{}
	node := model.Node{
		Id:     id,
		Config: *config,
	}
	dataBytes, _ := json.Marshal(node)
	urlValues.Set("eventType", "offlineNode")
	urlValues.Set("data", string(dataBytes))

	resp, err := http.PostForm(url1, urlValues)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New(string(body))
	}

	return nil
}

func PickEvent(addr string, id string, config *model.NodeConfig) (*model.Event, error) {
	url1 := "http://" + addr + "/controller"
	urlValues := url.Values{}
	node := model.Node{
		Id:     id,
		Config: *config,
	}
	dataBytes, _ := json.Marshal(node)
	urlValues.Set("eventType", "pickEvent")
	urlValues.Set("data", string(dataBytes))

	resp, err := http.PostForm(url1, urlValues)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(string(body))
	}

	if len(body) > 0 {
		log.Printf("receive event: %s", string(body))
		event := model.Event{}
		err = json.Unmarshal(body, &event)
		if err != nil {
			return nil, err
		}

		return &event, nil
	}

	return nil, nil
}

func FinishEvent(addr string, event *model.Event) error {
	url1 := "http://" + addr + "/controller"
	urlValues := url.Values{}
	dataBytes, _ := json.Marshal(*event)
	urlValues.Set("eventType", "finishEvent")
	urlValues.Set("data", string(dataBytes))

	resp, err := http.PostForm(url1, urlValues)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(string(body))
	}

	return nil
}

func AddEvent(addr string, event *model.Event) error {
	url1 := "http://" + addr + "/controller"
	urlValues := url.Values{}
	dataBytes, _ := json.Marshal(*event)
	urlValues.Set("eventType", "addEvent")
	urlValues.Set("data", string(dataBytes))

	resp, err := http.PostForm(url1, urlValues)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(string(body))
	}

	log.Printf("add event success. event: %s\n", string(dataBytes))
	return nil
}

func DeleteEvent(addr string, eventId int64, nodeId string, groupId string, eventType model.EventType) error {
	url1 := "http://" + addr + "/controller"
	urlValues := url.Values{}
	event := model.Event{
		Id:      eventId,
		NodeId:  nodeId,
		GroupId: groupId,
		Type:    eventType,
	}
	dataBytes, _ := json.Marshal(event)
	urlValues.Set("eventType", "deleteEvent")
	urlValues.Set("data", string(dataBytes))

	resp, err := http.PostForm(url1, urlValues)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(string(body))
	}

	return nil
}
