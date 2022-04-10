package chaojidou

import (
	"errors"
	"fmt"
	"github.com/gangcheng1030/game_script/utils"
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

type DifficultyType int

const (
	DIFFICULTY_TYPE_MAOXIAN DifficultyType = iota
	DIFFICULTY_TYPE_SHULIAN
	DIFFICULTY_TYPE_YINGXIONG
	DIFFICULTY_TYPE_CHUANSHUO
)

type JinBenType int

const (
	JINBEN_TYPE_SUXING JinBenType = iota
	JINBEN_TYPE_HEIAN
)

type LiuLangTuanType int

const (
	LIULANGTUAN_TYPE_1 LiuLangTuanType = iota
)

func init() {
	rand.Seed(time.Now().Unix())
}

type ChaoJiDou interface {
	// 通用
	Active() bool
	Confirm()
	Empty()
	Esc()
	// LeftClick x: 横坐标, y: 纵坐标, w: 窗口宽, h: 窗口高
	LeftClick(x, y, w, h int) bool
	JuQing(jt JuQingType)
	JiuYunDong(dt DifficultyType, times int)
	MeiRiTiaoZhan(mt MeiRiType, dt DifficultyType)
	LiuLangTuan(lt LiuLangTuanType, dt DifficultyType)
	JinBen(jt JinBenType, dt DifficultyType, times int)

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
	GameWindow     robotgo.Rect
	BigMap         BigMap
	RenWuMap       RenWuMap
	JuQingMap      JuQingMap
	ZhuiSuMap      ZhuiSuMap
	MeiRiMap       MeiRiMap
	LiuLangTuanMap LiuLangTuanMap
	JinBenMap      JinBenMap

	StartBattleButton robotgo.Rect
	EnterButton       robotgo.Rect // 入场按钮
	EnterDButton      robotgo.Rect // 推荐难度，确认入场按钮

	// 队员
	GroupAcceptButton robotgo.Rect
	EnterAcceptButton robotgo.Rect

	// 队长
}

type BigMap struct {
	ZhuiSu      robotgo.Rect
	MeiRi       robotgo.Rect
	LiuLangTuan robotgo.Rect
	Hermosi     robotgo.Rect
}

type RenWuMap struct {
	ZhuYe      robotgo.Rect
	FirstXunLu robotgo.Rect
}

type JuQingMap struct {
	TongHuaZhenButton robotgo.Rect

	FuBens map[JuQingType]FuBen
}

type ZhuiSuMap struct {
	EnterButton robotgo.Rect

	JiuYunDong FuBen
}

type MeiRiMap struct {
	Window              robotgo.Rect
	DifficultyTypePoses []robotgo.Rect

	ShenHaiZhuKe    []JuQingType
	ZuiQiangJianShi []JuQingType
	TieBiShouWei    []JuQingType
	YanHua          []JuQingType
	BingFengWangGuo []JuQingType
	XiaoXinChuDian  []JuQingType
}

type LiuLangTuanMap struct {
	EnterButton robotgo.Rect

	FuBenArray []FuBen
}

type JinBenMap struct {
	EnterButton     robotgo.Rect
	ChuanShaoButton robotgo.Rect

	FuBenArray []FuBen
}

type FuBen struct {
	Window              robotgo.Rect
	DifficultyTypePoses []robotgo.Rect
	SmallMap            []robotgo.Rect
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

func (c *chaoJiDou) LeftClick(x, y, w, h int) bool {
	active := c.Active()
	if !active {
		return false
	}

	if x < 0 || x > w || y < 0 || y > h {
		return true
	}

	xx := x*c.GameWindow.W/w + c.GameWindow.X
	yy := y*c.GameWindow.H/h + c.GameWindow.Y
	fmt.Printf("xx: %d, yy: %d\n", xx, yy)
	robotgo.Move(xx, yy)
	robotgo.Click()
	return true
}

func (c *chaoJiDou) JiuYunDong(dt DifficultyType, times int) {
	c.press(robotgo.KeyM, 1)

	// 大地图点追溯
	c.clickButton(c.BigMap.ZhuiSu, NpcWaitSecs)

	// 追溯地图点击‘九云洞’
	robotgo.KeyPress(robotgo.KeyF)
	robotgo.Sleep(3)
	c.clickButton(c.ZhuiSuMap.JiuYunDong.Window, 3)

	// 选难度
	c.clickButton(c.ZhuiSuMap.JiuYunDong.DifficultyTypePoses[dt], 2)

	// 入场
	c.clickButton(c.ZhuiSuMap.EnterButton, ReadMapWaitSecs)

	// 打怪
	c.jiuYunDongHelper()

	// 重复刷
	for remain := times - 1; remain > 0; remain-- {
		robotgo.KeyPress(robotgo.F11)
		robotgo.Sleep(3)
		robotgo.KeyPress(robotgo.KeyF)
		robotgo.Sleep(3)
		robotgo.KeyPress(robotgo.KeyF)
		robotgo.Sleep(30)
		c.jiuYunDongHelper()
	}

	// 返回主城
	robotgo.KeyPress(robotgo.F12)
	robotgo.Sleep(5)
	robotgo.KeyPress(robotgo.KeyF)
	robotgo.Sleep(3)
	robotgo.KeyPress(robotgo.KeyF)
	robotgo.Sleep(20)
}

func (c *chaoJiDou) jiuYunDongHelper() {
	// 开始战斗
	c.clickButton(c.StartBattleButton, 7)

	// 技能连招
	robotgo.KeyPress(robotgo.KeyQ)
	robotgo.Sleep(4)
	robotgo.KeyPress(robotgo.Key3)
	robotgo.Sleep(2)
	robotgo.KeyPress(robotgo.KeyD)
	robotgo.Sleep(1)
	robotgo.KeyPress(robotgo.KeyW)
	robotgo.Sleep(1)

	// 第1张怪物图
	c.clickButton(c.ZhuiSuMap.JiuYunDong.SmallMap[0], 10)

	c.move(950, 225, 10, 5)

	// 第2张怪物图
	c.clickButton(c.ZhuiSuMap.JiuYunDong.SmallMap[1], 8)

	c.press(robotgo.KeyD, 1)
	c.move(1086, 71, 2, 8)

	// 第3张怪物图
	c.clickButton(c.ZhuiSuMap.JiuYunDong.SmallMap[2], 8)

	c.press(robotgo.KeyD, 1)
	c.move(1086, 71, 2, 8)

	// 第4张怪物图
	c.clickButton(c.ZhuiSuMap.JiuYunDong.SmallMap[3], 8)

	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1082, 97, 2, 8)

	// 第5张怪物图
	c.clickButton(c.ZhuiSuMap.JiuYunDong.SmallMap[4], 8)

	c.press(robotgo.KeyD, 1)
	c.move(299, 141, 2, 8)

	// 第6张怪物图
	c.clickButton(c.ZhuiSuMap.JiuYunDong.SmallMap[5], 8)

	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(633, 91, 2, 8)

	// 第7张怪物图
	c.clickButton(c.ZhuiSuMap.JiuYunDong.SmallMap[6], 8)

	c.press(robotgo.KeyD, 1)
	c.move(1110, 422, 2, 4)
	c.move(1020, 106, 2, 8)

	// 第8张怪物图：boss
	c.clickButton(c.ZhuiSuMap.JiuYunDong.SmallMap[7], 8)

	robotgo.KeyPress(robotgo.Esc)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	robotgo.Sleep(22)
	c.move(357, 578, 2, 5)
	robotgo.KeyPress(robotgo.KeyF)
	robotgo.Sleep(2)
}

func (c *chaoJiDou) LiuLangTuan(lt LiuLangTuanType, dt DifficultyType) {
	c.press(robotgo.KeyM, 1)

	// 大地图点流浪团
	c.clickButton(c.BigMap.LiuLangTuan, NpcWaitSecs)

	// 对话
	robotgo.KeyPress(robotgo.KeyF)
	robotgo.Sleep(3)
	robotgo.KeyPress(robotgo.KeyF)
	robotgo.Sleep(3)

	// 选中相应地图
	c.clickButton(c.LiuLangTuanMap.FuBenArray[lt].Window, 3)

	// 选难度
	c.clickButton(c.LiuLangTuanMap.FuBenArray[lt].DifficultyTypePoses[dt], 2)

	// 入场
	c.clickButton(c.LiuLangTuanMap.EnterButton, 3)
	c.clickButton(c.EnterDButton, 1) // 这个入场老是出bug，所以点2次
	c.clickButton(c.EnterDButton, ReadMapWaitSecs)

	// 打怪
	switch lt {
	case LIULANGTUAN_TYPE_1:
		c.liuLangTuan1Helper()
	default:
		// do nothing
	}

	// 返回主城
	robotgo.KeyPress(robotgo.F12)
	robotgo.Sleep(5)
	robotgo.KeyPress(robotgo.KeyF)
	robotgo.Sleep(3)
	robotgo.KeyPress(robotgo.KeyF)
	robotgo.Sleep(20)
}

func (c *chaoJiDou) liuLangTuan1Helper() {
	// 开始战斗
	c.clickButton(c.StartBattleButton, 7)

	// 技能连招
	robotgo.KeyPress(robotgo.KeyQ)
	robotgo.Sleep(4)
	robotgo.KeyPress(robotgo.Key3)
	robotgo.Sleep(4)
	robotgo.KeyPress(robotgo.KeyD)
	robotgo.Sleep(3)
	robotgo.KeyPress(robotgo.KeyW)
	robotgo.Sleep(3)

	// 第一阶段
	c.move(1263, 49, 2, 5)
	c.continuedBattle(60)

	// 第二阶段
	for i := 0; i < 5; i++ {
		c.move(1077, 534, 1, 1)
	}
	for i := 0; i < 7; i++ {
		c.move(931, 111, 1, 1)
	}
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.KeyW, 1)
	c.continuedBattle(60)

	// 第三阶段
	for i := 0; i < 9; i++ {
		c.move(260, 711, 1, 1)
	}
	c.move(401, 41, 2, 4)
	robotgo.MoveSmooth(680, 258, 0.9, 0.9)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.KeyW, 1)
	c.continuedBattle(15)
	c.press(robotgo.KeyT, 4)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.KeyW, 1)
	c.continuedBattle(45)
}

