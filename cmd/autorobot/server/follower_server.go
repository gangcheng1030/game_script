package server

import (
	"encoding/json"
	"github.com/gangcheng1030/game_script/chaojidou"
	"github.com/gangcheng1030/game_script/utils/robotgoutil"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"net/http"
	"os/exec"
)

type FollowerHandler struct {
	CjdDir          string
	StartGameButton robotgo.Rect
	Captain         chaojidou.ChaoJiDou
	FullScreenMode  int
	IsRunning       bool
}

func (fh *FollowerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	eventType := r.Form.Get("eventType")
	if eventType == "signin" {
		fh.IsRunning = true
		data := r.Form.Get("data")
		account := Account{}
		json.Unmarshal([]byte(data), &account)
		go fh.signIn(account)
	} else if eventType == "select_role" {
		data := r.Form.Get("data")
		role := Role{}
		json.Unmarshal([]byte(data), &role)
		go fh.selectRole(role)
	} else if eventType == "quit" {
		data := r.Form.Get("data")
		role := Role{}
		json.Unmarshal([]byte(data), &role)
		go fh.quit(role)
	} else if eventType == "hook_event" {
		data := r.Form.Get("data")
		event := hook.Event{}
		json.Unmarshal([]byte(data), &event)
		fh.handleEvent(&event)
	}

	w.WriteHeader(http.StatusOK)
}

func (fh *FollowerHandler) signIn(account Account) {
	cjd := exec.Command(fh.CjdDir + "\\Launcher.exe")
	err := cjd.Run()
	if err != nil {
		panic(err)
	}
	robotgo.Sleep(45)

	// 开始游戏
	robotgoutil.ClickButton(fh.StartGameButton, 100)

	fh.Captain, err = chaojidou.Build(chaojidou.CLIENT_TYPE_OFFICIAL)
	if err != nil {
		panic(err)
	}

	// 将窗口置于前台
	robotgo.MoveSmooth(946, 190, 0.9, 0.9)
	robotgo.Click()
	robotgo.Sleep(3)

	// 登录
	fh.Captain.SignIn(account.AccountName, account.Password)
}

func (fh *FollowerHandler) selectRole(role Role) {
	fh.Captain.SelectRole(role.Id-1, role.First, fh.FullScreenMode)

	fh.Captain.RepairEquipment()
	chaojidou.NpcWaitSecs = 20
	if !role.DisablePreClearBag {
		fh.Captain.ClearBag(true)
	}
	chaojidou.NpcWaitSecs = 30
	fh.Captain.CardsUp()
}

func (fh *FollowerHandler) quit(role Role) {
	fh.Captain.RepairEquipment()
	if role.PostClearBag {
		chaojidou.NpcWaitSecs = 8
		fh.Captain.ClearBag(false)
		chaojidou.NpcWaitSecs = 30
	}
	fh.Captain.CardsDown()

	if !role.Last {
		fh.Captain.QuitRole()
	} else {
		fh.Captain.Quit()
		fh.IsRunning = false
	}
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

type Account struct {
	AccountName string
	Password    string
}

type Role struct {
	Id    int
	First bool
	Last  bool

	DisablePreClearBag bool
	PostClearBag       bool
}
