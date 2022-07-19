package core

import "time"

type MeiriManager struct {
	meiri       string
	updatedTime time.Time
}

func (mm *MeiriManager) unsafeSet(meiri string) {
	mm.meiri = meiri
	mm.updatedTime = time.Now()
}
