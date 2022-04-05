package chaojidou

import "C"
import "github.com/go-vgo/robotgo"

type MeiRiType string

const (
	MEIRI_TYPE_HEWEIANYING     MeiRiType = "hway"
	MEIRI_TYPE_SHENHAIZHUKE    MeiRiType = "shzk"
	MEIRI_TYPE_ZUIQIANGJIANSHI MeiRiType = "zqjs"
	MEIRI_TYPE_TIEBISHOUWEI    MeiRiType = "tbsw"
)

func (c *chaoJiDou) MeiRiTiaoZhan(mt MeiRiType, dt DifficultyType) {
	robotgo.KeyPress(robotgo.KeyM)
	robotgo.Sleep(1)

	// 大地图点每日挑战
	c.clickButton(c.BigMap.MeiRi, 30)

	// 对话
	c.press(robotgo.KeyF, 3)
	robotgo.KeyPress(robotgo.KeyF)
	robotgo.Sleep(8)

	// 点中地图
	c.clickButton(c.MeiRiMap.Window, 3)

	// 入场
	c.clickButton(c.EnterButton, 3)
	c.clickButton(c.EnterDButton, 3) // 点两次确认入场
	c.clickButton(c.EnterDButton, 60)

	// 打怪
	if mt == MEIRI_TYPE_HEWEIANYING {
		c.heweianying()
	} else if mt == MEIRI_TYPE_SHENHAIZHUKE {
		c.shenhaizhuke()
	} else if mt == MEIRI_TYPE_ZUIQIANGJIANSHI {
		c.zuiqiangjianshi()
	} else if mt == MEIRI_TYPE_TIEBISHOUWEI {
		c.tiebishouwei()
	}

	// 返回主城
	robotgo.Sleep(180)
	robotgo.KeyPress(robotgo.F12)
	robotgo.Sleep(5)
	c.press(robotgo.KeyF, 3)
	robotgo.KeyPress(robotgo.KeyF)
	robotgo.Sleep(20)
}

func (c *chaoJiDou) heweianying() {
	// 开始战斗
	c.clickButton(c.StartBattleButton, 13)

}

func (c *chaoJiDou) shenhaizhuke() {
	// 开始战斗
	c.clickButton(c.StartBattleButton, 13)

	c.shenhaiyizuMeiri()
	c.shijunzhizuiMeiri()
	c.tongxiangweiyijieMeiri()
	c.yexinzhongheMeiri()
}

func (c *chaoJiDou) zuiqiangjianshi() {
	// 开始战斗
	c.clickButton(c.StartBattleButton, 13)

	c.fengyinzhishiMeiri()
	c.yingxiongdezigeMeiri()
	c.tianheMeiri()
	c.tiankongdapingyuanMeiri()
}

func (c *chaoJiDou) tiebishouwei() {
	// 开始战斗
	c.clickButton(c.StartBattleButton, 13)

	c.xunnanguangchangrukouMeiri()
	c.bangongxiuxishiMeiri()
	c.tianbianyunxiaMeiri()
	c.zhongjibingqiMeiri()
}

func (c *chaoJiDou) shenhaiyizuMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.ShenHaiZhuKe[0]]
	c.press(robotgo.KeyS, 8)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyE, 2)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyW, 1)
	c.press(robotgo.Key3, 2)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(105, 181, 1, 1, 5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(620, 40, 1, 1, 6)
	robotgo.Sleep(5)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	robotgo.Sleep(10)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1134, 540, 1, 1, 3)
	robotgo.Sleep(8)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1027, 245, 1, 3)
	robotgo.Sleep(8)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 5)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1079, 367, 1, 1, 6)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	c.multiMove(350, 231, 1, 1, 3)
	robotgo.Sleep(5)

	// 第7张怪物图
	c.clickButton(fuben.SmallMap[6], 5)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(362, 161, 1, 3)
	robotgo.Sleep(5)

	// 第8张怪物图： boss
	c.clickButton(fuben.SmallMap[7], 8)
	c.press(robotgo.Esc, 1)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	robotgo.Sleep(8)

	// 下一张图
	c.move(930, 268, 1, 5)
	c.press(robotgo.KeyF, 3)
	robotgo.Sleep(40)
}

func (c *chaoJiDou) shijunzhizuiMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.ShenHaiZhuKe[1]]
	c.press(robotgo.KeyS, 8)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyE, 2)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyW, 1)
	c.press(robotgo.Key3, 2)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(829, 134, 1, 1, 3)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	c.multiMove(67, 298, 1, 1, 3)
	robotgo.Sleep(5)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(913, 134, 1, 1, 3)
	robotgo.Sleep(8)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 5)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1023, 64, 1, 3)
	robotgo.Sleep(8)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(100, 100, 1, 1, 3)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	c.multiMove(622, 638, 1, 1, 4)
	robotgo.Sleep(5)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 5)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1118, 652, 1, 1, 5)
	robotgo.Sleep(8)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 5)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(223, 675, 1, 1, 2)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	c.multiMove(1170, 549, 1, 1, 6)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	c.multiMove(160, 539, 1, 1, 3)
	robotgo.Sleep(5)

	// 第7张怪物图： boss
	c.clickButton(fuben.SmallMap[6], 5)
	c.press(robotgo.Esc, 1)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	robotgo.Sleep(8)

	// 下一张图
	c.move(1134, 498, 1, 5)
	c.press(robotgo.KeyF, 3)
	robotgo.Sleep(40)
}

