package client

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func SendEvent(addr string, eventType string, e interface{}) error {
	url1 := "http://" + addr + "/follower/autorobot"

	urlValues := url.Values{}
	eventBytes, _ := json.Marshal(e)
	urlValues.Set("eventType", eventType)
	urlValues.Set("data", string(eventBytes))

	_, err := http.PostForm(url1, urlValues)
	return err
}
