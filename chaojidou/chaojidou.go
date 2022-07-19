package chaojidou

import (
	"errors"
	"fmt"
	"github.com/gangcheng1030/game_script/utils"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"github.com/shirou/gopsutil/v3/process"
	"math/rand"
	"strings"
	"sync"
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

type ZhuiSuType string

const (
	ZHUISU_TYPE_GUTU         = "gt"
	ZHUISU_TYPE_DADUHUI      = "ddh"
	ZHUISU_TYPE_JIUYUNDONG   = "jyd"
	ZHUISU_TYPE_TONGHUAZHEN  = "thz"
	ZHUISU_TYPE_KAERJIAYIZHI = "kejyz"
	ZHUISU_TYPE_GELAXIYA     = "glxy"
	ZHUISU_TYPE_BULINDIXI    = "bldx"
	ZHUISU_TYPE_LALAIYE      = "lly"
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
	SignIn(account string, password string)
	SelectRole(n int, first bool, fullScreenMode int)
	RepairEquipment()
	ClearBag(sellToDaiXi bool)
	CardsUp()
	CreateGroup(followerFunTuanNames []string, first bool)
	CardsDown()
	QuitRole() // 返回到角色选择界面
	Quit()     // 退出游戏
	// LeftClick x: 横坐标, y: 纵坐标, w: 窗口宽, h: 窗口高
	LeftClick(x, y, w, h int) bool
	JuQing(jt JuQingType)
	JiuYunDong(dt DifficultyType)
	ZhuiSu(zt ZhuiSuType, dt DifficultyType)
	MeiRiTiaoZhan(mt MeiRiType, dt DifficultyType)
	LiuLangTuan(lt LiuLangTuanType, dt DifficultyType)
	AoDeSai(dt DifficultyType)
	JinBen(jt JinBenType, dt DifficultyType, times int)
	JiZhanYanSuan()

	GetGameWindow() robotgo.Rect
	GetPid() int32

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

type FuBenOption struct {
	ExcludeM bool
}

type chaoJiDou struct {
	Pid int32

	// 通用
	AccountBox       robotgo.Rect
	PasswordBox      robotgo.Rect
	SignInButton     robotgo.Rect
	GameWindow       robotgo.Rect
	RoleMap          RoleMap
	ShouChongMap     ShouChongMap
	ActivityMap      ActivityMap
	BigMap           BigMap
	MenuMap          MenuMap
	RenWuMap         RenWuMap
	JuQingMap        JuQingMap
	ZhuiSuMap        ZhuiSuMap
	MeiRiMap         MeiRiMap
	LiuLangTuanMap   LiuLangTuanMap
	AoDeSaiMap       AoDeSaiMap
	JinBenMap        JinBenMap
	JiZhanYanSuanMap JiZhanYanSuanMap

	StartBattleButton robotgo.Rect
	EnterButton       robotgo.Rect // 入场按钮
	EnterDButton      robotgo.Rect // 推荐难度，确认入场按钮
	EnterSButton      robotgo.Rect // 分数不够，确认入场按钮
	EnterSButton2     robotgo.Rect // 分数不够2，确认入场按钮

	// 队员
	GroupAcceptButton robotgo.Rect
	EnterAcceptButton robotgo.Rect

	// 队长
}

type RoleMap struct {
	PageButtonsOrigin []robotgo.Rect
	RoleButtonsOrigin []robotgo.Rect

	PageButtons []robotgo.Rect
	RoleButtons []robotgo.Rect
}

type ShouChongMap struct {
	CloseButtonOrigin robotgo.Rect
}

type ActivityMap struct {
	CloseButtonOrigin robotgo.Rect
}

type BigMap struct {
	ZhuiSu          robotgo.Rect
	MeiRi           robotgo.Rect
	LiuLangTuan     robotgo.Rect
	Hermosi         robotgo.Rect
	CiYuanChuanSong robotgo.Rect
	BeiYi           robotgo.Rect

	ZhuangBeiFenJie robotgo.Rect
	ShangDian       robotgo.Rect
	DaiXi           robotgo.Rect
}

type MenuMap struct {
	BaseSettingsButtonOrigin robotgo.Rect
	SelectRoleButton         robotgo.Rect
	QuitButton               robotgo.Rect
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

	FuBens     map[ZhuiSuType]FuBen
	JiuYunDong FuBen
}

type MeiRiMap struct {
	Window              robotgo.Rect
	DifficultyTypePoses []robotgo.Rect

	ShenHaiZhuKe      []JuQingType
	ZuiQiangJianShi   []JuQingType
	TieBiShouWei      []JuQingType
	YanHua            []JuQingType
	BingFengWangGuo   []JuQingType
	XiaoXinChuDian    []JuQingType
	RenXiaoGuiDa      []JuQingType
	JiuYunZhiDian     []JuQingType
	SaRanZhiHua       []JuQingType
	ShengWuShiYan     []JuQingType
	QiangZhanDaSha    []JuQingType
	HeWeiAnYing       []JuQingType
	WeiXianHuiZhuan   []JuQingType
	WangMingTuDeDaoLu []JuQingType
	ShenMiDongWuXue   []JuQingType
	HuoZhiJiDian      []JuQingType
	ShengDiXunLi      []JuQingType
}

type LiuLangTuanMap struct {
	EnterButton robotgo.Rect

	FuBenArray []FuBen
}

type AoDeSaiMap struct {
	AoDeSaiFuBen FuBen
}

type JinBenMap struct {
	EnterButton     robotgo.Rect
	ChuanShaoButton robotgo.Rect

	FuBenArray []FuBen
}

type JiZhanYanSuanMap struct {
	JiZhanYanSuanFuBen FuBen
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

func (c *chaoJiDou) SignIn(account string, password string) {
	robotgo.MoveSmooth(904, 646, mouseSpeedX, mouseSpeedY)
	robotgo.Click()
	robotgo.Sleep(4)
	c.press(robotgo.Esc, 4)

	c.clickButton(c.AccountBox, 2)
	c.clickButton(c.AccountBox, 2)
	for i := 0; i < 40; i++ {
		robotgo.KeyPress(robotgo.Backspace)
	}
	robotgo.Sleep(2)
	robotgo.TypeStr(account)
	robotgo.Sleep(3)

	c.clickButton(c.PasswordBox, 2)
	c.clickButton(c.PasswordBox, 2)
	robotgo.TypeStr(password)
	robotgo.Sleep(1)

	c.clickButton(c.SignInButton, 45)
}

func (c *chaoJiDou) SelectRole(n int, first bool, fullScreenMode int) {
	if first {
		pageNum := n / 6
		c.clickButton(c.RoleMap.PageButtonsOrigin[pageNum], 5)

		roleNum := n % 6
		c.clickButton(c.RoleMap.RoleButtonsOrigin[roleNum], 3)

		c.press(robotgo.KeyF, 80)

		// 关闭首充窗口
		c.clickButton(c.ShouChongMap.CloseButtonOrigin, 4)

		// 关闭活动窗口
		c.clickButton(c.ActivityMap.CloseButtonOrigin, 3)

		// 调整窗口位置和大小
		c.press(robotgo.F8, 3)
		robotgo.MoveSmooth(942, 349, mouseSpeedX, mouseSpeedY)
		robotgo.Click()
		robotgo.Sleep(2)
		robotgo.MoveSmooth(926, 390, mouseSpeedX, mouseSpeedY)
		robotgo.Click()
		robotgo.Sleep(2)
		robotgo.MoveSmooth(960, 723, mouseSpeedX, mouseSpeedY)
		robotgo.Click()
		robotgo.Sleep(8)
		c.press(robotgo.KeyF, 5)

		if fullScreenMode == 1 {
			robotgo.MoveSmooth(933, 273, mouseSpeedX, mouseSpeedY)
			robotgo.Click()
			robotgo.Sleep(2)
			robotgo.MoveSmooth(931, 302, mouseSpeedX, mouseSpeedY)
			robotgo.Click()
			robotgo.Sleep(2)
			robotgo.MoveSmooth(934, 325, mouseSpeedX, mouseSpeedY)
			robotgo.Click()
			robotgo.Sleep(2)
			robotgo.MoveSmooth(911, 436, mouseSpeedX, mouseSpeedY)
			robotgo.Click()
			robotgo.Sleep(2)
			robotgo.MoveSmooth(974, 797, mouseSpeedX, mouseSpeedY)
			robotgo.Click()
			robotgo.Sleep(8)
		} else {
			robotgo.MoveSmooth(569, 166, mouseSpeedX, mouseSpeedY)
			robotgo.Click()
			robotgo.Sleep(2)
			robotgo.MoveSmooth(569, 186, mouseSpeedX, mouseSpeedY)
			robotgo.Click()
			robotgo.Sleep(2)
			robotgo.MoveSmooth(569, 197, mouseSpeedX, mouseSpeedY)
			robotgo.Click()
			robotgo.Sleep(2)
			robotgo.MoveSmooth(544, 249, mouseSpeedX, mouseSpeedY)
			robotgo.Click()
			robotgo.Sleep(2)
			robotgo.MoveSmooth(585, 489, mouseSpeedX, mouseSpeedY)
			robotgo.Click()
			robotgo.Sleep(8)
		}

		// 调整移动方式
		robotgo.MoveSmooth(621, 134, mouseSpeedX, mouseSpeedY)
		robotgo.Click()
		robotgo.Sleep(2)
		robotgo.MoveSmooth(658, 225, mouseSpeedX, mouseSpeedY)
		robotgo.Click()
		robotgo.Sleep(2)
		robotgo.MoveSmooth(679, 565, mouseSpeedX, mouseSpeedY)
		robotgo.Click()
		robotgo.Sleep(2)

		c.press(robotgo.F8, 3)
	} else {
		pageNum := n / 6
		c.clickButton(c.RoleMap.PageButtons[pageNum], 5)

		roleNum := n % 6
		c.clickButton(c.RoleMap.RoleButtons[roleNum], 3)

		c.press(robotgo.KeyF, 45)
	}
}

func (c *chaoJiDou) RepairEquipment() {
	// 修理装备
	robotgo.MoveSmooth(999, 612, mouseSpeedX, mouseSpeedY)
	robotgo.Click()
	robotgo.Sleep(3)
	c.press(robotgo.KeyF, 3)
}

func (c *chaoJiDou) ClearBag(sellToDaiXi bool) {
	// 排序背包
	c.press(robotgo.KeyI, 3)
	robotgo.MoveSmooth(1307, 457, mouseSpeedX, mouseSpeedY)
	robotgo.Click()
	robotgo.Sleep(3)
	robotgo.MoveSmooth(1330, 102, mouseSpeedX, mouseSpeedY)
	robotgo.Click()
	robotgo.Sleep(3)

	// 分解装备
	c.press(robotgo.KeyM, 3)
	c.clickButton(c.BigMap.ZhuangBeiFenJie, NpcWaitSecs)
	robotgo.MoveSmooth(875, 424, mouseSpeedX, mouseSpeedY)
	robotgo.Click()
	robotgo.Sleep(3)
	robotgo.MoveSmooth(724, 582, mouseSpeedX, mouseSpeedY)
	robotgo.Click()
	robotgo.Sleep(3)
	c.press(robotgo.KeyF, 4)
	c.press(robotgo.KeyF, 3)
	robotgo.MoveSmooth(1292, 737, mouseSpeedX, mouseSpeedY)
	robotgo.Click()
	robotgo.Sleep(3)

	// 卖装备
	c.press(robotgo.KeyM, 3)
	if sellToDaiXi {
		c.clickButton(c.BigMap.DaiXi, 15)
		c.press(robotgo.KeyF, 3)
	} else {
		c.clickButton(c.BigMap.ShangDian, 5)
	}
	robotgo.KeyDown(robotgo.Shift)
	startX := 982
	startY := 491
	width := 37
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if i == 0 && j < 5 {
				continue
			}
			robotgo.MoveSmooth(startX+j*width, startY+i*width, mouseSpeedX, mouseSpeedY)
			robotgo.Click("right")
			robotgo.MilliSleep(400)
		}
	}
	robotgo.KeyUp(robotgo.Shift)
	robotgo.Sleep(3)
	robotgo.MoveSmooth(1292, 737, mouseSpeedX, mouseSpeedY)
	robotgo.Click()
	robotgo.Sleep(3)
}

func (c *chaoJiDou) CardsUp() {
	c.press(robotgo.KeyJ, 3)
	robotgo.MoveSmooth(430, 150, mouseSpeedX, mouseSpeedY)
	robotgo.Click("right")
	robotgo.Sleep(1)
	robotgo.MoveSmooth(510, 150, mouseSpeedX, mouseSpeedY)
	robotgo.Click("right")
	robotgo.Sleep(1)
	robotgo.MoveSmooth(590, 150, mouseSpeedX, mouseSpeedY)
	robotgo.Click("right")
	robotgo.Sleep(1)
	robotgo.MoveSmooth(670, 150, mouseSpeedX, mouseSpeedY)
	robotgo.Click("right")
	robotgo.Sleep(1)
	robotgo.MoveSmooth(750, 150, mouseSpeedX, mouseSpeedY)
	robotgo.Click("right")
	robotgo.Sleep(1)
	robotgo.MoveSmooth(430, 230, mouseSpeedX, mouseSpeedY)
	robotgo.Click("right")
	robotgo.Sleep(1)
	robotgo.MoveSmooth(510, 230, mouseSpeedX, mouseSpeedY)
	robotgo.Click("right")
	robotgo.Sleep(1)
	robotgo.MoveSmooth(590, 230, mouseSpeedX, mouseSpeedY)
	robotgo.Click("right")
	robotgo.Sleep(1)
	c.press(robotgo.KeyJ, 3)
}

func (c *chaoJiDou) CreateGroup(followerFunTuanNames []string, first bool) {
	c.press(robotgo.KeyP, 3)
	robotgo.MoveSmooth(621, 561, mouseSpeedX, mouseSpeedY)
	robotgo.Click()
	robotgo.Sleep(2)
	if first {
		robotgo.MoveSmooth(158, 461, mouseSpeedX, mouseSpeedY)
		robotgo.Click()
		robotgo.Sleep(2)

		robotgo.MoveSmooth(163, 481, mouseSpeedX, mouseSpeedY)
		robotgo.Click()
		robotgo.Sleep(2)
	}
	robotgo.MoveSmooth(172, 522, mouseSpeedX, mouseSpeedY)
	robotgo.Click()
	robotgo.Sleep(2)

	rects := []robotgo.Rect{
		{
			Point: robotgo.Point{X: 611, Y: 502},
			Size:  robotgo.Size{W: 4, H: 2},
		},
		{
			Point: robotgo.Point{X: 745, Y: 502},
			Size:  robotgo.Size{W: 4, H: 2},
		},
		{
			Point: robotgo.Point{X: 882, Y: 502},
			Size:  robotgo.Size{W: 4, H: 2},
		},
	}
	c.clickButton(rects[0], 2)

	for i := range followerFunTuanNames {
		robotgo.MoveSmooth(1001, 569, mouseSpeedX, mouseSpeedY)
		robotgo.Click()
		robotgo.Sleep(1)
		robotgo.MoveSmooth(1001, 569, mouseSpeedX, mouseSpeedY)
		robotgo.Click()
		robotgo.Sleep(1)

		robotgo.TypeStr(followerFunTuanNames[i])
		robotgo.Sleep(1)

		robotgo.MoveSmooth(1152, 567, mouseSpeedX, mouseSpeedY)
		robotgo.Click()
		robotgo.Sleep(5)

		tmpPoint := utils.GetRandomPointInRect(c.GroupAcceptButton)
		e := &hook.Event{
			Kind:   hook.MouseDown,
			X:      int16(tmpPoint.X),
			Y:      int16(tmpPoint.Y),
			Clicks: 1,
		}
		err := SendEvent(Follwers[i], e)
		if err != nil {
			panic(err)
		}
		robotgo.Sleep(2)
	}

	c.press(robotgo.KeyP, 3)
}

func (c *chaoJiDou) CardsDown() {
	c.press(robotgo.KeyJ, 3)
	robotgo.MoveSmooth(186, 240, mouseSpeedX, mouseSpeedY)
	robotgo.Click("right")
	robotgo.Sleep(1)
	robotgo.MoveSmooth(100, 328, mouseSpeedX, mouseSpeedY)
	robotgo.Click("right")
	robotgo.Sleep(1)
	robotgo.MoveSmooth(270, 325, mouseSpeedX, mouseSpeedY)
	robotgo.Click("right")
	robotgo.Sleep(1)
	robotgo.MoveSmooth(185, 415, mouseSpeedX, mouseSpeedY)
	robotgo.Click("right")
	robotgo.Sleep(1)
	c.press(robotgo.KeyJ, 3)
}

func (c *chaoJiDou) QuitRole() {
	c.press(robotgo.Esc, 3)
	c.clickButton(c.MenuMap.SelectRoleButton, 3)
	c.press(robotgo.KeyF, 15)
}

func (c *chaoJiDou) Quit() {
	c.press(robotgo.Esc, 3)
	c.clickButton(c.MenuMap.QuitButton, 3)
	c.press(robotgo.KeyF, 15)

	// 退出启动器
	pses, _ := process.Processes()
	for _, pss := range pses {
		name, _ := pss.Name()
		if strings.HasPrefix(name, "LootHoarder.exe") {
			pss.Terminate()
		}
	}
	robotgo.Sleep(15)
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

func (c *chaoJiDou) ZhuiSu(zt ZhuiSuType, dt DifficultyType) {
	c.press(robotgo.KeyM, 1)

	// 大地图点追溯
	c.clickButton(c.BigMap.ZhuiSu, NpcWaitSecs)

	// 追溯地图点击相应副本
	robotgo.KeyPress(robotgo.KeyF)
	robotgo.Sleep(2)
	robotgo.KeyPress(robotgo.KeyF)
	robotgo.Sleep(3)
	c.clickButton(c.ZhuiSuMap.FuBens[zt].Window, 3)

	// 选难度
	c.clickButton(c.ZhuiSuMap.FuBens[zt].DifficultyTypePoses[dt], 2)

	// 入场
	c.clickButton(c.ZhuiSuMap.EnterButton, 3)
	c.clickButton(c.EnterDButton, 1) // 这个入场老是出bug，所以点2次
	c.clickButton(c.EnterDButton, 3)
	c.clickButton(c.EnterSButton, 1) // 这个入场老是出bug，所以点2次
	c.clickButton(c.EnterSButton, 3)
	c.clickButton(c.EnterSButton2, 1) // 这个入场老是出bug，所以点2次
	c.clickButton(c.EnterSButton2, 1)
	c.handleFollowersClick(c.EnterAcceptButton, 2, 5, 5000, 0)
	robotgo.Sleep(ReadMapWaitSecs)

	// 打怪
	if zt == ZHUISU_TYPE_JIUYUNDONG {
		c.jiuYunDongHelper()
	} else if zt == ZHUISU_TYPE_GUTU {
		c.guTuHelper()
	} else if zt == ZHUISU_TYPE_DADUHUI {
		c.daDuHuiHelper()
	} else if zt == ZHUISU_TYPE_TONGHUAZHEN {
		c.tongHuaZhenHelper()
	} else if zt == ZHUISU_TYPE_KAERJIAYIZHI {
		c.kaErJiaYiZhiHelper()
	} else if zt == ZHUISU_TYPE_GELAXIYA {
		c.geLaXiYaHelper()
	} else if zt == ZHUISU_TYPE_BULINDIXI {
		c.buLinDiXiHelper()
	} else if zt == ZHUISU_TYPE_LALAIYE {
		c.laLaiYeHelper()
	} else {
		fmt.Println("invalid zhuisutype.")
		return
	}

	// 返回主城
	robotgo.KeyPress(robotgo.F12)
	robotgo.Sleep(7)
	robotgo.KeyPress(robotgo.KeyF)
	robotgo.Sleep(4)
	robotgo.KeyPress(robotgo.KeyF)
	c.handleFollowersPress(robotgo.KeyF, 5)
	robotgo.Sleep(25)
	c.handleFollowersMove(100, 100, 2, 0, 1000)
	c.move(246, 218, 2, 6)
}

func (c *chaoJiDou) JiuYunDong(dt DifficultyType) {
	c.press(robotgo.KeyM, 1)

	// 大地图点追溯
	c.clickButton(c.BigMap.ZhuiSu, NpcWaitSecs)

	// 追溯地图点击‘九云洞’
	robotgo.KeyPress(robotgo.KeyF)
	robotgo.Sleep(2)
	robotgo.KeyPress(robotgo.KeyF)
	robotgo.Sleep(3)
	c.clickButton(c.ZhuiSuMap.JiuYunDong.Window, 3)

	// 选难度
	c.clickButton(c.ZhuiSuMap.JiuYunDong.DifficultyTypePoses[dt], 2)

	// 入场
	c.clickButton(c.ZhuiSuMap.EnterButton, 3)
	c.clickButton(c.EnterDButton, 1) // 这个入场老是出bug，所以点2次
	c.clickButton(c.EnterDButton, 3)
	c.clickButton(c.EnterSButton, 1) // 这个入场老是出bug，所以点2次
	c.clickButton(c.EnterSButton, 3)
	c.clickButton(c.EnterSButton2, 1) // 这个入场老是出bug，所以点2次
	c.clickButton(c.EnterSButton2, 1)

	c.handleFollowersClick(c.EnterAcceptButton, 2, 5, 5000, 0)
	robotgo.Sleep(ReadMapWaitSecs)

	// 打怪
	c.jiuYunDongHelper()

	// 返回主城
	robotgo.KeyPress(robotgo.F12)
	robotgo.Sleep(5)
	robotgo.KeyPress(robotgo.KeyF)
	robotgo.Sleep(3)
	robotgo.KeyPress(robotgo.KeyF)
	c.handleFollowersPress(robotgo.KeyF, 5)
	robotgo.Sleep(25)
	c.handleFollowersMove(100, 100, 2, 0, 1000)
	c.move(246, 218, 2, 6)
}

func (c *chaoJiDou) guTuHelper() {

}

func (c *chaoJiDou) daDuHuiHelper() {
	// 开始战斗
	c.clickButton(c.StartBattleButton, 1)
	c.handleFollowersPress(robotgo.KeyF, 5)
	robotgo.Sleep(10)

	// 第1张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_DADUHUI].SmallMap[0], 1, 0, 3000, 0)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_DADUHUI].SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(568, 48, 3, 4)
	c.press(robotgo.KeyD, 1)
	c.prestart()

	// 第2张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_DADUHUI].SmallMap[1], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_DADUHUI].SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(216, 102, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_DADUHUI].SmallMap[2], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_DADUHUI].SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1031, 157, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_DADUHUI].SmallMap[3], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_DADUHUI].SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(442, 47, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_DADUHUI].SmallMap[4], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_DADUHUI].SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(303, 130, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第6张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_DADUHUI].SmallMap[5], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_DADUHUI].SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(612, 54, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第7张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_DADUHUI].SmallMap[6], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_DADUHUI].SmallMap[6], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1034, 121, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第8张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_DADUHUI].SmallMap[7], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_DADUHUI].SmallMap[7], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	robotgo.Sleep(30)
}

func (c *chaoJiDou) tongHuaZhenHelper() {
	// 开始战斗
	c.clickButton(c.StartBattleButton, 1)
	c.handleFollowersPress(robotgo.KeyF, 5)
	robotgo.Sleep(10)

	// 第1张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_TONGHUAZHEN].SmallMap[0], 1, 0, 3000, 0)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_TONGHUAZHEN].SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1090, 205, 3, 4)
	c.press(robotgo.KeyD, 1)
	c.prestart()

	// 第2张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_TONGHUAZHEN].SmallMap[1], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_TONGHUAZHEN].SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1027, 154, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_TONGHUAZHEN].SmallMap[2], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_TONGHUAZHEN].SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1138, 154, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.move(952, 140, 3, 5)

	// 第4张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_TONGHUAZHEN].SmallMap[3], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_TONGHUAZHEN].SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1165, 191, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_TONGHUAZHEN].SmallMap[4], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_TONGHUAZHEN].SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1037, 125, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第6张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_TONGHUAZHEN].SmallMap[5], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_TONGHUAZHEN].SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1002, 75, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第7张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_TONGHUAZHEN].SmallMap[6], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_TONGHUAZHEN].SmallMap[6], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(303, 200, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第8张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_TONGHUAZHEN].SmallMap[7], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_TONGHUAZHEN].SmallMap[7], 8)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	robotgo.Sleep(30)
}

func (c *chaoJiDou) kaErJiaYiZhiHelper() {
	// 开始战斗
	c.clickButton(c.StartBattleButton, 1)
	c.handleFollowersPress(robotgo.KeyF, 5)
	robotgo.Sleep(10)

	// 第1张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_KAERJIAYIZHI].SmallMap[0], 1, 0, 3000, 0)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_KAERJIAYIZHI].SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1015, 175, 3, 4)
	c.press(robotgo.KeyD, 1)
	c.prestart()

	// 第2张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_KAERJIAYIZHI].SmallMap[1], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_KAERJIAYIZHI].SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1140, 93, 3, 4)
	c.press(robotgo.KeyD, 1)
	c.move(1006, 225, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	for i := 0; i < len(Follwers); i += 3 {
		c.press(robotgo.KeyD, 4)
	}

	// 第3张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_KAERJIAYIZHI].SmallMap[2], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_KAERJIAYIZHI].SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(404, 195, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	for i := 0; i < len(Follwers); i += 3 {
		c.press(robotgo.KeyD, 4)
	}

	// 第4张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_KAERJIAYIZHI].SmallMap[3], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_KAERJIAYIZHI].SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(711, 50, 3, 4)
	c.press(robotgo.KeyD, 1)
	c.move(256, 72, 3, 4)
	c.press(robotgo.KeyD, 1)
	c.move(233, 311, 3, 4)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_KAERJIAYIZHI].SmallMap[4], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_KAERJIAYIZHI].SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1110, 71, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	for i := 0; i < len(Follwers); i += 3 {
		c.press(robotgo.KeyD, 4)
	}

	// 第6张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_KAERJIAYIZHI].SmallMap[5], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_KAERJIAYIZHI].SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1096, 83, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	for i := 0; i < len(Follwers); i += 3 {
		c.press(robotgo.KeyD, 4)
	}

	// 第7张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_KAERJIAYIZHI].SmallMap[6], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_KAERJIAYIZHI].SmallMap[6], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(355, 202, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	for i := 0; i < len(Follwers); i += 3 {
		c.press(robotgo.KeyD, 4)
	}

	// 第8张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_KAERJIAYIZHI].SmallMap[7], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_KAERJIAYIZHI].SmallMap[7], 8)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	robotgo.Sleep(30)
}

