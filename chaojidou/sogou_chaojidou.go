package chaojidou

import "github.com/gangcheng1030/game_script/utils"

type SogouChaoJiDou struct {
	chaoJiDou
}

func NewSogouChaoJiDou() ChaoJiDou {
	pss, _ := utils.FindProcessByName(SOGOU_PROCESS_NAME)
	c := chaoJiDou{Pid: pss.Pid}
	return &SogouChaoJiDou{chaoJiDou: c}
}
