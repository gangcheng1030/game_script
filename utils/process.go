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
			if sogouPs == nil || pss.Pid == parentPs.Pid || isParent(pss, parentPs) {
				sogouPs = pss
				parentPs, _ = pss.Parent()
			}
		}
	}

	return sogouPs, nil
}

func FindProcessByTimes(processName string) (*process.Process, error) {
	pses, _ := process.Processes()
	var sogouPs *process.Process
	var childrenNum int = 0
	for _, pss := range pses {
		name, _ := pss.Name()
		if strings.HasPrefix(name, processName) {
			children, _ := pss.Children()
			if len(children) >= childrenNum {
				sogouPs = pss
				childrenNum = len(children)
			}
		}
	}

	return sogouPs, nil
}

// 判断 p1 是否为 p2 的父进程
func isParent(p1 *process.Process, p2 *process.Process) bool {
	if p1.Pid == p2.Pid {
		return false
	}

	child := p2
	var parent *process.Process
	var err error
	for {
		parent, err = child.Parent()
		if err != nil {
			return false
		}
		if parent.Pid == p1.Pid {
			return true
		}
		child = parent
	}
}