func (c *chaoJiDou) tongxiangweiyijieMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.ShenHaiZhuKe[2]]
	c.press(robotgo.KeyS, 8)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyE, 2)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyW, 1)
	c.press(robotgo.Key3, 2)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(285, 634, 1, 1, 5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1080, 622, 1, 1, 6)
	robotgo.Sleep(6)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 5)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	robotgo.Sleep(12)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 11)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	robotgo.Sleep(12)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(750, 144, 1, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)

	// 第6张怪物图： boss
	c.clickButton(fuben.SmallMap[5], 8)
	c.press(robotgo.Esc, 1)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	robotgo.Sleep(8)

	// 下一张图
	c.move(808, 202, 1, 5)
	c.press(robotgo.KeyF, 3)
	robotgo.Sleep(40)
}

func (c *chaoJiDou) yexinzhongheMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.ShenHaiZhuKe[3]]
	c.press(robotgo.KeyS, 8)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyE, 2)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyW, 1)
	c.press(robotgo.Key3, 2)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(193, 379, 1, 1, 5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 6)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(186, 425, 1, 1, 9)
	robotgo.Sleep(5)

	// 第3张怪物图
	c.press(robotgo.KeyF, 5)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(282, 559, 1, 1, 2)
	robotgo.Sleep(3)
	c.multiMove(936, 638, 1, 1, 2)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(351, 591, 1, 1, 4)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1257, 149, 1, 5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1168, 110, 1, 10)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)

	// 第7张怪物图： boss
	c.clickButton(fuben.SmallMap[6], 5)
	c.press(robotgo.Esc, 1)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	robotgo.Sleep(20)
}

func (c *chaoJiDou) fengyinzhishiMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.ZuiQiangJianShi[0]]
	c.press(robotgo.KeyS, 8)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyE, 2)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyW, 1)
	c.press(robotgo.Key3, 2)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	robotgo.MoveSmooth(181, 571, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	robotgo.MoveSmooth(70, 307, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	robotgo.MoveSmooth(1078, 628, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	robotgo.MoveSmooth(1135, 401, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 11)
	robotgo.MoveSmooth(1119, 271, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 11)
	robotgo.MoveSmooth(1149, 273, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)

	// 第7张怪物图： boss
	c.clickButton(fuben.SmallMap[6], 8)
	c.press(robotgo.Esc, 1)
	robotgo.MoveSmooth(1079, 221, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(11)

	// 下一张图
	c.move(1036, 309, 1, 5)
	c.press(robotgo.KeyF, 3)
	robotgo.Sleep(40)
}

func (c *chaoJiDou) yingxiongdezigeMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.ZuiQiangJianShi[1]]
	c.press(robotgo.KeyS, 8)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyE, 2)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyW, 1)
	c.press(robotgo.Key3, 2)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	robotgo.MoveSmooth(533, 139, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 10)
	robotgo.MoveSmooth(188, 459, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	robotgo.MoveSmooth(18, 170, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(18, 170, 1, 5)
	c.press(robotgo.KeyD, 1)
	c.move(408, 46, 1, 6)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 5)
	robotgo.MoveSmooth(118, 216, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	robotgo.MoveSmooth(76, 618, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(254, 728, 1, 1, 5)
	robotgo.Sleep(5)

	// 第6张怪物图： boss
	c.clickButton(fuben.SmallMap[5], 5)
	c.press(robotgo.Esc, 1)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)

	// 下一张图
	c.move(1129, 441, 1, 5)
	c.press(robotgo.KeyF, 3)
	robotgo.Sleep(40)
}

func (c *chaoJiDou) tianheMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.ZuiQiangJianShi[2]]
	c.press(robotgo.KeyS, 8)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyE, 2)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyW, 1)
	c.press(robotgo.Key3, 2)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(790, 64, 1, 1, 4)
	robotgo.Sleep(2)
	c.multiMove(1260, 100, 1, 1, 2)
	robotgo.Sleep(2)
	c.press(robotgo.KeyF, 4)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	robotgo.Sleep(5)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(538, 121, 1, 1, 5)
	robotgo.Sleep(6)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 6)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(978, 131, 1, 1, 4)
	robotgo.Sleep(6)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 6)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)

	// 第6张怪物图： boss
	c.clickButton(fuben.SmallMap[5], 8)
	c.press(robotgo.Esc, 1)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)

	// 下一张图
	c.move(875, 180, 1, 5)
	c.press(robotgo.KeyF, 3)
	robotgo.Sleep(40)
}

