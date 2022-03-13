package main

/**
cjdmj用的宏脚本
*/

import (
	"flag"
	"fmt"
	"github.com/gangcheng1030/game_script/chaojidou"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"strings"
)

var captainStr = flag.String("c", "official_client", "client of captain")
var membersStr = flag.String("m", "sogou,qq,360", "clients of members")

var captain chaojidou.ChaoJiDou
var members []chaojidou.ChaoJiDou

func initComponent() {
	flag.Parse()
	var err error
	captain, err = chaojidou.Build(chaojidou.ClientType(*captainStr))
	if err != nil {
		panic(err)
	}

	members = make([]chaojidou.ChaoJiDou, 0)
	if len(*membersStr) > 0 {
		membersSplit := strings.Split(*membersStr, ",")
		for _, memberStr := range membersSplit {
			member, err := chaojidou.Build(chaojidou.ClientType(memberStr))
			if err != nil {
				panic(err)
			}
			members = append(members, member)
		}
	}
}

func main() {
	initComponent()
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
		if len(members) > 0 {
			currentPid := robotgo.GetPID()
			members[0].Empty()
			for i := range members {
				members[(i+1)%len(members)].Confirm()
			}
			robotgo.ActivePID(currentPid)
			robotgo.MilliSleep(chaojidou.WAITING_ACTIVE_PID_MILLI_SECONDS)
		}
	})

	fmt.Println("--- Please press shift + q to quit ---")
	hook.Register(hook.KeyDown, []string{robotgo.Shift, robotgo.KeyQ}, func(e hook.Event) {
		fmt.Println("shift-q")
		if len(members) > 0 {
			currentPid := robotgo.GetPID()
			members[0].Empty()
			for i := range members {
				members[(i+1)%len(members)].Esc()
			}
			robotgo.ActivePID(currentPid)
			robotgo.MilliSleep(chaojidou.WAITING_ACTIVE_PID_MILLI_SECONDS)
		}
	})

	fmt.Println("--- Please press shift + g to accept group ---")
	hook.Register(hook.KeyDown, []string{robotgo.Shift, robotgo.KeyG}, func(e hook.Event) {
		fmt.Println("shift-g")
		if len(members) > 0 {
			currentPid := robotgo.GetPID()
			members[0].Empty()
			for i := range members {
				members[(i+1)%len(members)].GroupAccept()
			}
			robotgo.ActivePID(currentPid)
			robotgo.MilliSleep(chaojidou.WAITING_ACTIVE_PID_MILLI_SECONDS)
		}
	})

	fmt.Println("--- Please press shift + e to enter instance zones ---")
	hook.Register(hook.KeyDown, []string{robotgo.Shift, robotgo.KeyE}, func(e hook.Event) {
		fmt.Println("shift-e")
		if len(members) > 0 {
			currentPid := robotgo.GetPID()
			members[0].Empty()
			for i := range members {
				members[(i+1)%len(members)].EnterAccept()
			}
			robotgo.ActivePID(currentPid)
			robotgo.MilliSleep(chaojidou.WAITING_ACTIVE_PID_MILLI_SECONDS)
		}
	})

	fmt.Println("--- Please press shift + w to confirm or enter ---")
	hook.Register(hook.KeyDown, []string{robotgo.Shift, robotgo.KeyW}, func(e hook.Event) {
		fmt.Println("shift-w")
		if len(members) > 0 {
			currentPid := robotgo.GetPID()
			members[0].Empty()
			for i := range members {
				members[(i+1)%len(members)].ConfirmOrAccept()
			}
			robotgo.ActivePID(currentPid)
			robotgo.MilliSleep(chaojidou.WAITING_ACTIVE_PID_MILLI_SECONDS)
		}
	})

	s := hook.Start()
	<-hook.Process(s)
}