func (c *chaoJiDou) JinBen(jt JinBenType, dt DifficultyType, times int) {
	c.press(robotgo.KeyM, 1)

	// 大地图点金本
	c.clickButton(c.BigMap.Hermosi, NpcWaitSecs)

	// 对话
	robotgo.KeyPress(robotgo.KeyF)
	robotgo.Sleep(2)
	robotgo.KeyPress(robotgo.KeyF)
	robotgo.Sleep(2)

	// 选中船艄
	c.clickButton(c.JinBenMap.ChuanShaoButton, 2)

	// 选中相应地图
	c.clickButton(c.JinBenMap.FuBenArray[jt].Window, 2)

	// 选难度
	c.clickButton(c.JinBenMap.FuBenArray[jt].DifficultyTypePoses[dt], 2)

	// 入场
	c.clickButton(c.JinBenMap.EnterButton, 2)
	c.clickButton(c.EnterDButton, 1) // 这个入场老是出bug，所以点2次
	c.clickButton(c.EnterDButton, ReadMapWaitSecs)

	// 打怪
	switch jt {
	case JINBEN_TYPE_SUXING:
		c.suXingDeChuanShuoHelper()
	case JINBEN_TYPE_HEIAN:
		c.heiAnQinShiZhiHuanHelper()
	default:
		// do nothing
	}

	// 重复刷
	for remain := times - 1; remain > 0; remain-- {
		robotgo.KeyPress(robotgo.F11)
		robotgo.Sleep(5)
		robotgo.KeyPress(robotgo.KeyF)
		robotgo.Sleep(3)
		robotgo.KeyPress(robotgo.KeyF)
		robotgo.Sleep(25)
		switch jt {
		case JINBEN_TYPE_SUXING:
			c.suXingDeChuanShuoHelper()
		case JINBEN_TYPE_HEIAN:
			c.heiAnQinShiZhiHuanHelper()
		default:
			// do nothing
		}
	}

	// 返回主城
	robotgo.KeyPress(robotgo.F12)
	robotgo.Sleep(5)
	robotgo.KeyPress(robotgo.KeyF)
	robotgo.Sleep(3)
	robotgo.KeyPress(robotgo.KeyF)
	robotgo.Sleep(20)
	//c.move(402, 258, 5, 3)
}

