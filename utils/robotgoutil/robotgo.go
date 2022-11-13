package robotgoutil

import (
	"github.com/gangcheng1030/game_script/utils"
	"github.com/go-vgo/robotgo"
	"log"
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

func Click(x, y int, err int, tm int) {
	if err > 0 {
		x = x + rand.Intn(2*err) - err
		y = y + rand.Intn(2*err) - err
	}

	robotgo.MoveSmooth(x, y, 0.9, 0.9)
	robotgo.MilliSleep(200)
	robotgo.Click()
	log.Printf("click: %d, %d", x, y)
	if tm > 0 {
		robotgo.Sleep(tm)
	}
}
