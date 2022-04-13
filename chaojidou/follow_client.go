package chaojidou

import (
	"encoding/json"
	hook "github.com/robotn/gohook"
	"net/http"
	"net/url"
)

func SendEvent(addr string, e *hook.Event) error {
	url1 := "http://" + addr + "/follower"

	urlValues := url.Values{}
	eventBytes, _ := json.Marshal(*e)
	urlValues.Set("event", string(eventBytes))

	_, err := http.PostForm(url1, urlValues)
	return err
}
