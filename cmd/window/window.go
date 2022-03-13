package main

import (
	"github.com/go-vgo/robotgo"
)

func main() {
	//fpid, err := robotgo.FindIds("Google")
	//if err == nil {
	//	fmt.Println("pids... ", fpid)
	//
	//	if len(fpid) > 0 {
	//		//robotgo.TypeStr("Hi galaxy!", int(fpid[0]))
	//		robotgo.ActivePID(fpid[0])
	//		robotgo.KeyTap("enter")
	//		//robotgo.KeyToggle("a")
	//		//robotgo.KeyToggle("cmd")
	//		//robotgo.Sleep(1)
	//		//robotgo.KeyToggle("a", "up")
	//		//robotgo.KeyToggle("cmd", "up")
	//	}
	//}

	// 5576 5064
	robotgo.ActivePID(5576)

	//isExist, err := robotgo.PidExists(100)
	//if err == nil && isExist {
	//	fmt.Println("pid exists is", isExist)
	//
	//	robotgo.Kill(100)
	//}
	//
	//abool := robotgo.Alert("test", "robotgo")
	//if abool {
	//	fmt.Println("ok@@@ ", "ok")
	//}
	//
	//title := robotgo.GetTitle()
	//fmt.Println("title@@@ ", title)
}
