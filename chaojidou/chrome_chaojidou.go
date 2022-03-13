package chaojidou

import (
	"github.com/gangcheng1030/game_script/utils"
	"github.com/go-vgo/robotgo"
)

type ChromeChaoJiDou struct {
	chaoJiDou
}

func NewChromeChaoJiDou() ChaoJiDou {
	pss, _ := utils.FindProcessByName(CHROME_PROCESS_NAME)
	gameWindow := robotgo.Rect{
		Point: robotgo.Point{X: 128, Y: 102},
		Size:  robotgo.Size{W: 1069, H: 603},
	}
	groupAcceptButton := robotgo.Rect{
		Point: robotgo.Point{X: 508, Y: 328},
		Size:  robotgo.Size{W: 54, H: 13},
	}
	enterAcceptButton := robotgo.Rect{
		Point: robotgo.Point{X: 460, Y: 355},
		Size:  robotgo.Size{W: 65, H: 10},
	}
	c := chaoJiDou{
		Pid:               pss.Pid,
		GameWindow:        gameWindow,
		GroupAcceptButton: groupAcceptButton,
		EnterAcceptButton: enterAcceptButton,
	}
	return &ChromeChaoJiDou{chaoJiDou: c}
}
