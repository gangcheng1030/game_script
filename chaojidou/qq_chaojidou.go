package chaojidou

import "github.com/gangcheng1030/game_script/utils"

type QqChaoJiDou struct {
	chaoJiDou
}

func NewQqChaoJiDou() ChaoJiDou {
	pss, _ := utils.FindProcessByName(QQ_PROCESS_NAME)
	c := chaoJiDou{Pid: pss.Pid}
	return &QqChaoJiDou{chaoJiDou: c}
}
