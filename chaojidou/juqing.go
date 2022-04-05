package chaojidou

import "C"
import "github.com/go-vgo/robotgo"

type JuQingType string

const (
	JUQING_TYPE_YISHENGYUZHAOZE       = "ysyzz"  // 医生与沼泽
	JUQING_TYPE_SHENHAIYIZU           = "shyz"   // 深海遗族
	JUQING_TYPE_SHIJUNZHIZUI          = "sjzz"   // 弑君之罪
	JUQING_TYPE_TONGXIANGWEIYIJIE     = "txwyj"  // 同相位异界
	JUQING_TYPE_YEXINZHONGHE          = "yxzh"   // 野心重荷
	JUQING_TYPE_FENGYINZHISHI         = "fyzs"   // 封印之石
	JUQING_TYPE_YINGXIONGDEZIGE       = "yxdzg"  // 英雄的资格
	JUQING_TYPE_TIANHE                = "th"     // 天河
	JUQING_TYPE_TIANKONGDAPINGYUAN    = "tkdpy"  // 天空大平原
	JUQING_TYPE_XUNNANGUANGCHANGRUKOU = "xngcrk" // 殉难广场入口
	JUQING_TYPE_BANGONGXIUXISHI       = "bgxxs"  // 办公休息室
	JUQING_TYPE_TIANBIANYUNXIA        = "tbyx"   // 天边云霞
	JUQING_TYPE_ZHONGJIBINGQI         = "zjbq"   // 终极兵器
)

func (c *chaoJiDou) JuQing(jt JuQingType) {
	// 打开任务面板
	c.press(robotgo.F5, 3)

	// 打开主页
	c.clickButton(c.RenWuMap.ZhuYe, 2)

	// 自动前往
	c.clickButton(c.RenWuMap.FirstXunLu, 30)

	if jt == JUQING_TYPE_YISHENGYUZHAOZE {
		c.yishengyuzhaozeHelper()
	}

}

func (c *chaoJiDou) yishengyuzhaozeHelper() {
	// 对话
	c.press(robotgo.KeyF, 5)

	// 选大地图
	c.clickButton(c.JuQingMap.TongHuaZhenButton, 3)

	// 选小地图
	c.clickButton(c.JuQingMap.FuBens[JUQING_TYPE_YISHENGYUZHAOZE].Window, 3)

	// 入场
	c.clickButton(c.EnterButton, 3)
	c.clickButton(c.EnterDButton, 60)

	// 开始战斗
	c.clickButton(c.StartBattleButton, 7)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)

	// 第1张怪物图
	c.clickButton(c.JuQingMap.FuBens[JUQING_TYPE_YISHENGYUZHAOZE].SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.KeyA, 1)
	c.move(1025, 219, 1, 3)
	robotgo.Sleep(9)
	c.press(robotgo.F1, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)

	// 第2张怪物图
	c.clickButton(c.JuQingMap.FuBens[JUQING_TYPE_YISHENGYUZHAOZE].SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.KeyA, 1)
	c.move(1061, 118, 1, 3)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	c.multiMove(993, 186, 1, 1, 2)
	robotgo.Sleep(3)
	c.press(robotgo.F1, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)

	// 第3张怪物图
	c.clickButton(c.JuQingMap.FuBens[JUQING_TYPE_YISHENGYUZHAOZE].SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.KeyA, 1)
	robotgo.Sleep(12)
	c.press(robotgo.F1, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)

	// 第4张怪物图
	c.clickButton(c.JuQingMap.FuBens[JUQING_TYPE_YISHENGYUZHAOZE].SmallMap[3], 6)
	c.press(robotgo.KeyA, 1)
	c.press(robotgo.KeyD, 1)
	c.move(1114, 238, 1, 4)
	robotgo.MoveSmooth(253, 657, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(10)
	c.press(robotgo.F1, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)

	// 第5张怪物图
	c.clickButton(c.JuQingMap.FuBens[JUQING_TYPE_YISHENGYUZHAOZE].SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.KeyA, 1)
	robotgo.Sleep(6)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)
	c.press(robotgo.F1, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)

	// 第6张怪物图
	c.clickButton(c.JuQingMap.FuBens[JUQING_TYPE_YISHENGYUZHAOZE].SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.KeyA, 1)
	c.move(1074, 343, 1, 3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)
	c.multiMove(815, 213, 1, 1, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)
	c.press(robotgo.F1, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)

	// 第7张怪物图: boss
	c.clickButton(c.JuQingMap.FuBens[JUQING_TYPE_YISHENGYUZHAOZE].SmallMap[6], 8)
	c.press(robotgo.Esc, 1)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.KeyA, 1)
	c.press(robotgo.KeyS, 1)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyE, 1)
	c.press(robotgo.KeyW, 1)
	c.press(robotgo.KeyQ, 1)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.F1, 1)
	c.multiMove(1160, 618, 1, 1, 6)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.KeyA, 1)
	c.multiMove(1160, 618, 1, 1, 6)
	c.press(robotgo.F1, 1)
	c.press(robotgo.KeyD, 1)
	c.multiMove(200, 618, 1, 1, 6)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.KeyA, 1)
	c.multiMove(1160, 618, 1, 1, 6)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)

	robotgo.Sleep(10)
	// 返回主城
	robotgo.KeyPress(robotgo.F12)
	robotgo.Sleep(5)
	robotgo.KeyPress(robotgo.KeyF)
	robotgo.Sleep(3)
	robotgo.KeyPress(robotgo.KeyF)
	robotgo.Sleep(20)
}

func SearchJuqing(juqingArray []JuQingType, jt JuQingType) int {
	for index := range juqingArray {
		if juqingArray[index] == jt {
			return index
		}
	}

	return -1
}
