package chaojidou

import (
	"github.com/gangcheng1030/game_script/utils"
	"github.com/go-vgo/robotgo"
)

type T360ChaoJiDou struct {
	chaoJiDou
}

func NewT360ChaoJiDou() ChaoJiDou {
	pss, _ := utils.FindProcessByName(T360_PROCESS_NAME)
	gameWindow := robotgo.Rect{
		Point: robotgo.Point{X: 131, Y: 103},
		Size:  robotgo.Size{W: 1070, H: 603},
	}
	groupAcceptButton := robotgo.Rect{
		Point: robotgo.Point{X: 507, Y: 327},
		Size:  robotgo.Size{W: 50, H: 11},
	}
	enterAcceptButton := robotgo.Rect{
		Point: robotgo.Point{X: 463, Y: 355},
		Size:  robotgo.Size{W: 62, H: 10},
	}
	c := chaoJiDou{
		Pid:               pss.Pid,
		GameWindow:        gameWindow,
		GroupAcceptButton: groupAcceptButton,
		EnterAcceptButton: enterAcceptButton,
	}
	return &T360ChaoJiDou{chaoJiDou: c}
}
