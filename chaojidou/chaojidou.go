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
	GroupAccept()
	EnterAccept()
	ConfirmOrAccept()
	Esc()
	Empty()
}

type chaoJiDou struct {
	Pid               int32
	GameWindow        robotgo.Rect
	GroupAcceptButton robotgo.Rect
	EnterAcceptButton robotgo.Rect
}

func (c *chaoJiDou) Confirm() {
	c.Active()
	c.confirm()
}

func (c *chaoJiDou) GroupAccept() {
	c.Active()
	c.groupAccept()
}

func (c *chaoJiDou) EnterAccept() {
	c.Active()
	c.enterAccept()
}

func (c *chaoJiDou) ConfirmOrAccept() {
	c.Active()
	c.confirm()
	c.enterAccept()
}

func (c *chaoJiDou) Esc() {
	c.Active()
	c.esc()
}

func (c *chaoJiDou) Active() {
	err := robotgo.ActivePID(c.Pid)
	if err != nil {
		fmt.Println(err)
	}
	robotgo.MilliSleep(WAITING_ACTIVE_PID_MILLI_SECONDS)
}

func (c *chaoJiDou) Empty() {
	c.Active()
	robotgo.Click(robotgo.Down)
}

func (c *chaoJiDou) groupAccept() {
	xr := rand.Intn(c.GroupAcceptButton.W)
	yr := rand.Intn(c.GroupAcceptButton.H)

	x := c.GameWindow.X + c.GroupAcceptButton.X + xr
	y := c.GameWindow.Y + c.GroupAcceptButton.Y + yr

	fmt.Printf("x: %d, y: %d", x, y)
	robotgo.Move(x, y)
	robotgo.Click()
}

func (c *chaoJiDou) enterAccept() {
	xr := rand.Intn(c.EnterAcceptButton.W)
	yr := rand.Intn(c.EnterAcceptButton.H)

	x := c.GameWindow.X + c.EnterAcceptButton.X + xr
	y := c.GameWindow.Y + c.EnterAcceptButton.Y + yr

	robotgo.Move(x, y)
	robotgo.Click()
}

func (c *chaoJiDou) confirm() {
	robotgo.Click(robotgo.Down)
	robotgo.KeyToggle(robotgo.KeyF, int(c.Pid))
	robotgo.KeyToggle(robotgo.KeyF, int(c.Pid), robotgo.Up)
}

func (c *chaoJiDou) esc() {
	robotgo.Click(robotgo.Down)
	robotgo.KeyToggle(robotgo.Esc, int(c.Pid))
	robotgo.KeyToggle(robotgo.Esc, int(c.Pid), robotgo.Up)
}