func (c *chaoJiDou) suXingDeChuanShuoHelper() {
	// 开始战斗
	c.clickButton(c.StartBattleButton, 7)

	// 第1张怪物图
	c.clickButton(c.JinBenMap.FuBenArray[0].SmallMap[0], 10)
	c.press(robotgo.KeyD, 1)
	robotgo.MoveSmooth(336, 345, 0.9, 0.9)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyR, 1)
	c.multiMove(336, 345, 2, 1, 3)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	c.move(602, 620, 2, 3)

	// 第2张怪物图
	c.clickButton(c.JinBenMap.FuBenArray[0].SmallMap[1], 6)
	c.press(robotgo.KeyD, 1)
	robotgo.MoveSmooth(325, 314, 0.9, 0.9)
	c.press(robotgo.Key1, 4)
	c.press(robotgo.KeyE, 3)
	c.press(robotgo.KeyW, 1)
	c.press(robotgo.KeyQ, 4)
	c.multiMove(325, 314, 2, 1, 3)
	robotgo.Sleep(3)

	// 第3张怪物图
	c.clickButton(c.JinBenMap.FuBenArray[0].SmallMap[2], 6)
	c.press(robotgo.KeyD, 1)
	robotgo.MoveSmooth(1055, 299, 0.9, 0.9)
	c.press(robotgo.Key3, 2)
	c.multiMove(1055, 299, 2, 1, 3)
	robotgo.Sleep(8)

	// 第4张怪物图
	c.clickButton(c.JinBenMap.FuBenArray[0].SmallMap[3], 6)
	c.press(robotgo.KeyD, 1)
	robotgo.MoveSmooth(957, 200, 0.9, 0.9)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyW, 1)
	c.multiMove(957, 200, 2, 1, 3)
	robotgo.Sleep(3)

	// 第5张怪物图
	c.clickButton(c.JinBenMap.FuBenArray[0].SmallMap[4], 6)
	c.press(robotgo.KeyD, 1)
	robotgo.MoveSmooth(1057, 319, 0.9, 0.9)
	c.multiMove(1057, 319, 2, 1, 3)
	c.press(robotgo.Key3, 2)
	c.multiMove(1057, 319, 2, 1, 3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)

	// 第6张怪物图
	c.clickButton(c.JinBenMap.FuBenArray[0].SmallMap[5], 6)
	c.press(robotgo.KeyD, 1)
	robotgo.MoveSmooth(841, 496, 0.9, 0.9)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyW, 1)
	c.multiMove(841, 496, 2, 1, 3)
	robotgo.Sleep(3)

	// 第7张怪物图
	c.clickButton(c.JinBenMap.FuBenArray[0].SmallMap[6], 5)
	robotgo.MoveSmooth(216, 389, 0.9, 0.9)
	c.press(robotgo.KeyT, 4)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.F1, 1)
	c.multiMove(216, 389, 2, 1, 3)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.KeyE, 1)
	robotgo.Sleep(3)

	// 第8张怪物图
	c.clickButton(c.JinBenMap.FuBenArray[0].SmallMap[7], 6)
	c.press(robotgo.KeyD, 1)
	robotgo.MoveSmooth(673, 550, 0.9, 0.9)
	c.multiMove(673, 550, 2, 1, 5)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyW, 1)
	robotgo.Sleep(15)

	// 第9张怪物图
	c.clickButton(c.JinBenMap.FuBenArray[0].SmallMap[8], 6)
	c.press(robotgo.KeyD, 1)
	robotgo.MoveSmooth(942, 242, 0.9, 0.9)
	c.press(robotgo.Key3, 2)
	c.multiMove(942, 242, 2, 1, 3)
	robotgo.Sleep(3)

	// 第10张怪物图：实际上是回到第8张
	c.press(robotgo.F2, 1)
	c.clickButton(c.JinBenMap.FuBenArray[0].SmallMap[9], 8)

	// 第11张怪物图：boss
	c.clickButton(c.JinBenMap.FuBenArray[0].SmallMap[10], 10)
	robotgo.MoveSmooth(490, 320, 0.9, 0.9)
	c.press(robotgo.Esc, 1)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyW, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(20)
	c.move(138, 654, 2, 3)
	robotgo.KeyPress(robotgo.KeyF)
	robotgo.Sleep(2)
}

