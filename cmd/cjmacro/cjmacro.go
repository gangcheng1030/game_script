package main

import (
	"fmt"
	"github.com/gangcheng1030/game_script/chaojidou"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

var sogouChaoJiDou chaojidou.ChaoJiDou
var qqChaoJiDow chaojidou.ChaoJiDou

/**
	cjdmj用的宏脚本，当队长发起进本、开始战斗、回城时，其它队员可以自动做出响应。
*/
func main() {
	add()
}

func add() {
	robotgo.KeySleep = 100
	robotgo.MouseSleep = 100

	sogouChaoJiDou = chaojidou.NewSogouChaoJiDou()
	qqChaoJiDow = chaojidou.NewQqChaoJiDou()

	fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
	hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("ctrl-shift-q")
		hook.End()
	})

	fmt.Println("--- Please press shift + f to confirm ---")
	hook.Register(hook.KeyDown, []string{robotgo.Shift, robotgo.KeyF}, func(e hook.Event) {
		fmt.Println("shift-f")
		currentPid := robotgo.GetPID()
		sogouChaoJiDou.Confirm()
		qqChaoJiDow.Confirm()
		robotgo.ActivePID(currentPid)
		robotgo.MilliSleep(chaojidou.WAITING_ACTIVE_PID_MILLI_SECONDS)
	})

	s := hook.Start()
	<-hook.Process(s)
}