func (c *chaoJiDou) geLaXiYaHelper() {
	// 开始战斗
	c.clickButton(c.StartBattleButton, 1)
	c.handleFollowersPress(robotgo.KeyF, 5)
	robotgo.Sleep(10)

	// 第1张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_GELAXIYA].SmallMap[0], 1, 0, 3000, 0)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_GELAXIYA].SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1002, 146, 10, 4)
	c.press(robotgo.KeyD, 1)
	c.prestart()

	// 第2张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_GELAXIYA].SmallMap[1], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_GELAXIYA].SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1319, 162, 10, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	for i := 0; i < len(Follwers); i += 3 {
		c.press(robotgo.KeyD, 4)
	}

	// 第3张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_GELAXIYA].SmallMap[2], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_GELAXIYA].SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(78, 332, 10, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	for i := 0; i < len(Follwers); i += 3 {
		c.press(robotgo.KeyD, 4)
	}

	// 第4张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_GELAXIYA].SmallMap[3], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_GELAXIYA].SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(110, 386, 10, 3)
	robotgo.MoveSmooth(110, 286, mouseSpeedX, mouseSpeedY)
	c.press(robotgo.KeyD, 1)
	c.move(110, 386, 3, 3)
	robotgo.MoveSmooth(110, 286, mouseSpeedX, mouseSpeedY)
	c.press(robotgo.KeyD, 1)
	c.move(110, 386, 3, 3)
	robotgo.MoveSmooth(110, 286, mouseSpeedX, mouseSpeedY)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_GELAXIYA].SmallMap[4], 1, 0, 3000, 6)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_GELAXIYA].SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(60, 347, 10, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	for i := 0; i < len(Follwers); i += 3 {
		c.press(robotgo.KeyD, 4)
	}

	// 第6张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_GELAXIYA].SmallMap[5], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_GELAXIYA].SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	for i := 0; i < len(Follwers); i += 3 {
		c.press(robotgo.KeyD, 4)
	}

	// 第7张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_GELAXIYA].SmallMap[6], 1, 0, 3000, 0)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_GELAXIYA].SmallMap[6], 9)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	robotgo.Sleep(30)
}