func (c *chaoJiDou) heiAnQinShiZhiHuanHelper() {
	// 开始战斗
	c.clickButton(c.StartBattleButton, 7)

	// 第1张怪物图
	c.clickButton(c.JinBenMap.FuBenArray[1].SmallMap[0], 8)
	c.press(robotgo.KeyS, 2)
	robotgo.Sleep(3)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	c.multiMove(1090, 345, 2, 1, 3)
	robotgo.MoveSmooth(956, 121, 0.9, 0.9)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyD, 1)
	c.multiMove(956, 121, 2, 1, 3)
	c.press(robotgo.KeyW, 1)
	c.press(robotgo.F1, 1)
	c.multiMove(158, 196, 2, 1, 1)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(5)
	c.press(robotgo.F2, 1)

	// 第2张怪物图
	c.clickButton(c.JinBenMap.FuBenArray[1].SmallMap[1], 5)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	c.multiMove(314, 159, 2, 1, 3)
	robotgo.Sleep(3)
	c.multiMove(314, 159, 2, 1, 2)
	robotgo.Sleep(2)

	// 第3张怪物图
	c.clickButton(c.JinBenMap.FuBenArray[1].SmallMap[2], 4)
	robotgo.MoveSmooth(392, 370, 0.9, 0.9)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyW, 1)
	c.multiMove(392, 370, 2, 1, 3)
	c.press(robotgo.F1, 1)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)

	// 第4张怪物图
	c.clickButton(c.JinBenMap.FuBenArray[1].SmallMap[3], 6)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyW, 1)
	c.multiMove(595, 167, 2, 1, 4)
	c.press(robotgo.KeyD, 1)
	c.multiMove(135, 197, 2, 1, 3)
	robotgo.Sleep(3)

	// 第5张怪物图
	c.clickButton(c.JinBenMap.FuBenArray[1].SmallMap[4], 5)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyW, 1)
	c.multiMove(303, 618, 2, 1, 3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)
	c.multiMove(303, 618, 2, 1, 3)
	robotgo.Sleep(3)

	// 第6张怪物图
	c.clickButton(c.JinBenMap.FuBenArray[1].SmallMap[5], 10)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyW, 1)
	c.multiMove(1121, 490, 2, 1, 4)
	robotgo.Sleep(6)

	// 第7张怪物图
	c.clickButton(c.JinBenMap.FuBenArray[1].SmallMap[6], 8)

	// 第8张怪物图
	robotgo.Sleep(3)
	c.clickButton(c.JinBenMap.FuBenArray[1].SmallMap[7], 10)
	c.press(robotgo.KeyD, 1)
	robotgo.MoveSmooth(433, 143, 0.9, 0.9)
	c.press(robotgo.Key1, 3)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyW, 1)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(10)

	// 第9张怪物图
	c.clickButton(c.JinBenMap.FuBenArray[1].SmallMap[8], 6)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyW, 1)
	c.multiMove(386, 195, 2, 1, 3)
	c.press(robotgo.KeyD, 1)
	c.multiMove(386, 195, 2, 1, 2)
	robotgo.Sleep(5)

	// 第10张怪物图：实际上是回到第8张
	c.clickButton(c.JinBenMap.FuBenArray[1].SmallMap[9], 5)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyW, 1)
	c.multiMove(942, 205, 2, 1, 5)
	c.press(robotgo.KeyD, 1)
	c.multiMove(942, 205, 2, 1, 5)
	robotgo.Sleep(3)

	// 第11张怪物图：boss
	c.clickButton(c.JinBenMap.FuBenArray[1].SmallMap[10], 3)
	c.press(robotgo.Esc, 1)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyT, 4)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyW, 1)
	c.multiMove(1018, 158, 2, 1, 3)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyE, 4)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(20)
}

