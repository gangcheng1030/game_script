package chaojidou

import (
	"encoding/json"
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
	fmt.Println(event)

	fh.handleEvent(&event)

	w.WriteHeader(http.StatusOK)
}

func (fh *FollowerHandler) handleEvent(e *hook.Event) {
	switch e.Kind {
	case hook.KeyDown:
		robotgo.KeyPress(string(e.Keychar))
	case hook.MouseDown:
		robotgo.MoveSmooth(int(e.X), int(e.Y), 0.9, 0.9)
		robotgo.MilliSleep(300)
		robotgo.Click()
		for i := 1; i < int(e.Clicks); i++ {
			robotgo.Sleep(2)
			robotgo.Click()
		}
	case hook.MouseMove:
		robotgo.MoveSmooth(int(e.X), int(e.Y), 0.9, 0.9)
		robotgo.MilliSleep(300)
		robotgo.Click("right")
	default:
	}
}