func (c *chaoJiDou) buLinDiXiHelper() {
	// 开始战斗
	c.clickButton(c.StartBattleButton, 1)
	c.handleFollowersPress(robotgo.KeyF, 5)
	robotgo.Sleep(10)

	// 第1张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_BULINDIXI].SmallMap[0], 1, 0, 3000, 0)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_BULINDIXI].SmallMap[0], 6)
	c.press(robotgo.Key3, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyS, 3)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyE, 3)
	robotgo.Sleep(7)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyW, 1)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)

	// 第2张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_BULINDIXI].SmallMap[1], 1, 0, 3000, 0)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_BULINDIXI].SmallMap[1], 8)
	robotgo.MoveSmooth(321, 351, mouseSpeedX, mouseSpeedY)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyD, 1)
	for i := 0; i < len(Follwers); i += 3 {
		c.press(robotgo.KeyD, 4)
	}
	robotgo.Sleep(6)

	// 第3张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_BULINDIXI].SmallMap[2], 1, 0, 3000, 0)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_BULINDIXI].SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(292, 213, 2, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	for i := 0; i < len(Follwers); i += 3 {
		c.press(robotgo.KeyD, 4)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() { // 为了防止卡死，需要提前走位
		defer wg.Done()
		c.handleFollowersMove(292, 213, 2, 0, 0)
		robotgo.Sleep(4)
		for i := 0; i < 10; i++ {
			c.handleFollowersMove(272, 216, 1, 0, 0)
			robotgo.Sleep(1)
		}
	}()
	c.multiMove(272, 216, 1, 1, 10)

	// 第4张怪物图
	wg.Wait()
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_BULINDIXI].SmallMap[3], 1, 0, 3000, 0)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_BULINDIXI].SmallMap[3], 6)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(82, 58, 2, 6)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.refreshD4()

	// 第5张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_BULINDIXI].SmallMap[4], 1, 0, 3000, 6)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_BULINDIXI].SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	robotgo.MoveSmooth(394, 173, mouseSpeedX, mouseSpeedY)
	c.press(robotgo.Key1, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)
	for i := 0; i < len(Follwers); i += 3 {
		c.press(robotgo.KeyD, 4)
	}

	// 第6张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_BULINDIXI].SmallMap[5], 1, 0, 3000, 0)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_BULINDIXI].SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(35, 618, 1, 1, 3)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	for i := 0; i < len(Follwers); i += 3 {
		c.press(robotgo.KeyD, 4)
	}

	// 第7张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_BULINDIXI].SmallMap[6], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_BULINDIXI].SmallMap[6], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1101, 84, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	for i := 0; i < len(Follwers); i += 3 {
		c.press(robotgo.KeyD, 4)
	}

	// 第8张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_BULINDIXI].SmallMap[7], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_BULINDIXI].SmallMap[7], 7)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 4)
	c.continuedBattle(20)
	robotgo.Sleep(10)
}