func (c *chaoJiDou) GetGameWindow() robotgo.Rect {
	return c.GameWindow
}

func (c *chaoJiDou) Active() bool {
	err := robotgo.ActivePID(c.Pid)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if robotgo.GetPID() != c.Pid {
		return false
	}
	robotgo.MilliSleep(WAITING_ACTIVE_PID_MILLI_SECONDS)
	return true
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

// tm单位：秒
func (c *chaoJiDou) clickButton(button robotgo.Rect, tm int) {
	tmpPoint := utils.GetRandomPointInRect(button)
	robotgo.MoveSmooth(tmpPoint.X+c.GameWindow.X, tmpPoint.Y+c.GameWindow.Y, mouseSpeedX, mouseSpeedY)
	robotgo.MilliSleep(200)
	robotgo.Click()
	if tm > 0 {
		robotgo.Sleep(tm)
	}
}

// errors：误差
func (c *chaoJiDou) multiMove(x, y int, errors int, tm int, nums int) {
	for i := 0; i < nums; i++ {
		c.move(x, y, errors, tm)
	}
}

// errors：误差
func (c *chaoJiDou) move(x, y int, errors int, tm int) {
	tmpPoint := utils.GetRandomPointInRect1(x-errors, y-errors, 2*errors, 2*errors)
	robotgo.MoveSmooth(tmpPoint.X+c.GameWindow.X, tmpPoint.Y+c.GameWindow.Y, mouseSpeedX, mouseSpeedY)
	robotgo.Click("right")
	if tm > 0 {
		rtm := tm*1000 + rand.Intn(400) - 200
		robotgo.MilliSleep(rtm)
	}
}

func (c *chaoJiDou) press(key string, tm int) {
	robotgo.KeyPress(key)
	if tm > 0 {
		rtm := tm*1000 + rand.Intn(400) - 200
		robotgo.MilliSleep(rtm)
	}
}

func (c *chaoJiDou) continuedBattle(tm int) {
	doneTimer := time.NewTimer(time.Second * time.Duration(tm))
	d1Timer := time.NewTimer(8 * time.Second)
	d2Timer := time.NewTimer(3 * time.Second)
	d3Timer := time.NewTimer(8 * time.Second)
	d4Timer := time.NewTimer(10 * time.Second)
	d5Timer := time.NewTimer(13 * time.Second)
	//d6Timer := time.NewTimer(5 * time.Second)
	gpTimer := time.NewTimer(5 * time.Second)
	rTimer := time.NewTimer(15 * time.Second)
	bTimer := time.NewTimer(15 * time.Second)
	defer func() {
		d1Timer.Stop()
		d2Timer.Stop()
		d3Timer.Stop()
		d4Timer.Stop()
		d5Timer.Stop()
		//d6Timer.Stop()
		gpTimer.Stop()
		rTimer.Stop()
		bTimer.Stop()
	}()

	for {
		select {
		case <-doneTimer.C:
			return
		case <-d1Timer.C:
			if rand.Intn(2) == 0 {
				c.press(robotgo.Key3, 2)
			}
			d1Timer.Reset(8 * time.Second)
		case <-d2Timer.C:
			if rand.Intn(2) == 0 {
				c.press(robotgo.KeyD, 1)
			}
			d2Timer.Reset(3 * time.Second)
		case <-d3Timer.C:
			if rand.Intn(2) == 0 {
				c.press(robotgo.KeyR, 1)
			}
			d3Timer.Reset(8 * time.Second)
		case <-d4Timer.C:
			if rand.Intn(2) == 0 {
				c.press(robotgo.KeyW, 1)
			}
			d4Timer.Reset(10 * time.Second)
		case <-d5Timer.C:
			if rand.Intn(2) == 0 {
				c.press(robotgo.KeyE, 1)
			}
			d5Timer.Reset(13 * time.Second)
		//case <-d6Timer.C:
		//	if rand.Intn(2) == 0 {
		//		c.press(robotgo.KeyT, 4)
		//	}
		//	gpTimer.Reset(60 * time.Second)
		//	d6Timer.Reset(60 * time.Second)
		case <-gpTimer.C:
			if rand.Intn(2) == 0 {
				c.press(robotgo.KeyQ, 4)
			}
			gpTimer.Reset(60 * time.Second)
			//d6Timer.Reset(60 * time.Second)
		case <-rTimer.C:
			if rand.Intn(2) == 0 {
				c.press(robotgo.F1, 1)
			}
			rTimer.Reset(15 * time.Second)
		case <-bTimer.C:
			if rand.Intn(2) == 0 {
				c.press(robotgo.F2, 1)
			}
			bTimer.Reset(15 * time.Second)
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}
