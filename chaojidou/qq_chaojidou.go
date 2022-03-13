package chaojidou

import (
	"github.com/gangcheng1030/game_script/utils"
	"github.com/go-vgo/robotgo"
)

type QqChaoJiDou struct {
	chaoJiDou
}

func NewQqChaoJiDou() ChaoJiDou {
	pss, _ := utils.FindProcessByName(QQ_PROCESS_NAME)
	gameWindow := robotgo.Rect{
		Point: robotgo.Point{X: 121, Y: 96},
		Size:  robotgo.Size{W: 1077, H: 607},
	}
	groupAcceptButton := robotgo.Rect{
		Point: robotgo.Point{X: 512, Y: 329},
		Size:  robotgo.Size{W: 55, H: 12},
	}
	enterAcceptButton := robotgo.Rect{
		Point: robotgo.Point{X: 465, Y: 357},
		Size:  robotgo.Size{W: 64, H: 10},
	}
	c := chaoJiDou{
		Pid:               pss.Pid,
		GameWindow:        gameWindow,
		GroupAcceptButton: groupAcceptButton,
		EnterAcceptButton: enterAcceptButton,
	}
	return &QqChaoJiDou{chaoJiDou: c}
}
