package follower_server

import (
	"encoding/json"
	"github.com/gangcheng1030/game_script/huojuzhiguang/cmd/sync/synccore"
	"github.com/gangcheng1030/game_script/utils"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"log"
	"net/http"
)

type FollowerHandler struct {
}

func (fh *FollowerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	eventStr := r.Form.Get("event")
	event := hook.Event{}
	json.Unmarshal([]byte(eventStr), &event)
	//fmt.Printf("%d -- %s", event.Kind, event.String())

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
		k, ok := utils.Raw2key[e.Rawcode]
		if !ok {
			log.Printf("wrong rawCode: %d", e.Rawcode)
			return nil
		}
		robotgo.KeyPress(k)
		if e.Rawcode == 83 { // s
			synccore.HandleS()
		}
	//case hook.KeyUp:
	//	k, ok := utils.Raw2key[e.Rawcode]
	//	if !ok {
	//		log.Printf("wrong rawCode: %d", e.Rawcode)
	//		return nil
	//	}
	//	robotgo.KeyUp(k)
	case hook.MouseDown:
		if e.Button == 2 {
			robotgo.Click("right")
		} else {
			//fmt.Println("receive mouse down event")
			robotgo.Click()
		}
	//case hook.MouseUp:
	//	if e.Button == 2 {
	//		robotgo.MouseUp("right")
	//	} else {
	//		robotgo.MouseUp()
	//	}
	case hook.MouseMove:
		//fmt.Println("receive mouse move event ")
		robotgo.Move(int(e.X), int(e.Y))
	default:
	}

	return nil
}
