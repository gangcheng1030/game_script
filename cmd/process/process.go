package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/shirou/gopsutil/v3/process"
	"strings"
)

func main() {
	robotgo.MouseSleep = 100
	robotgo.KeySleep = 100

	pses, _ := process.Processes()
	//sogouPses := make([]*process.Process, 0)
	for _, pss := range pses {
		name, _ := pss.Name()
		if strings.HasPrefix(name, "QQBrowser.exe") {
			//sogouPses = append(sogouPses, pss)
			isBackground, _ := pss.Background()
			parentPs, _ := pss.Parent()
			fmt.Printf("%d, %s, %d, %d\n", pss.Pid, name, parentPs.Pid, isBackground)
			robotgo.ActivePID(pss.Pid)
			robotgo.Sleep(5)
		}
	}
	//qqPs, _ := processutil.FindQQProcess()
	//robotgo.ActivePID(qqPs.Pid)
	//robotgo.Sleep(5)
}

func contains(pses []*process.Process, pss *process.Process) bool {
	for _, tmp := range pses {
		if tmp.Pid == pss.Pid {
			return true
		}
	}

	return false
}
