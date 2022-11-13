package main

import (
	"flag"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

var mode = flag.Int("m", 1, "mode")
var open = false

func main() {
	flag.Parse()

	evChan := hook.Start()
	defer hook.End()

	//preTime := time.Now()
	for ev := range evChan {
		if ev.Kind == hook.KeyDown {
			if ev.Rawcode == 110 || ev.Rawcode == 40 { // decimal point
				break
			}

			if ev.Rawcode == 83 { // s
				open = !open
				if open {
					switch *mode {
					case 1:
						go runMode1()
						break
					case 2:
						go runMode2()
						break
					default:
						go runMode1()
						break
					}
				}
				continue
			}
		}

		if !open {
			continue
		}
	}
}

func runMode1() {
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

	robotgo.MouseDown()

	go func() {
		for open {
			robotgo.MilliSleep(1500)
			robotgo.KeyPress(robotgo.KeyQ)
			robotgo.MilliSleep(1500)
		}
	}()

	for open {
		robotgo.Click("right")
		robotgo.Sleep(1)
		robotgo.KeyPress(robotgo.KeyW)
		robotgo.Sleep(1)
	}
}

func runMode2() {
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

	robotgo.MouseDown()

	go func() {
		for open {
			robotgo.MilliSleep(1500)
			robotgo.KeyPress(robotgo.KeyQ)
			robotgo.MilliSleep(1500)
		}
	}()

	go func() {
		for open {
			robotgo.MilliSleep(2500)
			robotgo.KeyPress(robotgo.KeyE)
			robotgo.MilliSleep(2000)
		}
	}()

	for open {
		robotgo.Click("right")
		robotgo.Sleep(1)
		robotgo.KeyPress(robotgo.KeyW)
		robotgo.Sleep(1)
	}
}
