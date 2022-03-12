package chaojidou

import "github.com/gangcheng1030/game_script/utils"

type ChromeChaoJiDou struct {
	chaoJiDou
}

func NewChromeChaoJiDou() ChaoJiDou {
	pss, _ := utils.FindProcessByName(CHROME_PROCESS_NAME)
	c := chaoJiDou{Pid: pss.Pid}
	return &ChromeChaoJiDou{chaoJiDou: c}
}