func (c *chaoJiDou) tiankongdapingyuanMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.ZuiQiangJianShi[3]]
	c.press(robotgo.KeyS, 8)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyE, 2)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyW, 1)
	c.press(robotgo.Key3, 2)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(992, 153, 1, 1, 5)
	robotgo.Sleep(3)
	c.press(robotgo.KeyF, 4)
	c.press(robotgo.KeyD, 1)
	c.multiMove(543, 606, 1, 1, 3)
	robotgo.Sleep(5)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(935, 125, 1, 1, 4)
	robotgo.Sleep(6)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1139, 133, 1, 1, 6)
	robotgo.Sleep(4)
	c.press(robotgo.KeyF, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(9)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(755, 87, 1, 1, 4)
	robotgo.Sleep(4)
	c.press(robotgo.KeyF, 4)
	c.press(robotgo.KeyD, 1)
	c.multiMove(670, 60, 1, 1, 4)
	robotgo.Sleep(4)
	c.press(robotgo.KeyF, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 6)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(675, 54, 1, 5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第7张怪物图： boss
	c.clickButton(fuben.SmallMap[6], 13)
	c.press(robotgo.Esc, 1)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	robotgo.Sleep(25)
}

func (c *chaoJiDou) xunnanguangchangrukouMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.TieBiShouWei[0]]
	c.press(robotgo.KeyS, 8)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyE, 2)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyW, 1)
	c.press(robotgo.Key3, 2)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(11)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 14)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 10)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(670, 39, 1, 1, 4)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)

	// 第6张怪物图： boss
	c.clickButton(fuben.SmallMap[5], 7)
	c.press(robotgo.Esc, 1)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(11)

	// 下一张图
	c.move(1084, 377, 1, 5)
	c.press(robotgo.KeyF, 3)
	robotgo.Sleep(40)
}

func (c *chaoJiDou) bangongxiuxishiMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.TieBiShouWei[1]]
	c.press(robotgo.KeyS, 8)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyE, 2)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyW, 1)
	c.press(robotgo.Key3, 2)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(679, 125, 1, 1, 3)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	c.multiMove(679, 125, 1, 1, 2)
	robotgo.Sleep(6)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 10)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(973, 176, 1, 1, 2)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(574, 107, 1, 5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 6)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(373, 108, 2, 5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	c.move(244, 270, 2, 5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(672, 83, 2, 5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)

	// 第7张怪物图
	c.clickButton(fuben.SmallMap[6], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1022, 80, 2, 5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)

	// 第8张怪物图： boss
	c.clickButton(fuben.SmallMap[7], 8)
	c.press(robotgo.Esc, 1)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(11)

	// 下一张图
	c.move(946, 229, 1, 5)
	c.press(robotgo.KeyF, 3)
	robotgo.Sleep(40)
}

func (c *chaoJiDou) tianbianyunxiaMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.TieBiShouWei[2]]
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyE, 2)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyW, 1)
	c.press(robotgo.Key3, 2)
	robotgo.Sleep(3)

	// 第1张怪物图：无法自动寻径
	c.move(614, 138, 1, 5)
	c.press(robotgo.KeyF, 5)

	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1083, 112, 2, 5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1098, 101, 1, 1, 4)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 10)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1094, 74, 2, 5)
	c.press(robotgo.KeyD, 1)
	// 这张图有bug，轰炸机会消失
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyT, 4)
	c.press(robotgo.KeyW, 1)
	robotgo.Sleep(3)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1127, 673, 2, 5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1091, 611, 2, 5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(488, 109, 1, 1, 4)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	c.move(260, 331, 2, 3)
	c.press(robotgo.KeyW, 1)
	robotgo.Sleep(3)

	// 第7张怪物图： boss
	c.clickButton(fuben.SmallMap[6], 8)
	c.press(robotgo.Esc, 1)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(10)

	// 下一张图
	c.move(953, 259, 1, 5)
	c.press(robotgo.KeyF, 3)
	robotgo.Sleep(40)
}

func (c *chaoJiDou) zhongjibingqiMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.TieBiShouWei[3]]
	c.press(robotgo.KeyS, 8)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyE, 2)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyW, 1)
	c.press(robotgo.Key3, 2)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)

	// 第1张怪物图：无法自动寻径
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1036, 123, 2, 5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(322, 112, 2, 5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 10)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1234, 188, 2, 5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 10)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(79, 77, 2, 5)
	c.press(robotgo.KeyD, 1)
	c.move(139, 320, 2, 3)
	robotgo.Sleep(5)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(178, 85, 2, 5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(327, 153, 2, 5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)

	// 第7张怪物图： boss
	c.clickButton(fuben.SmallMap[6], 8)
	c.press(robotgo.Esc, 1)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(25)
}
