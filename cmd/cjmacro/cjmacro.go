package main

import (
	"fmt"
	"github.com/gangcheng1030/game_script/chaojidou"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

var sogouChaoJiDou chaojidou.ChaoJiDou
var qqChaoJiDou chaojidou.ChaoJiDou
var t360ChaoJiDou chaojidou.ChaoJiDou
var chromeChaoJiDou chaojidou.ChaoJiDou

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
	qqChaoJiDou = chaojidou.NewQqChaoJiDou()
	t360ChaoJiDou = chaojidou.NewT360ChaoJiDou()
	chromeChaoJiDou = chaojidou.NewChromeChaoJiDou()

	fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
	hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("ctrl-shift-q")
		hook.End()
	})

	fmt.Println("--- Please press shift + f to confirm ---")
	hook.Register(hook.KeyDown, []string{robotgo.Shift, robotgo.KeyF}, func(e hook.Event) {
		fmt.Println("shift-f")
		currentPid := robotgo.GetPID()
		sogouChaoJiDou.Empty()
		qqChaoJiDou.Confirm()
		t360ChaoJiDou.Confirm()
		sogouChaoJiDou.Confirm()
		//chromeChaoJiDou.Confirm()
		robotgo.ActivePID(currentPid)
		robotgo.MilliSleep(chaojidou.WAITING_ACTIVE_PID_MILLI_SECONDS)
	})

	fmt.Println("--- Please press shift + q to quit ---")
	hook.Register(hook.KeyDown, []string{robotgo.Shift, robotgo.KeyQ}, func(e hook.Event) {
		fmt.Println("shift-q")
		currentPid := robotgo.GetPID()
		sogouChaoJiDou.Empty()
		qqChaoJiDou.Esc()
		t360ChaoJiDou.Esc()
		sogouChaoJiDou.Esc()
		//chromeChaoJiDou.Esc()
		robotgo.ActivePID(currentPid)
		robotgo.MilliSleep(chaojidou.WAITING_ACTIVE_PID_MILLI_SECONDS)
	})

	fmt.Println("--- Please press shift + g to accept group ---")
	hook.Register(hook.KeyDown, []string{robotgo.Shift, robotgo.KeyG}, func(e hook.Event) {
		fmt.Println("shift-g")
		currentPid := robotgo.GetPID()
		sogouChaoJiDou.Empty()
		qqChaoJiDou.GroupAccept()
		t360ChaoJiDou.GroupAccept()
		sogouChaoJiDou.GroupAccept()
		//chromeChaoJiDou.GroupAccept()
		robotgo.ActivePID(currentPid)
		robotgo.MilliSleep(chaojidou.WAITING_ACTIVE_PID_MILLI_SECONDS)
	})

	fmt.Println("--- Please press shift + e to enter instance zones ---")
	hook.Register(hook.KeyDown, []string{robotgo.Shift, robotgo.KeyE}, func(e hook.Event) {
		fmt.Println("shift-e")
		currentPid := robotgo.GetPID()
		sogouChaoJiDou.Empty()
		qqChaoJiDou.EnterAccept()
		t360ChaoJiDou.EnterAccept()
		sogouChaoJiDou.EnterAccept()
		//chromeChaoJiDou.EnterAccept()
		robotgo.ActivePID(currentPid)
		robotgo.MilliSleep(chaojidou.WAITING_ACTIVE_PID_MILLI_SECONDS)
	})

	fmt.Println("--- Please press shift + w to confirm or enter ---")
	hook.Register(hook.KeyDown, []string{robotgo.Shift, robotgo.KeyW}, func(e hook.Event) {
		fmt.Println("shift-w")
		currentPid := robotgo.GetPID()
		sogouChaoJiDou.Empty()
		qqChaoJiDou.ConfirmOrAccept()
		t360ChaoJiDou.ConfirmOrAccept()
		sogouChaoJiDou.ConfirmOrAccept()
		//chromeChaoJiDou.ConfirmOrAccept()
		robotgo.ActivePID(currentPid)
		robotgo.MilliSleep(chaojidou.WAITING_ACTIVE_PID_MILLI_SECONDS)
	})

	s := hook.Start()
	<-hook.Process(s)
}