func (c *chaoJiDou) laLaiYeHelper() {
	// 开始战斗
	c.clickButton(c.StartBattleButton, 1)
	c.handleFollowersPress(robotgo.KeyF, 5)
	robotgo.Sleep(10)

	// 第1张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_LALAIYE].SmallMap[0], 1, 0, 3000, 0)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_LALAIYE].SmallMap[0], 6)
	c.press(robotgo.Key3, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyS, 3)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyE, 3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(7)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyW, 1)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)

	// 第2张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_LALAIYE].SmallMap[1], 1, 0, 3000, 0)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_LALAIYE].SmallMap[1], 10)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	c.multiMove(1080, 361, 1, 1, 2)
	robotgo.Sleep(2)
	robotgo.MoveSmooth(1080, 500, mouseSpeedX, mouseSpeedY)
	c.press(robotgo.KeyD, 1)
	for i := 0; i < len(Follwers); i += 3 {
		c.press(robotgo.KeyD, 4)
	}
	robotgo.Sleep(6)

	// 第3张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_LALAIYE].SmallMap[2], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_LALAIYE].SmallMap[2], 6)
	c.press(robotgo.KeyQ, 1)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	for i := 0; i < len(Follwers); i += 3 {
		c.press(robotgo.KeyD, 4)
	}

	// 第4张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_LALAIYE].SmallMap[3], 1, 0, 3000, 0)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_LALAIYE].SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(268, 128, 2, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.refreshD4()

	// 第5张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_LALAIYE].SmallMap[4], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_LALAIYE].SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1178, 238, 5, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.move(1110, 548, 5, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	for i := 0; i < len(Follwers); i += 3 {
		c.press(robotgo.KeyD, 4)
	}

	// 第6张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_LALAIYE].SmallMap[5], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_LALAIYE].SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1074, 359, 1, 1, 3)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	for i := 0; i < len(Follwers); i += 3 {
		c.press(robotgo.KeyD, 4)
	}

	// 第7张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_LALAIYE].SmallMap[6], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_LALAIYE].SmallMap[6], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1075, 512, 3, 1, 3)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.multiMove(1045, 169, 3, 1, 2)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	for i := 0; i < len(Follwers); i += 3 {
		c.press(robotgo.KeyD, 4)
	}

	// 第8张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_LALAIYE].SmallMap[7], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.FuBens[ZHUISU_TYPE_LALAIYE].SmallMap[7], 8)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 4)
	c.continuedBattle(20)
	robotgo.Sleep(10)
}

