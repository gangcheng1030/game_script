package chaojidou

import "github.com/gangcheng1030/game_script/utils"

type T360ChaoJiDou struct {
	chaoJiDou
}

func NewT360ChaoJiDou() ChaoJiDou {
	pss, _ := utils.FindProcessByName(T360_PROCESS_NAME)
	c := chaoJiDou{Pid: pss.Pid}
	return &T360ChaoJiDou{chaoJiDou: c}
}