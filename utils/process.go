package utils

import (
	"github.com/shirou/gopsutil/v3/process"
	"strings"
)

const (
	SOGOU_PROCESS_NAME = "SogouExplorer.exe"
	QQ_PROCESS_NAME    = "QQBrowser.exe"
)

func FindSogouProcess() (*process.Process, error) {
	pses, _ := process.Processes()
	var sogouPs *process.Process
	var parentPs *process.Process
	for _, pss := range pses {
		name, _ := pss.Name()
		if strings.HasPrefix(name, SOGOU_PROCESS_NAME) {
			if sogouPs == nil || pss.Pid == parentPs.Pid {
				sogouPs = pss
				parentPs, _ = pss.Parent()
			}
		}
	}

	return sogouPs, nil
}

func FindQQProcess() (*process.Process, error) {
	pses, _ := process.Processes()
	var sogouPs *process.Process
	var parentPs *process.Process
	for _, pss := range pses {
		name, _ := pss.Name()
		if strings.HasPrefix(name, QQ_PROCESS_NAME) {
			if sogouPs == nil || pss.Pid == parentPs.Pid {
				sogouPs = pss
				parentPs, _ = pss.Parent()
			}
		}
	}

	return sogouPs, nil
}