func (c *chaoJiDou) jiuYunDongHelper() {
	// 开始战斗
	c.clickButton(c.StartBattleButton, 1)
	c.handleFollowersPress(robotgo.KeyF, 5)
	robotgo.Sleep(10)

	// 第1张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.JiuYunDong.SmallMap[0], 1, 0, 3000, 0)
	c.clickButton(c.ZhuiSuMap.JiuYunDong.SmallMap[0], 10)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	c.move(950, 225, 10, 5)
	c.prestart()

	// 第2张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.JiuYunDong.SmallMap[1], 1, 0, 3000, 3)
	c.clickButton(c.ZhuiSuMap.JiuYunDong.SmallMap[1], 8)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	c.move(1086, 71, 2, 8)
	for i := 0; i < len(Follwers); i += 2 {
		c.press(robotgo.KeyD, 3)
	}

	// 第3张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.JiuYunDong.SmallMap[2], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.JiuYunDong.SmallMap[2], 8)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	c.move(1086, 71, 2, 8)
	for i := 0; i < len(Follwers); i += 2 {
		c.press(robotgo.KeyD, 3)
	}

	// 第4张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.JiuYunDong.SmallMap[3], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.JiuYunDong.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1082, 97, 2, 6)
	c.move(1082, 97, 2, 6)
	for i := 0; i < len(Follwers); i += 2 {
		c.press(robotgo.KeyD, 3)
	}

	// 第5张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.JiuYunDong.SmallMap[4], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.JiuYunDong.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(299, 141, 1, 1, 2)
	robotgo.Sleep(6)
	for i := 0; i < len(Follwers); i += 2 {
		c.press(robotgo.KeyD, 3)
	}

	// 第6张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.JiuYunDong.SmallMap[5], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.JiuYunDong.SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(633, 91, 2, 8)
	for i := 0; i < len(Follwers); i += 2 {
		c.press(robotgo.KeyD, 3)
	}

	// 第7张怪物图
	c.handleFollowersClick(c.ZhuiSuMap.JiuYunDong.SmallMap[6], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.JiuYunDong.SmallMap[6], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1110, 422, 2, 4)
	c.move(1020, 106, 2, 8)
	for i := 0; i < len(Follwers); i += 2 {
		c.press(robotgo.KeyD, 3)
	}

	// 第8张怪物图：boss
	c.handleFollowersClick(c.ZhuiSuMap.JiuYunDong.SmallMap[7], 1, 0, 3000, 4)
	c.clickButton(c.ZhuiSuMap.JiuYunDong.SmallMap[7], 8)
	c.press(robotgo.KeyD, 2)
	c.press(robotgo.KeyD, 2)
	c.press(robotgo.KeyD, 2)
	c.press(robotgo.Key3, 2)
	for i := 0; i < len(Follwers); i += 2 {
		c.press(robotgo.KeyD, 5)
	}
	robotgo.Sleep(25)
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
	c.clickButton(c.EnterDButton, 3)
	c.clickButton(c.EnterSButton, 1) // 这个入场老是出bug，所以点2次
	c.clickButton(c.EnterSButton, 3)
	c.clickButton(c.EnterSButton2, 1) // 这个入场老是出bug，所以点2次
	c.clickButton(c.EnterSButton2, 1)

	c.handleFollowersClick(c.EnterAcceptButton, 2, 5, 5000, 0)
	robotgo.Sleep(ReadMapWaitSecs)

	// 打怪
	switch lt {
	case LIULANGTUAN_TYPE_1:
		c.liuLangTuan1Helper()
	default:
		// do nothing
	}

	// 返回主城
	robotgo.KeyPress(robotgo.F12)
	robotgo.Sleep(7)
	robotgo.KeyPress(robotgo.KeyF)
	robotgo.Sleep(4)
	robotgo.KeyPress(robotgo.KeyF)
	c.handleFollowersPress(robotgo.KeyF, 5)
	robotgo.Sleep(20)
}

func (c *chaoJiDou) liuLangTuan1Helper() {
	// 开始战斗
	c.clickButton(c.StartBattleButton, 1)
	c.handleFollowersPress(robotgo.KeyF, 5)
	robotgo.Sleep(10)

	// 技能连招
	c.move(86, 696, 2, 2)
	c.prestart()
	c.move(839, 331, 2, 2)

	// 第一阶段
	c.move(1263, 49, 2, 5)
	go c.handleFollowersMove(1263, 49, 2, 0, 3000)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	c.continuedBattle(40)
	for i := 0; i < len(Follwers); i++ {
		c.press(robotgo.KeyD, 1)
		c.continuedBattle(6)
	}
	c.press(robotgo.KeyS, 2)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	c.continuedBattle(10)
	c.press(robotgo.F2, 1)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)
	c.press(robotgo.KeyA, 3)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.KeyW, 2)

	// 第二阶段
	go func() {
		for i := 0; i < 5; i++ {
			c.handleFollowersMove(1077, 534, 3, 0, 200)
			robotgo.MilliSleep(800)
		}
		for i := 0; i < 5; i++ {
			c.handleFollowersMove(931, 111, 3, 0, 200)
			robotgo.MilliSleep(800)
		}
	}()
	for i := 0; i < 5; i++ {
		c.move(1077, 534, 1, 1)
	}
	for i := 0; i < 5; i++ {
		c.move(931, 111, 1, 1)
	}
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.KeyW, 1)
	c.continuedBattle(40)
	for i := 0; i < len(Follwers); i++ {
		c.press(robotgo.KeyD, 1)
		c.continuedBattle(6)
	}
	c.press(robotgo.KeyS, 2)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	c.continuedBattle(10)
	c.press(robotgo.F2, 1)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)
	c.press(robotgo.KeyA, 3)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.KeyW, 2)

	// 第三阶段
	go func() {
		for i := 0; i < 9; i++ {
			c.handleFollowersMove(260, 711, 3, 0, 200)
			robotgo.MilliSleep(800)
		}
		c.handleFollowersMove(401, 41, 3, 0, 3000)
	}()
	for i := 0; i < 9; i++ {
		c.move(260, 711, 1, 1)
	}
	c.move(401, 41, 2, 4)
	robotgo.MoveSmooth(680, 258, 0.9, 0.9)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyT, 4)
	c.press(robotgo.KeyD, 1)
	c.continuedBattle(50)
	for i := 0; i < len(Follwers); i++ {
		c.press(robotgo.KeyD, 1)
		c.continuedBattle(6)
	}
}

