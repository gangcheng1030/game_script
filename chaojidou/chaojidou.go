package chaojidou

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"math/rand"
	"time"
)

//type ClientType uint8
//
//const (
//	CLIENT_TYPE_WINDOWS ClientType = iota
//	CLIENT_TYPE_SOGOU
//	CLIENT_TYPE_QQ
//	CLIENT_TYPE_T360
//	CLIENT_TYPE_CHROME
//)

func init() {
	rand.Seed(time.Now().Unix())
}

type ChaoJiDou interface {
	Active()
	Confirm()
	Accept()
	ConfirmOrAccept()
}

type chaoJiDou struct {
	Pid int32
	GameWindow robotgo.Rect
	AcceptButton robotgo.Rect
}

func (c *chaoJiDou) Confirm() {
	c.Active()
	c.confirm()
}

func (c *chaoJiDou) Accept()  {
	c.Active()
	c.accept()
}

func (c *chaoJiDou) ConfirmOrAccept() {
	c.Active()
	c.confirm()
	c.accept()
}

func (c *chaoJiDou) Active() {
	err := robotgo.ActivePID(c.Pid)
	if err != nil {
		fmt.Println(err)
	}
	robotgo.MilliSleep(WAITING_ACTIVE_PID_MILLI_SECONDS)
}

func (c *chaoJiDou) accept() {
	xr := rand.Intn(c.AcceptButton.W)
	yr := rand.Intn(c.AcceptButton.H)

	x := c.GameWindow.X + c.AcceptButton.X + xr
	y := c.GameWindow.Y + c.AcceptButton.Y + yr

	robotgo.MoveSmooth(x, y)
	robotgo.Click()
}

func (c *chaoJiDou) confirm() {
	robotgo.MilliSleep(WAITING_ACTIVE_PID_MILLI_SECONDS)
	robotgo.KeyToggle(robotgo.KeyF, int(c.Pid))
	robotgo.KeyToggle(robotgo.KeyF, int(c.Pid), robotgo.Up)
}