package utils

import (
	"github.com/go-vgo/robotgo"
	"math/rand"
)

func GetRandomPointInRect(rect robotgo.Rect) robotgo.Point {
	xr := rand.Intn(rect.W)
	yr := rand.Intn(rect.H)

	x := rect.X + xr
	y := rect.Y + yr

	return robotgo.Point{X: x, Y: y}
}

func GetRandomPointInRect1(x1, y1, w1, h1 int) robotgo.Point {
	xr := rand.Intn(w1)
	yr := rand.Intn(h1)

	x := x1 + xr
	y := y1 + yr

	return robotgo.Point{X: x, Y: y}
}