func (c *chaoJiDou) AoDeSai(dt DifficultyType) {
	c.press(robotgo.KeyM, 2)

	// 大地图点次元传送
	c.clickButton(c.BigMap.CiYuanChuanSong, NpcWaitSecs)

	// 对话
	//robotgo.KeyPress(robotgo.KeyF)
	//robotgo.Sleep(4)

	// 选中奥德赛地图
	c.clickButton(c.AoDeSaiMap.AoDeSaiFuBen.Window, 2)

	// 选难度
	c.clickButton(c.AoDeSaiMap.AoDeSaiFuBen.DifficultyTypePoses[dt], 2)

	// 入场
	c.clickButton(c.EnterButton, 2)
	c.clickButton(c.EnterDButton, 1) // 这个入场老是出bug，所以点2次
	c.clickButton(c.EnterDButton, 3)
	c.clickButton(c.EnterSButton, 1) // 这个入场老是出bug，所以点2次
	c.clickButton(c.EnterSButton, 3)
	c.clickButton(c.EnterSButton2, 1) // 这个入场老是出bug，所以点2次
	c.clickButton(c.EnterSButton2, 1)
	c.handleFollowersClick(c.EnterAcceptButton, 2, 5, 5000, 0)
	robotgo.Sleep(ReadMapWaitSecs)

	// 打怪
	c.aoDeSaiHelper()

	// 返回主城
	robotgo.KeyPress(robotgo.F12)
	robotgo.Sleep(5)
	robotgo.KeyPress(robotgo.KeyF)
	robotgo.Sleep(3)
	robotgo.KeyPress(robotgo.KeyF)
	c.handleFollowersPress(robotgo.KeyF, 5)
	robotgo.Sleep(20)
}

func (c *chaoJiDou) aoDeSaiHelper() {
	// 开始战斗
	c.clickButton(c.StartBattleButton, 1)
	c.handleFollowersPress(robotgo.KeyF, 5)
	robotgo.Sleep(11)

	robotgo.MoveSmooth(904, 518, mouseSpeedX, mouseSpeedY)
	c.press(robotgo.KeyS, 12)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyE, 3)
	c.press(robotgo.F2, 1)
	c.press(robotgo.Key3, 3)
	c.press(robotgo.KeyW, 1)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)
	c.move(632, 132, 1, 6)

	robotgo.MoveSmooth(1001, 175, mouseSpeedX, mouseSpeedY)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.KeyQ, 4)
	c.move(1001, 175, 3, 4)
	robotgo.MoveSmooth(650, 340, mouseSpeedX, mouseSpeedY)
	ts := 15 + len(Follwers)*3
	c.continuedBattle(ts)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyT, 4)
	ts = 30 + len(Follwers)*4
	c.continuedBattle(ts)
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
	c.clickButton(c.EnterDButton, 3)
	c.clickButton(c.EnterSButton, 1) // 这个入场老是出bug，所以点2次
	c.clickButton(c.EnterSButton, 3)
	c.clickButton(c.EnterSButton2, 1) // 这个入场老是出bug，所以点2次
	c.clickButton(c.EnterSButton2, 1)
	c.handleFollowersClick(c.EnterAcceptButton, 2, 5, 5000, 0)
	robotgo.Sleep(ReadMapWaitSecs)

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
	c.handleFollowersPress(robotgo.KeyF, 5)
	robotgo.Sleep(20)
}

func (c *chaoJiDou) suXingDeChuanShuoHelper() {
	// 开始战斗
	c.clickButton(c.StartBattleButton, 1)
	c.handleFollowersPress(robotgo.KeyF, 5)
	robotgo.Sleep(7)

	// 第1张怪物图
	c.handleFollowersClick(c.JinBenMap.FuBenArray[0].SmallMap[0], 1, 0, 3000, 0)
	c.clickButtonWithAlt(c.JinBenMap.FuBenArray[0].SmallMap[0], 10)
	c.press(robotgo.KeyD, 1)
	robotgo.MoveSmooth(336, 345, 0.9, 0.9)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyR, 1)
	c.multiMove(336, 345, 2, 1, 3)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	c.move(602, 620, 2, 3)
	c.press(robotgo.KeyS, 12)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyE, 3)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyW, 1)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)

	// 第2张怪物图
	c.handleFollowersClick(c.JinBenMap.FuBenArray[0].SmallMap[1], 1, 0, 3000, 5)
	c.clickButtonWithAlt(c.JinBenMap.FuBenArray[0].SmallMap[1], 6)
	c.press(robotgo.KeyD, 1)
	robotgo.MoveSmooth(325, 314, 0.9, 0.9)
	c.press(robotgo.Key1, 4)
	c.press(robotgo.KeyE, 3)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyW, 1)
	c.multiMove(325, 314, 2, 1, 3)
	robotgo.Sleep(3)

	// 第3张怪物图
	c.handleFollowersClick(c.JinBenMap.FuBenArray[0].SmallMap[2], 1, 0, 3000, 2)
	c.clickButtonWithAlt(c.JinBenMap.FuBenArray[0].SmallMap[2], 6)
	c.press(robotgo.KeyD, 1)
	robotgo.MoveSmooth(1055, 299, 0.9, 0.9)
	c.press(robotgo.Key3, 2)
	c.multiMove(1055, 299, 2, 1, 3)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(2)
	for i := 0; i < len(Follwers); i += 1 {
		robotgo.Sleep(1)
	}
	c.multiMove(685, 614, 1, 1, 3)

	// 第4张怪物图
	c.handleFollowersClick(c.JinBenMap.FuBenArray[0].SmallMap[3], 1, 0, 3000, 2)
	c.clickButtonWithAlt(c.JinBenMap.FuBenArray[0].SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	robotgo.MoveSmooth(957, 200, 0.9, 0.9)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyW, 1)
	c.multiMove(957, 200, 2, 1, 3)
	robotgo.Sleep(3)
	c.refreshD4WithoutSleep()

	// 第5张怪物图
	c.handleFollowersClick(c.JinBenMap.FuBenArray[0].SmallMap[4], 1, 0, 3000, 5)
	c.clickButtonWithAlt(c.JinBenMap.FuBenArray[0].SmallMap[4], 6)
	c.press(robotgo.KeyD, 1)
	robotgo.MoveSmooth(1057, 319, 0.9, 0.9)
	c.multiMove(1057, 319, 2, 1, 3)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	for i := 0; i < len(Follwers); i += 1 {
		robotgo.Sleep(1)
	}
	c.multiMove(1057, 319, 2, 1, 3)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)
	for i := 0; i < len(Follwers); i += 1 {
		robotgo.Sleep(1)
	}

	// 第6张怪物图
	c.handleFollowersClick(c.JinBenMap.FuBenArray[0].SmallMap[5], 1, 0, 3000, 6)
	c.clickButtonWithAlt(c.JinBenMap.FuBenArray[0].SmallMap[5], 6)
	c.press(robotgo.KeyD, 1)
	robotgo.MoveSmooth(841, 496, 0.9, 0.9)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyW, 1)
	c.multiMove(841, 496, 2, 1, 3)
	robotgo.Sleep(3)
	for i := 0; i < len(Follwers); i += 1 {
		robotgo.Sleep(1)
	}

	// 第7张怪物图
	c.handleFollowersClick(c.JinBenMap.FuBenArray[0].SmallMap[6], 1, 0, 3000, 4)
	c.clickButtonWithAlt(c.JinBenMap.FuBenArray[0].SmallMap[6], 4)
	robotgo.MoveSmooth(216, 389, 0.9, 0.9)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
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
	c.handleFollowersClick(c.JinBenMap.FuBenArray[0].SmallMap[7], 1, 0, 3000, 2)
	c.clickButtonWithAlt(c.JinBenMap.FuBenArray[0].SmallMap[7], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	robotgo.MoveSmooth(673, 550, 0.9, 0.9)
	c.multiMove(673, 550, 2, 1, 5)
	c.press(robotgo.KeyD, 5)
	c.press(robotgo.KeyD, 5)
	c.press(robotgo.KeyD, 1)
	c.refreshD4WithoutSleep()

	// 第9张怪物图
	c.handleFollowersClick(c.JinBenMap.FuBenArray[0].SmallMap[8], 1, 0, 3000, 4)
	c.clickButtonWithAlt(c.JinBenMap.FuBenArray[0].SmallMap[8], 6)
	c.press(robotgo.KeyD, 1)
	robotgo.MoveSmooth(942, 242, 0.9, 0.9)
	c.press(robotgo.Key3, 2)
	c.multiMove(942, 242, 2, 1, 3)
	robotgo.Sleep(3)
	for i := 0; i < len(Follwers); i += 1 {
		robotgo.Sleep(1)
	}

	// 第10张怪物图：实际上是回到第8张
	c.handleFollowersClick(c.JinBenMap.FuBenArray[0].SmallMap[9], 1, 0, 3000, 0)
	c.clickButtonWithAlt(c.JinBenMap.FuBenArray[0].SmallMap[9], 8)

	// 第11张怪物图：boss
	c.handleFollowersClick(c.JinBenMap.FuBenArray[0].SmallMap[10], 1, 0, 3000, 1)
	c.clickButtonWithAlt(c.JinBenMap.FuBenArray[0].SmallMap[10], 13)
	robotgo.MoveSmooth(490, 320, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.KeyS, 2)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	c.continuedBattle(20)
	for i := 0; i < len(Follwers); i += 1 {
		c.press(robotgo.KeyD, 3)
	}
	c.move(138, 654, 2, 3)
	robotgo.KeyPress(robotgo.KeyF)
	robotgo.Sleep(2)
}

