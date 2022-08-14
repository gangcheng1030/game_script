package chaojidou

import (
	"encoding/json"
	"errors"
	hook "github.com/robotn/gohook"
	"io/ioutil"
	"net/http"
	"net/url"
)

func SendEvent(addr string, e *hook.Event) error {
	url1 := "http://" + addr + "/follower"

	urlValues := url.Values{}
	eventBytes, _ := json.Marshal(*e)
	urlValues.Set("event", string(eventBytes))

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
