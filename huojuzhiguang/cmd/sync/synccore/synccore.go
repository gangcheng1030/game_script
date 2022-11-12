package synccore

import (
	"github.com/gangcheng1030/game_script/huojuzhiguang/cmd/sync/follower_client"
	"github.com/gangcheng1030/game_script/huojuzhiguang/cmd/sync/global"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"math/rand"
	"sync"
)

func HandleS() {
	global.Open = !global.Open
	if global.Open {
		switch global.Mode {
		case 1:
			go RunMode1()
			break
		case 2:
			go RunMode2()
			break
		default:
			go RunMode1()
			break
		}
	}
}

func HandleFollowers(e *hook.Event, st int, st1 int) {
	if len(global.Followers) == 0 {
		return
	}

	if st > 0 {
		robotgo.Sleep(st)
	}
	var waitGroup sync.WaitGroup
	for _, addr := range global.Followers {
		tmpAddr := addr
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			if st1 > 0 {
				st1 = rand.Intn(st1 * 1000)
				robotgo.MilliSleep(st1)
			}

			err := follower_client.SendEvent(tmpAddr, e)
			if err != nil {
				panic(err)
			}
		}()
	}

	waitGroup.Wait()
}

func RunMode1() {
	robotgo.KeyPress(robotgo.KeyE)
	robotgo.MilliSleep(500)
	robotgo.KeyPress(robotgo.KeyR)
	robotgo.MilliSleep(500)
	robotgo.KeyPress(robotgo.KeyE)
	robotgo.MilliSleep(500)
	robotgo.KeyPress(robotgo.KeyR)
	robotgo.MilliSleep(500)
	robotgo.KeyPress(robotgo.KeyE)
	robotgo.MilliSleep(500)
	robotgo.KeyPress(robotgo.KeyR)
	robotgo.MilliSleep(500)
	robotgo.KeyPress(robotgo.KeyE)
	robotgo.MilliSleep(500)
	robotgo.KeyPress(robotgo.KeyR)
	robotgo.MilliSleep(500)

	go func() {
		robotgo.MilliSleep(4500)
		for global.Open {
			robotgo.MouseDown()
			robotgo.MilliSleep(900)
			robotgo.MouseUp()
			robotgo.MilliSleep(100)
		}
	}()

	go func() {
		for global.Open {
			robotgo.MilliSleep(1500)
			robotgo.KeyPress(robotgo.KeyQ)
			robotgo.MilliSleep(1500)
		}
	}()

	for global.Open {
		robotgo.Click("right")
		robotgo.Sleep(1)
		robotgo.KeyPress(robotgo.KeyW)
		robotgo.Sleep(1)
	}
}

func RunMode2() {
	robotgo.KeyPress(robotgo.KeyR)
	robotgo.MilliSleep(500)
	robotgo.KeyPress(robotgo.KeyR)
	robotgo.MilliSleep(500)
	robotgo.KeyPress(robotgo.KeyR)
	robotgo.MilliSleep(500)
	robotgo.KeyPress(robotgo.KeyR)
	robotgo.MilliSleep(500)
	robotgo.KeyPress(robotgo.KeyR)
	robotgo.MilliSleep(500)
	robotgo.KeyPress(robotgo.KeyR)
	robotgo.MilliSleep(500)

	go func() {
		robotgo.MilliSleep(4500)
		for global.Open {
			robotgo.MouseDown()
			robotgo.MilliSleep(900)
			robotgo.MouseUp()
			robotgo.MilliSleep(100)
		}
	}()

	go func() {
		for global.Open {
			robotgo.MilliSleep(1500)
			robotgo.KeyPress(robotgo.KeyQ)
			robotgo.MilliSleep(1500)
		}
	}()

	go func() {
		for global.Open {
			robotgo.MilliSleep(2500)
			robotgo.KeyPress(robotgo.KeyE)
			robotgo.MilliSleep(2000)
		}
	}()

	for global.Open {
		robotgo.Click("right")
		robotgo.Sleep(1)
		robotgo.KeyPress(robotgo.KeyW)
		robotgo.Sleep(1)
	}
}