func (c *chaoJiDou) heiAnQinShiZhiHuanHelper() {
	// 开始战斗
	c.clickButton(c.StartBattleButton, 1)
	c.handleFollowersPress(robotgo.KeyF, 5)
	robotgo.Sleep(7)

	// 第1张怪物图
	c.handleFollowersClick(c.JinBenMap.FuBenArray[1].SmallMap[0], 1, 0, 3000, 0)
	c.clickButtonWithAlt(c.JinBenMap.FuBenArray[1].SmallMap[0], 13)
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
	var wg sync.WaitGroup
	wg.Add(1)
	go func() { // 为了防止卡死，需要提前走位
		defer wg.Done()
		for i := 0; i < 3; i++ {
			c.handleFollowersMove(1090, 345, 2, 0, 0)
			robotgo.Sleep(1)
		}
		robotgo.Sleep(2)
		for i := 0; i < 3; i++ {
			c.handleFollowersMove(956, 121, 2, 0, 0)
			robotgo.Sleep(1)
		}
	}()
	c.press(robotgo.F1, 1)
	c.multiMove(158, 350, 2, 1, 1)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)
	c.press(robotgo.F2, 1)
	c.multiMove(685, 614, 1, 1, 3)
	robotgo.Sleep(2)
	for i := 0; i < len(Follwers); i += 1 {
		robotgo.Sleep(3)
	}

	// 第2张怪物图
	wg.Wait()
	c.handleFollowersClick(c.JinBenMap.FuBenArray[1].SmallMap[1], 1, 0, 3000, 4)
	c.clickButtonWithAlt(c.JinBenMap.FuBenArray[1].SmallMap[1], 8)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	c.multiMove(314, 159, 2, 1, 3)
	robotgo.Sleep(3)
	robotgo.MoveSmooth(204, 159, mouseSpeedX, mouseSpeedY)
	c.press(robotgo.KeyD, 1)
	for i := 0; i < len(Follwers); i += 1 {
		robotgo.Sleep(1)
	}
	c.multiMove(314, 159, 2, 1, 2)
	robotgo.Sleep(2)
	for i := 0; i < len(Follwers); i += 1 {
		robotgo.Sleep(1)
	}

	// 第3张怪物图
	c.handleFollowersClick(c.JinBenMap.FuBenArray[1].SmallMap[2], 1, 0, 3000, 6)
	c.clickButtonWithAlt(c.JinBenMap.FuBenArray[1].SmallMap[2], 4)
	robotgo.MoveSmooth(392, 370, 0.9, 0.9)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(392, 370, 2, 1, 3)
	c.press(robotgo.F1, 1)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)

	// 第4张怪物图
	c.handleFollowersClick(c.JinBenMap.FuBenArray[1].SmallMap[3], 1, 0, 3000, 4)
	c.clickButtonWithAlt(c.JinBenMap.FuBenArray[1].SmallMap[3], 6)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(595, 167, 2, 1, 4)
	c.press(robotgo.KeyD, 1)
	for i := 0; i < len(Follwers); i += 1 {
		robotgo.Sleep(1)
	}
	c.multiMove(135, 197, 2, 1, 3)
	robotgo.Sleep(2)
	for i := 0; i < len(Follwers); i += 1 {
		robotgo.Sleep(1)
	}

	// 第5张怪物图
	c.handleFollowersClick(c.JinBenMap.FuBenArray[1].SmallMap[4], 1, 0, 3000, 7)
	c.clickButtonWithAlt(c.JinBenMap.FuBenArray[1].SmallMap[4], 5)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyW, 1)
	c.multiMove(303, 618, 2, 1, 3)
	robotgo.Sleep(1)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)
	c.multiMove(303, 618, 2, 1, 3)
	robotgo.Sleep(3)
	c.refreshD4WithoutSleep()

	//// 第6张怪物图
	//c.handleFollowersClick(c.JinBenMap.FuBenArray[1].SmallMap[5], 1, 0, 3000, 2)
	//c.clickButtonWithAlt(c.JinBenMap.FuBenArray[1].SmallMap[5], 10)
	//c.press(robotgo.KeyD, 1)
	//c.press(robotgo.Key3, 2)
	//c.multiMove(1121, 490, 2, 1, 3)
	//robotgo.Sleep(2)
	//c.press(robotgo.KeyD, 1)
	//c.press(robotgo.KeyD, 4)
	//c.press(robotgo.KeyD, 3)
	//for i := 0; i < len(Follwers); i += 1 {
	//	robotgo.Sleep(1)
	//}
	//
	//// 第7张怪物图
	//c.handleFollowersClick(c.JinBenMap.FuBenArray[1].SmallMap[6], 1, 0, 3000, 0)
	//c.clickButtonWithAlt(c.JinBenMap.FuBenArray[1].SmallMap[6], 10)

	// 第8张怪物图
	robotgo.Sleep(3)
	c.handleFollowersClick(c.JinBenMap.FuBenArray[1].SmallMap[7], 1, 0, 3000, 0)
	c.clickButtonWithAlt(c.JinBenMap.FuBenArray[1].SmallMap[7], 10)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.F1, 1)
	robotgo.MoveSmooth(433, 143, 0.9, 0.9)
	c.press(robotgo.Key1, 3)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)

	// 第9张怪物图
	c.handleFollowersClick(c.JinBenMap.FuBenArray[1].SmallMap[8], 1, 0, 3000, 0)
	c.clickButtonWithAlt(c.JinBenMap.FuBenArray[1].SmallMap[8], 6)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyW, 1)
	c.multiMove(386, 195, 2, 1, 3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)
	for i := 0; i < len(Follwers); i += 2 {
		c.press(robotgo.KeyD, 4)
	}
	c.press(robotgo.KeyD, 1)
	c.multiMove(386, 195, 2, 1, 2)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	for i := 0; i < len(Follwers); i += 2 {
		c.press(robotgo.KeyD, 4)
	}

	// 第10张怪物图
	c.handleFollowersClick(c.JinBenMap.FuBenArray[1].SmallMap[9], 1, 0, 3000, 6)
	c.clickButtonWithAlt(c.JinBenMap.FuBenArray[1].SmallMap[9], 5)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(942, 205, 2, 1, 5)
	c.press(robotgo.KeyD, 1)
	for i := 0; i < len(Follwers); i += 1 {
		robotgo.Sleep(1)
	}
	c.multiMove(942, 205, 2, 1, 5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)
	for i := 0; i < len(Follwers); i += 1 {
		robotgo.Sleep(1)
	}

	// 第11张怪物图：boss
	c.handleFollowersClick(c.JinBenMap.FuBenArray[1].SmallMap[10], 1, 0, 3000, 10)
	c.press(robotgo.F1, 1)
	c.clickButtonWithAlt(c.JinBenMap.FuBenArray[1].SmallMap[10], 8)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 4)
	c.press(robotgo.KeyW, 1)
	c.press(robotgo.F1, 1)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyE, 4)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyD, 1)
	c.continuedBattle(25)
	for i := 0; i < len(Follwers); i += 1 {
		c.press(robotgo.KeyD, 3)
	}
}

