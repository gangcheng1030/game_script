package chaojidou

import (
	"github.com/gangcheng1030/game_script/utils"
	"github.com/go-vgo/robotgo"
)

type SogouChaoJiDou struct {
	chaoJiDou
}

func NewSogouChaoJiDou() ChaoJiDou {
	pss, _ := utils.FindProcessByTimes(SOGOU_PROCESS_NAME)
	gameWindow := robotgo.Rect{
		Point: robotgo.Point{X: 142, Y: 95},
		Size:  robotgo.Size{W: 1048, H: 613},
	}
	groupAcceptButton := robotgo.Rect{
		Point: robotgo.Point{X: 501, Y: 322},
		Size:  robotgo.Size{W: 49, H: 13},
	}
	enterAcceptButton := robotgo.Rect{
		Point: robotgo.Point{X: 451, Y: 350},
		Size:  robotgo.Size{W: 65, H: 9},
	}
	c := chaoJiDou{
		Pid:               pss.Pid,
		GameWindow:        gameWindow,
		GroupAcceptButton: groupAcceptButton,
		EnterAcceptButton: enterAcceptButton,
	}
	return &SogouChaoJiDou{chaoJiDou: c}
}
