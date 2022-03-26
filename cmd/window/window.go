package main

import (
	"fmt"
	"github.com/gangcheng1030/game_script/utils"
	"github.com/go-vgo/robotgo"
	"github.com/lxn/win"
	"syscall"
)

func main() {
	pss, _ := utils.FindProcessByTimes("SogouExplorer.exe")
	handle, err := syscall.OpenProcess(0x0400|0x0010, false, uint32(pss.Pid))
	if err != nil {
		panic(err)
	}
	defer syscall.CloseHandle(handle)
	robotgo.ActivePID(pss.Pid)
	success := win.SetWindowPos(win.HWND(handle), win.HWND_TOPMOST, 0, 0, 0, 0, win.SWP_NOMOVE|win.SWP_NOSIZE)
	fmt.Println(success)
	robotgo.Sleep(10)
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
	//robotgo.ActivePID(5576)

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