func (c *chaoJiDou) JiZhanYanSuan() {
	c.press(robotgo.KeyM, 2)

	// 大地图点金本
	c.clickButton(c.BigMap.BeiYi, NpcWaitSecs)

	// 对话
	robotgo.KeyPress(robotgo.KeyF)
	robotgo.Sleep(2)
	robotgo.KeyPress(robotgo.KeyF)
	robotgo.Sleep(2)

	// 选中地图
	c.clickButton(c.JiZhanYanSuanMap.JiZhanYanSuanFuBen.Window, 2)

	// 入场
	c.clickButton(c.JinBenMap.EnterButton, 2)

	robotgo.Sleep(ReadMapWaitSecs)

	// 开始战斗
	c.clickButton(c.StartBattleButton, 7)

	c.press(robotgo.KeyS, 2)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyW, 1)
	c.press(robotgo.KeyT, 4)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyE, 3)
	c.press(robotgo.F2, 1)

	c.continuedBattle(60)
	robotgo.Sleep(5)
	robotgo.KeyPress(robotgo.F12)
	robotgo.Sleep(5)
	c.press(robotgo.KeyF, 3)
	robotgo.KeyPress(robotgo.KeyF)
	robotgo.Sleep(20)
}

func (c *chaoJiDou) GetGameWindow() robotgo.Rect {
	return c.GameWindow
}

func (c *chaoJiDou) GetPid() int32 {
	return c.Pid
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

// tm单位：秒
func (c *chaoJiDou) clickButtonWithAlt(button robotgo.Rect, tm int) {
	tmpPoint := utils.GetRandomPointInRect(button)
	robotgo.KeyDown(robotgo.Alt)
	robotgo.MoveSmooth(tmpPoint.X+c.GameWindow.X, tmpPoint.Y+c.GameWindow.Y, mouseSpeedX, mouseSpeedY)
	robotgo.MilliSleep(200)
	robotgo.Click()
	robotgo.KeyUp(robotgo.Alt)
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
	//gpTimer := time.NewTimer(5 * time.Second)
	rTimer := time.NewTimer(5 * time.Second)
	//bTimer := time.NewTimer(5 * time.Second)
	defer func() {
		d1Timer.Stop()
		d2Timer.Stop()
		d3Timer.Stop()
		d4Timer.Stop()
		d5Timer.Stop()
		//d6Timer.Stop()
		//gpTimer.Stop()
		rTimer.Stop()
		//bTimer.Stop()
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
		//case <-gpTimer.C:
		//	if rand.Intn(2) == 0 {
		//		c.press(robotgo.KeyQ, 4)
		//	}
		//	gpTimer.Reset(60 * time.Second)
		//d6Timer.Reset(60 * time.Second)
		case <-rTimer.C:
			if rand.Intn(2) == 0 {
				c.press(robotgo.F1, 1)
			}
			rTimer.Reset(15 * time.Second)
		//case <-bTimer.C:
		//	if rand.Intn(2) == 0 {
		//		c.press(robotgo.F2, 1)
		//	}
		//	bTimer.Reset(15 * time.Second)
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func (c *chaoJiDou) prestart() {
	c.press(robotgo.KeyS, 10)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyE, 3)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyW, 1)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)
}

func (c *chaoJiDou) refreshD4() {
	c.press(robotgo.KeyA, 3)
	c.press(robotgo.KeyS, 2)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.F2, 1)
	robotgo.Sleep(6)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.KeyW, 1)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)
}

func (c *chaoJiDou) refreshD4WithoutSleep() {
	c.press(robotgo.KeyA, 3)
	c.press(robotgo.KeyS, 2)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.F2, 1)
	robotgo.Sleep(3)
	c.press(robotgo.KeyW, 1)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)
}

func (c *chaoJiDou) handleFollowersClick(r robotgo.Rect, clicks int, preSleepSec int, aroundSleepMs int, postSleepSec int) {
	if len(Follwers) == 0 {
		return
	}

	if preSleepSec > 0 {
		robotgo.Sleep(preSleepSec)
	}
	var waitGroup sync.WaitGroup
	for _, addr := range Follwers {
		tmpAddr := addr
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			if aroundSleepMs > 0 {
				st1 := rand.Intn(aroundSleepMs)
				robotgo.MilliSleep(st1)
			}
			tmpPoint := utils.GetRandomPointInRect(r)
			e := &hook.Event{
				Kind:   hook.MouseDown,
				X:      int16(tmpPoint.X),
				Y:      int16(tmpPoint.Y),
				Clicks: uint16(clicks),
			}
			err := SendEvent(tmpAddr, e)
			if err != nil {
				panic(err)
			}
		}()
	}

	waitGroup.Wait()

	if postSleepSec > 0 {
		robotgo.Sleep(postSleepSec)
	}
}

func (c *chaoJiDou) handleFollowersMove(x int, y int, errors int, preSleepSec int, aroundSleepMs int) {
	if len(Follwers) == 0 {
		return
	}

	if preSleepSec > 0 {
		robotgo.Sleep(preSleepSec)
	}
	var waitGroup sync.WaitGroup
	for _, addr := range Follwers {
		tmpAddr := addr
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			if aroundSleepMs > 0 {
				st1 := rand.Intn(aroundSleepMs)
				robotgo.MilliSleep(st1)
			}
			tmpPoint := utils.GetRandomPointInRect1(x-errors, y-errors, 2*errors, 2*errors)
			e := &hook.Event{
				Kind: hook.MouseMove,
				X:    int16(tmpPoint.X),
				Y:    int16(tmpPoint.Y),
			}
			err := SendEvent(tmpAddr, e)
			if err != nil {
				panic(err)
			}
		}()
	}

	waitGroup.Wait()
}

func (c *chaoJiDou) handleFollowersPress(key string, st int) {
	if len(Follwers) == 0 {
		return
	}

	if st > 0 {
		robotgo.Sleep(st)
	}
	var waitGroup sync.WaitGroup
	for _, addr := range Follwers {
		tmpAddr := addr
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			st1 := rand.Intn(5000)
			robotgo.MilliSleep(st1)
			e := &hook.Event{
				Kind:    hook.KeyDown,
				Keychar: []rune(key)[0],
			}
			err := SendEvent(tmpAddr, e)
			if err != nil {
				panic(err)
			}
		}()
	}

	waitGroup.Wait()
}
