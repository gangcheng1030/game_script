package robotgoutil

import (
	"github.com/gangcheng1030/game_script/utils"
	"github.com/go-vgo/robotgo"
	"math/rand"
)

func Press(key string, tm int) {
	robotgo.KeyPress(key)
	if tm > 0 {
		rtm := tm*1000 + rand.Intn(400) - 200
		robotgo.MilliSleep(rtm)
	}
}

// ClickButton tm单位：秒
func ClickButton(button robotgo.Rect, tm int) {
	tmpPoint := utils.GetRandomPointInRect(button)
	robotgo.MoveSmooth(tmpPoint.X, tmpPoint.Y, 0.9, 0.9)
	robotgo.MilliSleep(200)
	robotgo.Click()
	if tm > 0 {
		robotgo.Sleep(tm)
	}
}
