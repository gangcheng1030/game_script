package chaojidou

import (
	"errors"
	"fmt"
	"github.com/go-vgo/robotgo"
	"math/rand"
	"time"
)

type ClientType string

const (
	CLIENT_TYPE_OFFICIAL ClientType = "official_client"
	CLIENT_TYPE_SOGOU    ClientType = "sogou"
	CLIENT_TYPE_QQ       ClientType = "qq"
	CLIENT_TYPE_T360     ClientType = "360"
	CLIENT_TYPE_CHROME   ClientType = "chrome"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type ChaoJiDou interface {
	// 通用
	Active()
	Confirm()
	Empty()
	Esc()
	// LeftClick x: 横坐标, y: 纵坐标, w: 窗口宽, h: 窗口高
	LeftClick(x, y, w, h int)

	GetGameWindow() robotgo.Rect

	// 队员
	GroupAccept()
	EnterAccept()
	ConfirmOrAccept()

	// 队长
}

func Build(clientType ClientType) (ChaoJiDou, error) {
	if clientType == CLIENT_TYPE_OFFICIAL {
		return NewOfficialClientChaoJiDou(), nil
	} else if clientType == CLIENT_TYPE_SOGOU {
		return NewSogouChaoJiDou(), nil
	} else if clientType == CLIENT_TYPE_QQ {
		return NewQqChaoJiDou(), nil
	} else if clientType == CLIENT_TYPE_T360 {
		return NewT360ChaoJiDou(), nil
	} else if clientType == CLIENT_TYPE_CHROME {
		return NewChromeChaoJiDou(), nil
	}

	return nil, errors.New("clientType is invalid")
}

type chaoJiDou struct {
	Pid int32

	// 通用
	GameWindow robotgo.Rect

	// 队员
	GroupAcceptButton robotgo.Rect
	EnterAcceptButton robotgo.Rect

	// 队长
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
	c.enterAccept()
	c.confirm()
}

func (c *chaoJiDou) Esc() {
	c.Active()
	c.esc()
}

func (c *chaoJiDou) LeftClick(x, y, w, h int) {
	if x < 0 || x > w || y < 0 || y > h {
		return
	}

	xx := x*c.GameWindow.W/w + c.GameWindow.X
	yy := y*c.GameWindow.H/h + c.GameWindow.Y
	//fmt.Printf("xx: %d, yy: %d\n", xx, yy)
	robotgo.Move(xx, yy)
	robotgo.Click()
}

func (c *chaoJiDou) GetGameWindow() robotgo.Rect {
	return c.GameWindow
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
	robotgo.Click(robotgo.Right)
}

func (c *chaoJiDou) groupAccept() {
	xr := rand.Intn(c.GroupAcceptButton.W)
	yr := rand.Intn(c.GroupAcceptButton.H)

	x := c.GameWindow.X + c.GroupAcceptButton.X + xr
	y := c.GameWindow.Y + c.GroupAcceptButton.Y + yr

	//fmt.Printf("x: %d, y: %d", x, y)
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
	robotgo.Click()
	robotgo.KeyToggle(robotgo.KeyF, int(c.Pid))
	robotgo.KeyToggle(robotgo.KeyF, int(c.Pid), robotgo.Up)
}

func (c *chaoJiDou) esc() {
	robotgo.Click()
	robotgo.KeyToggle(robotgo.Esc, int(c.Pid))
	robotgo.KeyToggle(robotgo.Esc, int(c.Pid), robotgo.Up)
}
