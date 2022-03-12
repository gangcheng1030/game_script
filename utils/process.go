package utils

import (
	"github.com/shirou/gopsutil/v3/process"
	"strings"
)

func FindProcessByName(processName string) (*process.Process, error) {
	pses, _ := process.Processes()
	var sogouPs *process.Process
	var parentPs *process.Process
	for _, pss := range pses {
		name, _ := pss.Name()
		if strings.HasPrefix(name, processName) {
			if sogouPs == nil || pss.Pid == parentPs.Pid {
				sogouPs = pss
				parentPs, _ = pss.Parent()
			}
		}
	}

	return sogouPs, nil
}
