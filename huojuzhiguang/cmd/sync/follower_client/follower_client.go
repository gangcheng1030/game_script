package follower_client

import (
	"encoding/json"
	"errors"
	hook "github.com/robotn/gohook"
	"io/ioutil"
	"net/http"
	"net/url"
)

func SendEvent(addr string, e *hook.Event) error {
	//if e.Kind == hook.MouseDown {
	//	fmt.Println("send mouse down event")
	//} else if e.Kind == hook.MouseUp {
	//	fmt.Println("send mouse up event")
	//}
	//if e.Kind == hook.MouseMove {
	//	fmt.Println("send mouse move event")
	//}

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
