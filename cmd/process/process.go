package main

import (
	"fmt"
	"github.com/gangcheng1030/game_script/chaojidou"
	"github.com/go-vgo/robotgo"
	"github.com/shirou/gopsutil/v3/process"
	"strings"
)

func main() {
	robotgo.MouseSleep = 100
	robotgo.KeySleep = 100

	//robotgo.ActivePID(15624)

	pses, _ := process.Processes()
	//sogouPses := make([]*process.Process, 0)
	for _, pss := range pses {
		name, _ := pss.Name()
		if strings.HasPrefix(name, chaojidou.QQ_PROCESS_NAME) {
			//sogouPses = append(sogouPses, pss)
			isBackground, _ := pss.Background()
			isForeground, _ := pss.Foreground()
			parentPs, _ := pss.Parent()
			children, _ := pss.Children()
			createTime, _ := pss.CreateTime()
			statuses, _ := pss.Status()
			fmt.Printf("pid: %d, name: %s, parent: %d, children: %d, isBackground: %v, isForeground: %v, createTime: %d, statuses: %v \n",
				pss.Pid, name, parentPs.Pid, len(children), isBackground, isForeground, createTime, statuses)
			//robotgo.ActivePID(pss.Pid)
			//robotgo.Sleep(5)
		}
	}
}

func contains(pses []*process.Process, pss *process.Process) bool {
	for _, tmp := range pses {
		if tmp.Pid == pss.Pid {
			return true
		}
	}

	return false
}
