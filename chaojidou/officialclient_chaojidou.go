package chaojidou

import (
	"github.com/gangcheng1030/game_script/utils"
	"github.com/go-vgo/robotgo"
)

type OfficialClientChaoJiDou struct {
	chaoJiDou
}

func NewOfficialClientChaoJiDou() ChaoJiDou {
	pss, _ := utils.FindProcessByName(OFFICIALCLIENT_PROCESS_NAME)
	gameWindow := robotgo.Rect{
		Point: robotgo.Point{X: 0, Y: 0},
		Size:  robotgo.Size{W: 1360, H: 768},
	}
	c := chaoJiDou{
		Pid:        pss.Pid,
		GameWindow: gameWindow,
	}
	return &OfficialClientChaoJiDou{chaoJiDou: c}
}
