package chaojidou

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"net/http"
)

type FollowerHandler struct {
}

func (fh *FollowerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	eventStr := r.Form.Get("event")
	event := hook.Event{}
	json.Unmarshal([]byte(eventStr), &event)
	fmt.Printf("%d -- %s", event.Kind, event.String())

	err := fh.handleEvent(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (fh *FollowerHandler) handleEvent(e *hook.Event) error {
	switch e.Kind {
	case hook.KeyDown:
		robotgo.KeyPress(string(e.Keychar))
	case hook.MouseDown:
		robotgo.MoveSmooth(int(e.X), int(e.Y), 0.9, 0.9)
		robotgo.MilliSleep(300)
		if e.Button == 2 {
			robotgo.Click("right")
		} else {
			robotgo.Click()
		}
		for i := 1; i < int(e.Clicks); i++ {
			robotgo.Sleep(2)
			robotgo.Click()
		}
	case 101: // 老爹
		DefaultCaptain.LaoDie()
	case 102: // isOnline
		isOnline := DefaultCaptain.IsOnline(false)
		if !isOnline {
			return errors.New("i am offline")
		}
	case 103: // forceQuit
		go DefaultCaptain.ForceQuit(false)
	default:
	}

	return nil
}
