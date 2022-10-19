package main

import (
	"fmt"
	"github.com/gangcheng1030/game_script/huojuzhiguang/common"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"time"
)

func initComponent() {
	robotgo.KeySleep = 100
	robotgo.MouseSleep = 100
}

func main() {
	fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
	hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("ctrl-shift-q")
		hook.End()
	})

	fmt.Println("--- Please press shift + 1 to 'bfhy - xyzq' ---")
	var endTime time.Time = time.Now()
	hook.Register(hook.KeyDown, []string{robotgo.Shift, robotgo.Key1}, func(e hook.Event) {
		fmt.Println("shift-1")
		if endTime.Add(time.Second).After(time.Now()) {
			return
		}

		common.Xinyangzhiqiang(1)
		fmt.Println("shift-1 end")
		endTime = time.Now()
	})

	s := hook.Start()
	<-hook.Process(s)
}
