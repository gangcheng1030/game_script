package chaojidou

import "github.com/gangcheng1030/game_script/utils"

type OfficialClientChaoJiDou struct {
	chaoJiDou
}

func NewOfficialClientChaoJiDou() ChaoJiDou {
	pss, _ := utils.FindProcessByName(OFFICIALCLIENT_PROCESS_NAME)
	c := chaoJiDou{Pid: pss.Pid}
	return &OfficialClientChaoJiDou{chaoJiDou: c}
}