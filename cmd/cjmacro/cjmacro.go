package main

import (
	"fmt"
	"github.com/gangcheng1030/game_script/utils"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

/**
cjdmj用的宏脚本，当队长发起进本、开始战斗、回城时，其它队员可以自动做出响应。
*/
func main() {
	add()
}

func add() {
	robotgo.KeySleep = 100
	robotgo.MouseSleep = 100

	fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
	hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("ctrl-shift-q")
		hook.End()
	})

	fmt.Println("--- Please press shift + f to confirm ---")
	hook.Register(hook.KeyDown, []string{robotgo.Shift, robotgo.KeyF}, func(e hook.Event) {
		fmt.Println("shift-f")

		// sogou浏览器
		sogouPs, _ := utils.FindSogouProcess()
		fmt.Printf("sogouPid: %d\n", sogouPs.Pid)
		err := robotgo.ActivePID(sogouPs.Pid)
		if err != nil {
			fmt.Println(err)
		}
		robotgo.MilliSleep(1000)
		fmt.Printf("activePid: %d\n", robotgo.GetPID())
		//robotgo.Click(robotgo.Right)
		robotgo.KeyToggle(robotgo.KeyF, int(sogouPs.Pid))
		robotgo.MilliSleep(100)
		robotgo.KeyToggle(robotgo.KeyF, int(sogouPs.Pid), robotgo.Up)
		robotgo.MilliSleep(100)

		// qq浏览器
		qqPs, _ := utils.FindQQProcess()
		fmt.Printf("qqPid: %d\n", qqPs.Pid)
		err = robotgo.ActivePID(qqPs.Pid)
		if err != nil {
			fmt.Println(err)
		}
		robotgo.MilliSleep(1000)
		fmt.Printf("activePid: %d\n", robotgo.GetPID())
		//robotgo.Click(robotgo.Right)
		robotgo.KeyToggle(robotgo.KeyF, int(qqPs.Pid))
		robotgo.MilliSleep(100)
		robotgo.KeyToggle(robotgo.KeyF, int(qqPs.Pid), robotgo.Up)
		robotgo.MilliSleep(100)
	})

	s := hook.Start()
	<-hook.Process(s)
}
