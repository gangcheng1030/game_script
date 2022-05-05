package chaojidou

import "C"
import "github.com/go-vgo/robotgo"

type MeiRiType string

const (
	MEIRI_TYPE_SHENHAIZHUKE      MeiRiType = "shzk"
	MEIRI_TYPE_ZUIQIANGJIANSHI   MeiRiType = "zqjs"
	MEIRI_TYPE_TIEBISHOUWEI      MeiRiType = "tbsw"
	MEIRI_TYPE_YANHUA            MeiRiType = "yh"
	MEIRI_TYPE_BINGFENGWANGGUO   MeiRiType = "bfwg"
	MEIRI_TYPE_XIAOXINCHUDIAN    MeiRiType = "xxcd"
	MEIRI_TYPE_RENXIAOGUIDA      MeiRiType = "rxgd"
	MEIRI_TYPE_JIUYUNZHIDIAN     MeiRiType = "jyzd"
	MEIRI_TYPE_SARANZHIHUA       MeiRiType = "srzh"
	MEIRI_TYPE_SHENGWUSHIYAN     MeiRiType = "swsy"
	MEIRI_TYPE_QIANGZHANDASHA    MeiRiType = "qzds"
	MEIRI_TYPE_HEWEIANYING       MeiRiType = "hway"
	MEIRI_TYPE_WEIXIANHUIZHUAN   MeiRiType = "wxhz"
	MEIRI_TYPE_WANGMINGTUDEDAOLU MeiRiType = "wmtddl"
	MEIRI_TYPE_SHENMIDONGWUXUE   MeiRiType = "smdwx"
	MEIRI_TYPE_HUOZHIJIDIAN      MeiRiType = "hzjd"
)

func (c *chaoJiDou) MeiRiTiaoZhan(mt MeiRiType, dt DifficultyType) {
	robotgo.KeyPress(robotgo.KeyM)
	robotgo.Sleep(1)

	// 大地图点每日挑战
	c.clickButton(c.BigMap.MeiRi, NpcWaitSecs)

	// 对话
	c.press(robotgo.KeyF, 3)
	robotgo.KeyPress(robotgo.KeyF)
	robotgo.Sleep(8)

	// 点中地图
	c.clickButton(c.MeiRiMap.Window, 3)

	// 入场
	c.clickButton(c.EnterButton, 3)
	c.clickButton(c.EnterDButton, 3) // 点两次确认入场
	c.clickButton(c.EnterDButton, ReadMapWaitSecs)

	// 打怪
	if mt == MEIRI_TYPE_HEWEIANYING {
		c.heweianying()
	} else if mt == MEIRI_TYPE_SHENHAIZHUKE {
		c.shenhaizhuke()
	} else if mt == MEIRI_TYPE_ZUIQIANGJIANSHI {
		c.zuiqiangjianshi()
	} else if mt == MEIRI_TYPE_TIEBISHOUWEI {
		c.tiebishouwei()
	} else if mt == MEIRI_TYPE_YANHUA {
		c.yanhua()
	} else if mt == MEIRI_TYPE_BINGFENGWANGGUO {
		c.bingfengwangguo()
	} else if mt == MEIRI_TYPE_XIAOXINCHUDIAN {
		c.xiaoxinchudian()
	} else if mt == MEIRI_TYPE_RENXIAOGUIDA {
		c.renxiaoguida()
	} else if mt == MEIRI_TYPE_JIUYUNZHIDIAN {
		c.jiuyunzhidian()
	} else if mt == MEIRI_TYPE_SARANZHIHUA {
		c.saranzhihua()
	} else if mt == MEIRI_TYPE_SHENGWUSHIYAN {
		c.shengwushiyan()
	} else if mt == MEIRI_TYPE_QIANGZHANDASHA {
		c.qiangzhandasha()
	} else if mt == MEIRI_TYPE_WEIXIANHUIZHUAN {
		c.weixianhuizhuan()
	} else if mt == MEIRI_TYPE_WANGMINGTUDEDAOLU {
		c.wangmingtudedaolu()
	} else if mt == MEIRI_TYPE_SHENMIDONGWUXUE {
		c.shenmidongwuxue()
	} else if mt == MEIRI_TYPE_HUOZHIJIDIAN {
		c.huozhijidian()
	}

	// 返回主城
	robotgo.Sleep(180)
	robotgo.KeyPress(robotgo.F12)
	robotgo.Sleep(5)
	c.press(robotgo.KeyF, 3)
	robotgo.KeyPress(robotgo.KeyF)
	robotgo.Sleep(20)
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

func (c *chaoJiDou) yanhua() {
	// 开始战斗
	c.clickButton(c.StartBattleButton, 13)

	c.quanjianzuozhanMeiri()
	c.qtshamokaifabuMeiri()
	c.tiankongzhizuanMeiri()
	c.xintianebaoMeiri()
}

func (c *chaoJiDou) bingfengwangguo() {
	// 开始战斗
	c.clickButton(c.StartBattleButton, 13)

	c.bingheMeiri()
	c.bingxuebuluomubacunMeiri()
	c.buguixiaguMeiri()
	c.bingfengchengbaoMeiri()
}

func (c *chaoJiDou) xiaoxinchudian() {
	// 开始战斗
	c.clickButton(c.StartBattleButton, 13)

	c.jianruixiaguMeiri()
	c.yueguangdadaoMeiri()
	c.shixianzhengyiMeiri()
}

func (c *chaoJiDou) renxiaoguida() {
	// 开始战斗
	c.clickButton(c.StartBattleButton, 13)

	c.shijiantingyuanMeiri()
	c.huweiduixiongmeiMeiri()
	c.tiancaishaonvMeiri()
}

func (c *chaoJiDou) jiuyunzhidian() {
	// 开始战斗
	c.clickButton(c.StartBattleButton, 13)

	c.qingzhulinMeiri()
	c.emeibieyuanMeiri()
	c.gaiyanMeiri()
	c.beimangshanMeiri()
	c.biwudasaijuesaiMeiri()
}

func (c *chaoJiDou) saranzhihua() {
	// 开始战斗
	c.clickButton(c.StartBattleButton, 13)

	c.kongzhonganxiMeiri()
	c.dixiatingchechangMeiri()
	c.nizongshensheshouMeiri()
	c.youlingMeiri()
}

func (c *chaoJiDou) shengwushiyan() {
	// 开始战斗
	c.clickButton(c.StartBattleButton, 13)

	c.dixiajidishiyanshiMeiri()
	c.yishengyuzhaozeMeiri()
	c.shaleduonaihuoshankouMeiri()
	c.haidelushiyanshiMeiri()
}

func (c *chaoJiDou) qiangzhandasha() {
	// 开始战斗
	c.clickButton(c.StartBattleButton, 13)

	c.kaisaliyagangMeiri()
	c.zhuqiaoMeiri()
	c.youlingcangkuMeiri()
	c.dixiajidizuizhongjuezhanMeiri()
}

func (c *chaoJiDou) heweianying() {
	// 开始战斗
	c.clickButton(c.StartBattleButton, 13)

	c.zaicimianduiweixianMeiri()
	c.heimeiguitingyuanMeiri()
	c.bolizhidiMeiri()
	c.ranheidewanxia1Meiri()
}

func (c *chaoJiDou) weixianhuizhuan() {
	// 开始战斗
	c.clickButton(c.StartBattleButton, 13)

	c.senlinbianjieMeiri()
	c.migongyuqiutuMeiri()
	c.denglulalaiyeMeiri()
	c.luosudedixiashijieMeiri()
}

func (c *chaoJiDou) wangmingtudedaolu() {
	// 开始战斗
	c.clickButton(c.StartBattleButton, 13)

	c.haishangqiaoliangMeiri()
	c.heihaiMeiri()
	c.langzimoluMeiri()
}

func (c *chaoJiDou) shenmidongwuxue() {
	// 开始战斗
	c.clickButton(c.StartBattleButton, 13)

	c.hongseshamoMeiri()
	c.xunzhaochuluMeiri()
	c.yanhanxueshanMeiri()
	c.bingxuefugaizhiqiangMeiri()
}

func (c *chaoJiDou) huozhijidian() {
	// 开始战斗
	c.clickButton(c.StartBattleButton, 13)

	c.dixiajidijinjizuozhanMeiri()
	c.huoyanzhuisuizheMeiri()
	c.feiqirongluMeiri()
	c.huoyanzhinvMeiri()
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
	robotgo.Sleep(6)

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
	c.move(100, 100, 1, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
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
	c.move(908, 202, 1, 5)
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
	c.move(933, 613, 1, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	robotgo.Sleep(30)
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
	robotgo.MoveSmooth(575, 54, mouseSpeedX, mouseSpeedY)
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
	c.move(466, 311, 1, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	c.move(1135, 648, 1, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
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

func (c *chaoJiDou) quanjianzuozhanMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.YanHua[0]]
	c.press(robotgo.KeyS, 8)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyE, 2)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyW, 1)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 15)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1022, 64, 1, 1, 4)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(14)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(492, 46, 1, 1, 5)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(11)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(372, 117, 1, 1, 5)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(11)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(948, 116, 1, 1, 4)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(11)

	// 第5张怪物图: 最后一张图，无boss
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(681, 69, 1, 1, 3)
	robotgo.Sleep(3)
	robotgo.MoveSmooth(690, 80, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)
	c.multiMove(681, 69, 1, 1, 3)
	robotgo.Sleep(3)
	c.press(robotgo.KeyF, 19)

	// 下一张图
	c.move(1143, 560, 1, 5)
	c.press(robotgo.KeyF, 3)
	robotgo.Sleep(40)
}

func (c *chaoJiDou) qtshamokaifabuMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.YanHua[1]]
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
	c.multiMove(1226, 140, 1, 1, 3)
	robotgo.Sleep(3)
	robotgo.MoveSmooth(1216, 140, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	c.multiMove(1226, 140, 1, 1, 2)
	robotgo.Sleep(13)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1057, 106, 3, 5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1113, 415, 1, 1, 7)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1252, 142, 3, 6)
	c.press(robotgo.KeyD, 1)
	c.multiMove(1186, 681, 1, 1, 2)
	robotgo.Sleep(5)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1065, 87, 3, 5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)

	// 第6张怪物图： boss
	c.clickButton(fuben.SmallMap[5], 7)
	c.press(robotgo.Esc, 1)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(11)

	// 下一张图
	c.move(830, 186, 1, 5)
	c.press(robotgo.KeyF, 3)
	robotgo.Sleep(40)
}

func (c *chaoJiDou) tiankongzhizuanMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.YanHua[2]]
	c.press(robotgo.KeyS, 8)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyE, 2)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyW, 1)
	c.press(robotgo.Key3, 2)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)

	// 第1张怪物图：只有1张图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.KeyD, 1)
	c.multiMove(74, 655, 10, 1, 3)
	robotgo.MoveSmooth(1153, 164, 0.9, 0.9)
	c.continuedBattle(130)

	// 下一张图
	c.move(1091, 372, 1, 5)
	c.press(robotgo.KeyF, 3)
	robotgo.Sleep(40)
}

func (c *chaoJiDou) xintianebaoMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.YanHua[3]]
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
	c.move(1084, 635, 2, 5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1221, 658, 2, 5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1076, 102, 2, 5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 10)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(23, 325, 2, 5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(24, 723, 2, 5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1041, 127, 2, 5)
	c.press(robotgo.KeyD, 1)
	c.move(248, 92, 2, 5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)

	// 第7张怪物图： boss
	c.clickButton(fuben.SmallMap[6], 8)
	c.press(robotgo.Esc, 1)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(25)
}

func (c *chaoJiDou) bingheMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.BingFengWangGuo[0]]
	c.press(robotgo.KeyS, 8)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyE, 3)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyW, 1)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1075, 156, 3, 4)
	robotgo.MoveSmooth(1065, 156, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1277, 268, 3, 4)
	robotgo.MoveSmooth(1267, 268, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1258, 437, 1, 1, 2)
	robotgo.Sleep(3)
	robotgo.MoveSmooth(1248, 437, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(735, 49, 1, 1, 2)
	robotgo.Sleep(3)
	robotgo.MoveSmooth(725, 49, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)

	// 第6张怪物图： boss
	c.clickButton(fuben.SmallMap[4], 7)
	c.press(robotgo.Esc, 1)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(11)

	// 下一张图
	c.move(882, 219, 1, 5)
	c.press(robotgo.KeyF, 3)
	c.move(788, 472, 1, 3)
	c.press(robotgo.KeyF, 3)
	c.move(456, 217, 1, 4)
	c.press(robotgo.KeyF, 3)
	robotgo.Sleep(40)
}

func (c *chaoJiDou) bingxuebuluomubacunMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.BingFengWangGuo[1]]
	c.press(robotgo.KeyS, 8)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyE, 3)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyW, 1)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(733, 56, 3, 4)
	robotgo.MoveSmooth(722, 46, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(688, 37, 3, 5)
	robotgo.MoveSmooth(678, 37, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(702, 76, 3, 4)
	robotgo.MoveSmooth(692, 76, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(340, 181, 3, 4)
	robotgo.MoveSmooth(330, 181, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)

	// 第6张怪物图： boss
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.Esc, 1)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(11)

	// 下一张图
	c.move(770, 180, 1, 5)
	c.press(robotgo.KeyF, 3)
	c.move(888, 499, 1, 4)
	c.press(robotgo.KeyF, 3)
	robotgo.Sleep(40)
}

func (c *chaoJiDou) buguixiaguMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.BingFengWangGuo[2]]
	c.press(robotgo.KeyS, 8)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyE, 3)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyW, 1)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1145, 184, 3, 4)
	robotgo.MoveSmooth(1135, 184, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1319, 279, 1, 1, 2)
	robotgo.Sleep(3)
	robotgo.MoveSmooth(1309, 279, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(867, 164, 3, 4)
	robotgo.MoveSmooth(857, 164, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1002, 154, 3, 4)
	robotgo.MoveSmooth(992, 154, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)

	// 第7张怪物图： boss
	c.clickButton(fuben.SmallMap[6], 8)
	c.press(robotgo.Esc, 1)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(11)

	// 下一张图
	c.move(745, 191, 1, 5)
	c.press(robotgo.KeyF, 3)
	c.move(890, 506, 1, 4)
	c.press(robotgo.KeyF, 3)
	c.move(890, 506, 1, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	robotgo.Sleep(40)
}

func (c *chaoJiDou) bingfengchengbaoMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.BingFengWangGuo[3]]
	c.press(robotgo.KeyS, 8)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyE, 3)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyW, 1)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(258, 46, 3, 4)
	robotgo.MoveSmooth(248, 46, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1146, 126, 3, 4)
	robotgo.MoveSmooth(1136, 126, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(204, 80, 3, 4)
	robotgo.MoveSmooth(194, 80, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(157, 110, 3, 4)
	robotgo.MoveSmooth(147, 110, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(68, 278, 3, 4)
	robotgo.MoveSmooth(58, 278, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(11)

	// 第7张怪物图： boss
	c.clickButton(fuben.SmallMap[6], 8)
	c.press(robotgo.Esc, 1)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(25)
}

func (c *chaoJiDou) jianruixiaguMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.XiaoXinChuDian[0]]
	c.press(robotgo.KeyS, 8)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyE, 3)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyW, 1)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(781, 150, 3, 4)
	robotgo.MoveSmooth(681, 150, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)
	c.move(28, 70, 3, 4)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.F2, 1)
	c.move(63, 639, 3, 4)
	robotgo.MoveSmooth(53, 639, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(231, 83, 3, 4)
	robotgo.MoveSmooth(131, 83, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)
	c.move(254, 83, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(763, 56, 3, 4)
	robotgo.MoveSmooth(663, 56, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1138, 49, 3, 4)
	robotgo.MoveSmooth(1038, 49, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1274, 340, 3, 4)
	robotgo.MoveSmooth(1174, 340, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)

	// 第7张怪物图
	c.clickButton(fuben.SmallMap[6], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.F1, 1)
	robotgo.Sleep(5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.move(1185, 213, 3, 4)
	robotgo.MoveSmooth(1085, 213, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)

	// 第8张怪物图： boss
	c.clickButton(fuben.SmallMap[7], 8)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)

	// 下一张图
	c.move(891, 207, 1, 5)
	c.press(robotgo.KeyF, 3)
	c.move(926, 586, 1, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	robotgo.Sleep(40)
}

func (c *chaoJiDou) yueguangdadaoMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.XiaoXinChuDian[1]]
	c.press(robotgo.KeyS, 8)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyE, 3)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyW, 1)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	robotgo.Sleep(4)
	c.move(1198, 164, 3, 4)
	robotgo.MoveSmooth(1098, 164, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.F2, 1)
	robotgo.Sleep(4)
	c.move(1093, 114, 3, 4)
	robotgo.MoveSmooth(993, 114, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	robotgo.Sleep(4)
	c.multiMove(577, 37, 1, 1, 2)
	robotgo.Sleep(4)
	robotgo.MoveSmooth(477, 37, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	robotgo.Sleep(4)
	c.move(1174, 93, 3, 4)
	robotgo.MoveSmooth(1074, 93, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)
	c.multiMove(1327, 452, 1, 1, 2)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 10)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(11)

	// 第6张怪物图： boss
	c.clickButton(fuben.SmallMap[5], 10)
	c.press(robotgo.Esc, 1)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(11)

	// 下一张图
	c.move(832, 188, 1, 5)
	c.press(robotgo.KeyF, 3)
	c.move(949, 650, 1, 5)
	c.press(robotgo.KeyF, 3)
	robotgo.Sleep(40)
}

func (c *chaoJiDou) shixianzhengyiMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.XiaoXinChuDian[2]]
	c.press(robotgo.KeyS, 8)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyE, 3)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyW, 1)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(10)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.F2, 1)
	c.move(223, 130, 1, 4)
	robotgo.MoveSmooth(123, 130, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)
	c.move(1150, 153, 1, 6)
	robotgo.MoveSmooth(1050, 153, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)
	c.move(1010, 261, 1, 6)
	robotgo.MoveSmooth(910, 261, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)
	c.multiMove(204, 77, 1, 1, 5)
	robotgo.MoveSmooth(104, 77, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)
	c.multiMove(204, 77, 1, 1, 5)
	robotgo.MoveSmooth(104, 77, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)
	c.multiMove(204, 77, 1, 1, 5)
	robotgo.MoveSmooth(104, 77, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(35, 311, 3, 4)
	robotgo.MoveSmooth(25, 311, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 10)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(218, 629, 3, 4)
	robotgo.MoveSmooth(118, 629, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(110, 637, 1, 1, 3)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)
	c.multiMove(338, 95, 1, 1, 2)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)

	// 第7张怪物图： boss
	c.clickButton(fuben.SmallMap[6], 8)
	c.press(robotgo.Esc, 1)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(25)
}

func (c *chaoJiDou) shijiantingyuanMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.RenXiaoGuiDa[0]]
	c.press(robotgo.KeyS, 10)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyE, 3)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyW, 1)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(956, 132, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.F2, 1)
	c.multiMove(664, 48, 1, 1, 2)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(99, 393, 1, 1, 3)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 12)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.multiMove(455, 87, 1, 1, 4)
	robotgo.Sleep(2)
	c.press(robotgo.KeyT, 4)
	c.press(robotgo.KeyE, 2)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.Key3, 2)

	// 第7张怪物图： boss
	c.clickButton(fuben.SmallMap[6], 8)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)

	// 下一张图
	c.move(935, 236, 1, 5)
	c.press(robotgo.KeyF, 3)
	c.move(354, 337, 1, 5)
	c.press(robotgo.KeyF, 3)
	c.move(1227, 625, 1, 4)
	c.press(robotgo.KeyF, 3)
	robotgo.Sleep(40)
}

func (c *chaoJiDou) huweiduixiongmeiMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.RenXiaoGuiDa[1]]
	c.press(robotgo.KeyS, 10)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyE, 3)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyW, 1)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1021, 72, 3, 4)
	robotgo.MoveSmooth(1221, 72, 0.9, 0.9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 10)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.F2, 1)
	c.move(225, 69, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1032, 64, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1097, 680, 1, 1, 2)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1142, 178, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1113, 90, 1, 1, 2)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第7张怪物图： boss
	c.clickButton(fuben.SmallMap[6], 8)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)

	// 下一张图
	c.move(859, 182, 1, 5)
	c.press(robotgo.KeyF, 3)
	c.move(851, 532, 1, 4)
	c.press(robotgo.KeyF, 3)
	robotgo.Sleep(40)
}

func (c *chaoJiDou) tiancaishaonvMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.RenXiaoGuiDa[2]]
	c.press(robotgo.KeyS, 10)
	c.press(robotgo.KeyQ, 4)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyE, 3)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyW, 1)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(3)

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(566, 148, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.F2, 1)
	c.move(1144, 266, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.F2, 1)
	c.move(617, 62, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1277, 445, 1, 1, 3)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 10)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1277, 370, 1, 1, 2)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.F2, 1)
	c.move(1211, 312, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第7张怪物图： boss
	c.clickButton(fuben.SmallMap[6], 8)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(25)
}

func (c *chaoJiDou) qingzhulinMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.JiuYunZhiDian[0]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(201, 75, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.F2, 1)
	c.move(530, 48, 3, 6)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(289, 126, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 5)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(230, 675, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(44, 149, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)

	// 第7张怪物图
	c.clickButton(fuben.SmallMap[6], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(791, 116, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)
	c.press(robotgo.KeyR, 1)
	c.press(robotgo.KeyE, 3)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyT, 4)
	robotgo.Sleep(1)

	// 第8张怪物图： boss
	c.clickButton(fuben.SmallMap[7], 6)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)

	// 下一张图
	c.move(858, 208, 1, 5)
	c.press(robotgo.KeyF, 3)
	c.move(872, 511, 1, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	robotgo.Sleep(40)
}

func (c *chaoJiDou) emeibieyuanMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.JiuYunZhiDian[1]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1120, 49, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.F2, 1)
	c.move(1145, 65, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(179, 101, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1109, 78, 1, 1, 3)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第6张怪物图： boss
	c.clickButton(fuben.SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)

	// 下一张图
	c.move(840, 190, 1, 5)
	c.press(robotgo.KeyF, 3)
	c.move(1073, 692, 1, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	robotgo.Sleep(40)
}

func (c *chaoJiDou) gaiyanMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.JiuYunZhiDian[2]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1102, 91, 3, 5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.F2, 1)
	c.move(278, 93, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(633, 65, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 5)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(593, 30, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(3)
	c.move(7, 277, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1135, 91, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第7张怪物图： boss
	c.clickButton(fuben.SmallMap[6], 8)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(8)

	// 下一张图
	c.move(973, 288, 1, 5)
	c.press(robotgo.KeyF, 3)
	c.move(513, 242, 1, 4)
	c.press(robotgo.KeyF, 3)
	robotgo.Sleep(40)
}

func (c *chaoJiDou) beimangshanMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.JiuYunZhiDian[3]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 10)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.F2, 1)
	c.move(1070, 125, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 13)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1091, 120, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)
	c.move(1094, 28, 1, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第6张怪物图： boss
	c.clickButton(fuben.SmallMap[5], 7)
	c.multiMove(50, 715, 1, 1, 9)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 下一张图
	c.move(867, 203, 1, 4)
	c.press(robotgo.KeyF, 3)
	c.move(872, 550, 1, 4)
	c.press(robotgo.KeyF, 3)
	robotgo.Sleep(40)
}

func (c *chaoJiDou) biwudasaijuesaiMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.JiuYunZhiDian[4]]
	c.prestart()

	// 第1张怪物图：boss图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.continuedBattle(70)
}

func (c *chaoJiDou) kongzhonganxiMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.SaRanZhiHua[0]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1148, 52, 3, 3)
	robotgo.MoveSmooth(1148, 152, mouseSpeedX, mouseSpeedY)
	c.press(robotgo.KeyD, 1)
	c.move(1148, 52, 3, 3)
	robotgo.MoveSmooth(1148, 152, mouseSpeedX, mouseSpeedY)
	c.press(robotgo.KeyD, 1)
	c.move(1148, 52, 3, 3)
	robotgo.MoveSmooth(1148, 152, mouseSpeedX, mouseSpeedY)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.F2, 1)
	c.multiMove(673, 71, 1, 1, 3)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	c.multiMove(1083, 356, 1, 1, 3)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1160, 74, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(249, 156, 1, 1, 3)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	c.move(652, 57, 3, 4)
	c.press(robotgo.KeyD, 1)
	c.multiMove(1037, 93, 1, 1, 3)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第6张怪物图： boss
	c.press(robotgo.KeyT, 4)
	c.clickButton(fuben.SmallMap[5], 8)
	robotgo.Sleep(25)

	// 下一张图
	c.move(897, 225, 1, 4)
	c.press(robotgo.KeyF, 3)
	//c.move(872, 550, 1, 4)
	//c.press(robotgo.KeyF, 3)
	robotgo.Sleep(40)
}

func (c *chaoJiDou) dixiatingchechangMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.SaRanZhiHua[1]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(430, 51, 3, 4)
	robotgo.MoveSmooth(430, 151, mouseSpeedX, mouseSpeedY)
	c.press(robotgo.KeyD, 1)
	c.move(430, 51, 3, 4)
	robotgo.MoveSmooth(430, 151, mouseSpeedX, mouseSpeedY)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.F2, 1)
	c.multiMove(72, 346, 1, 1, 3)
	robotgo.Sleep(3)
	robotgo.MoveSmooth(72, 446, mouseSpeedX, mouseSpeedY)
	c.press(robotgo.KeyD, 1)
	c.multiMove(72, 346, 1, 1, 4)
	robotgo.Sleep(3)
	robotgo.MoveSmooth(72, 446, mouseSpeedX, mouseSpeedY)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(261, 357, 1, 1, 2)
	robotgo.Sleep(3)
	robotgo.MoveSmooth(261, 457, mouseSpeedX, mouseSpeedY)
	c.press(robotgo.KeyD, 1)
	c.move(261, 357, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(162, 67, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(224, 34, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第6张怪物图： boss
	c.press(robotgo.KeyT, 4)
	c.clickButton(fuben.SmallMap[5], 8)
	robotgo.Sleep(25)

	// 下一张图
	c.move(915, 229, 1, 4)
	c.press(robotgo.KeyF, 3)
	//c.move(872, 550, 1, 4)
	//c.press(robotgo.KeyF, 3)
	robotgo.Sleep(40)
}

func (c *chaoJiDou) nizongshensheshouMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.SaRanZhiHua[2]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(278, 138, 1, 1, 6)
	robotgo.Sleep(2)
	robotgo.MoveSmooth(278, 238, mouseSpeedX, mouseSpeedY)
	c.press(robotgo.KeyD, 1)
	c.multiMove(278, 138, 1, 1, 6)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	c.multiMove(1099, 457, 1, 1, 3)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	// 轰炸机可能会消失
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyT, 4)
	c.press(robotgo.KeyW, 1)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.move(414, 123, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(236, 363, 1, 1, 4)
	c.press(robotgo.KeyD, 1)
	c.multiMove(404, 659, 1, 1, 5)
	c.press(robotgo.KeyD, 1)
	c.multiMove(156, 63, 1, 1, 4)
	c.press(robotgo.KeyD, 1)
	c.move(1146, 309, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(277, 658, 1, 1, 4)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	c.move(300, 158, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.move(272, 148, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(248, 66, 1, 1, 3)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	c.multiMove(1138, 408, 1, 1, 3)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	c.multiMove(173, 165, 1, 1, 12)
	c.press(robotgo.KeyD, 1)
	c.multiMove(1099, 358, 1, 1, 3)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	c.multiMove(1131, 456, 1, 1, 3)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	c.multiMove(926, 169, 1, 1, 4)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第7张怪物图： boss
	c.press(robotgo.KeyT, 4)
	c.clickButton(fuben.SmallMap[6], 8)
	robotgo.Sleep(25)

	// 下一张图
	c.move(879, 215, 1, 4)
	c.press(robotgo.KeyF, 3)
	//c.move(872, 550, 1, 4)
	//c.press(robotgo.KeyF, 3)
	robotgo.Sleep(40)
}

func (c *chaoJiDou) youlingMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.SaRanZhiHua[3]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(992, 217, 1, 1, 3)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	c.move(1142, 261, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.F2, 1)
	c.move(1130, 379, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1098, 101, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(546, 35, 3, 4)
	c.press(robotgo.KeyD, 1)
	c.move(1267, 156, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(398, 219, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第7张怪物图： boss
	c.press(robotgo.KeyT, 4)
	c.clickButton(fuben.SmallMap[6], 8)
	robotgo.Sleep(25)
}

func (c *chaoJiDou) dixiajidishiyanshiMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.ShengWuShiYan[0]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.F2, 1)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(5)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1104, 230, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(811, 69, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 5)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1180, 60, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第7张怪物图： boss
	c.clickButton(fuben.SmallMap[6], 13)
	c.multiMove(641, 667, 1, 1, 5)
	robotgo.MoveSmooth(702, 34, mouseSpeedX, mouseSpeedY)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 1)
	c.press(robotgo.KeyT, 4)
	for i := 0; i < 4; i++ {
		c.press(robotgo.KeyD, 5)
	}

	// 下一张图
	c.move(683, 158, 1, 4)
	c.press(robotgo.KeyF, 3)
	c.move(442, 424, 1, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 3)
	c.move(1132, 362, 1, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	robotgo.Sleep(30)
}

func (c *chaoJiDou) yishengyuzhaozeMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.ShengWuShiYan[1]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1083, 74, 3, 5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 6)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.F2, 1)
	c.move(1063, 55, 3, 4)
	c.press(robotgo.KeyD, 1)
	c.move(1039, 89, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 6)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1142, 53, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 6)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1105, 627, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 10)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 11)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1270, 216, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第7张怪物图： boss
	c.clickButton(fuben.SmallMap[6], 11)
	c.press(robotgo.Esc, 1)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 下一张图
	c.move(929, 225, 1, 4)
	c.press(robotgo.KeyF, 3)
	c.move(722, 540, 1, 3)
	c.press(robotgo.KeyF, 3)
	robotgo.Sleep(30)
}

func (c *chaoJiDou) shaleduonaihuoshankouMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.ShengWuShiYan[2]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(217, 42, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.F2, 1)
	c.move(55, 312, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(564, 63, 3, 4)
	c.press(robotgo.KeyD, 1)
	c.move(176, 173, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(742, 38, 1, 1, 2)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 1)
	c.move(699, 53, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第6张怪物图： boss
	c.clickButton(fuben.SmallMap[5], 8)
	c.multiMove(630, 661, 1, 1, 6)
	robotgo.MoveSmooth(670, 94, mouseSpeedX, mouseSpeedY)
	for i := 0; i < 7; i++ {
		c.press(robotgo.KeyD, 5)
	}

	// 下一张图
	c.move(671, 161, 1, 4)
	c.press(robotgo.KeyF, 3)
	c.move(944, 363, 1, 3)
	c.press(robotgo.KeyF, 3)
	robotgo.Sleep(30)
}

func (c *chaoJiDou) haidelushiyanshiMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.ShengWuShiYan[3]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(610, 176, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.F2, 1)
	c.move(678, 74, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1110, 149, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 6)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1284, 355, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 10)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1070, 668, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.F2, 1)
	c.multiMove(140, 675, 1, 1, 2)
	robotgo.Sleep(3)
	c.press(robotgo.KeyT, 4)
	for i := 0; i < 2; i++ {
		c.press(robotgo.KeyD, 5)
	}

	// 第7张怪物图： boss
	c.clickButton(fuben.SmallMap[6], 8)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(20)
}

func (c *chaoJiDou) kaisaliyagangMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.QiangZhanDaSha[0]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1061, 116, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.F2, 1)
	c.move(1217, 655, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(197, 80, 3, 5)
	c.press(robotgo.KeyD, 1)
	c.move(23, 183, 3, 5)
	c.press(robotgo.KeyD, 1)
	c.move(113, 55, 3, 4)
	c.press(robotgo.KeyD, 1)
	c.move(1041, 258, 3, 3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1126, 687, 1, 1, 4)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	c.multiMove(1126, 687, 1, 1, 3)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 1)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(2)
	c.press(robotgo.KeyT, 4)

	// 第6张怪物图： boss
	c.clickButton(fuben.SmallMap[5], 6)
	c.press(robotgo.Esc, 1)
	robotgo.Sleep(15)

	// 下一张图
	c.move(914, 223, 1, 4)
	c.press(robotgo.KeyF, 3)
	c.move(348, 330, 1, 4)
	c.press(robotgo.KeyF, 3)
	robotgo.Sleep(30)
}

func (c *chaoJiDou) zhuqiaoMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.QiangZhanDaSha[1]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(435, 154, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(670, 28, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(131, 377, 1, 1, 3)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(270, 80, 3, 4)
	robotgo.MoveSmooth(270, 180, mouseSpeedX, mouseSpeedY)
	c.press(robotgo.KeyD, 1)
	c.move(270, 80, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 1)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.move(423, 436, 1, 4)
	c.press(robotgo.KeyF, 3)
	for i := 0; i < 5; i++ {
		c.press(robotgo.KeyD, 5)
	}
	c.press(robotgo.KeyT, 4)

	// 第6张怪物图： boss
	c.clickButton(fuben.SmallMap[5], 6)
	c.press(robotgo.Esc, 1)
	robotgo.Sleep(15)

	// 下一张图
	c.move(999, 293, 1, 5)
	c.press(robotgo.KeyF, 3)
	c.move(474, 228, 1, 4)
	c.press(robotgo.KeyF, 3)
	robotgo.Sleep(30)
}

func (c *chaoJiDou) youlingcangkuMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.QiangZhanDaSha[2]]
	c.multiMove(970, 190, 1, 1, 4)
	robotgo.Sleep(2)
	c.prestart()

	// 第1张怪物图
	c.press(robotgo.KeyF, 4)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	c.move(1142, 96, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1241, 636, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(746, 645, 1, 1, 3)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 1)
	c.move(42, 379, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 1)
	c.move(358, 174, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyT, 4)

	// 第7张怪物图： boss
	c.clickButton(fuben.SmallMap[6], 7)
	c.press(robotgo.Esc, 1)
	robotgo.Sleep(15)

	// 下一张图
	c.move(699, 166, 1, 4)
	c.press(robotgo.KeyF, 3)
	c.move(979, 482, 1, 4)
	c.press(robotgo.KeyF, 3)
	robotgo.Sleep(30)
}

func (c *chaoJiDou) dixiajidizuizhongjuezhanMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.QiangZhanDaSha[3]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(872, 645, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.F2, 1)
	c.move(1172, 87, 3, 4)
	robotgo.MoveSmooth(1172, 187, mouseSpeedX, mouseSpeedY)
	c.press(robotgo.KeyD, 1)
	c.move(952, 53, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 6)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(94, 741, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(925, 633, 1, 1, 2)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(40, 312, 3, 4)
	c.press(robotgo.KeyD, 1)
	c.move(40, 312, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 6)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(175, 104, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第7张怪物图： boss
	c.clickButton(fuben.SmallMap[6], 8)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(20)
}

func (c *chaoJiDou) zaicimianduiweixianMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.HeWeiAnYing[0]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(115, 107, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1273, 276, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 1)
	c.move(946, 132, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 1)
	c.move(1209, 161, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第7张怪物图
	c.clickButton(fuben.SmallMap[6], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第8张怪物图
	c.clickButton(fuben.SmallMap[7], 3)
	c.multiMove(148, 182, 1, 1, 18)
	robotgo.Sleep(2)

	// 第9张怪物图
	c.clickButton(fuben.SmallMap[8], 5)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 1)
	c.move(949, 126, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第10张怪物图
	c.clickButton(fuben.SmallMap[9], 3)
	c.multiMove(1168, 337, 1, 1, 18)
	robotgo.Sleep(2)

	// 第11张怪物图： boss
	c.clickButton(fuben.SmallMap[10], 5)
	robotgo.Sleep(15)

	// 下一张图
	c.move(773, 164, 1, 4)
	c.press(robotgo.KeyF, 3)
	robotgo.Sleep(30)
}

func (c *chaoJiDou) heimeiguitingyuanMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.HeWeiAnYing[1]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	c.move(1137, 69, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1043, 531, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1006, 162, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 1)
	c.move(503, 262, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyT, 4)

	// 第6张怪物图： boss
	c.clickButton(fuben.SmallMap[5], 4)
	c.press(robotgo.Esc, 1)
	robotgo.Sleep(15)

	// 下一张图
	c.move(997, 292, 1, 4)
	c.press(robotgo.KeyF, 3)
	c.move(572, 268, 1, 4)
	c.press(robotgo.KeyF, 3)
	robotgo.Sleep(30)
}

func (c *chaoJiDou) bolizhidiMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.HeWeiAnYing[2]]
	c.multiMove(84, 78, 1, 1, 5)
	robotgo.Sleep(2)
	c.prestart()

	// 第1张怪物图
	c.press(robotgo.KeyF, 4)
	c.move(443, 32, 3, 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(451, 31, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 11)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1014, 44, 3, 4)
	c.press(robotgo.KeyD, 1)
	c.move(1144, 141, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1156, 211, 3, 4)
	c.press(robotgo.KeyD, 1)
	c.move(1207, 685, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(936, 95, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 1)
	c.move(837, 79, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第6张怪物图： boss
	c.clickButton(fuben.SmallMap[5], 10)
	c.multiMove(686, 553, 1, 1, 3)
	robotgo.Sleep(45)

	// 下一张图
	c.move(1026, 286, 1, 4)
	c.press(robotgo.KeyF, 3)
	c.move(479, 241, 1, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	robotgo.Sleep(30)
}

func (c *chaoJiDou) ranheidewanxia1Meiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.HeWeiAnYing[3]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(217, 98, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.F2, 1)
	c.multiMove(924, 119, 1, 1, 4)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 6)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(853, 60, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(159, 52, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(141, 330, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyT, 4)

	// 第6张怪物图： boss
	c.clickButton(fuben.SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(20)
}

func (c *chaoJiDou) senlinbianjieMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.WeiXianHuiZhuan[0]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1047, 53, 3, 4)
	c.press(robotgo.KeyD, 1)
	c.move(1016, 125, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 5)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1159, 59, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(248, 69, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(127, 692, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyT, 4)

	// 第6张怪物图： boss
	c.clickButton(fuben.SmallMap[4], 8)
	robotgo.Sleep(20)

	// 下一张图
	c.move(787, 180, 1, 4)
	c.press(robotgo.KeyF, 3)
	//c.move(479, 241, 1, 1)
	//c.press(robotgo.KeyF, 1)
	//c.press(robotgo.KeyF, 1)
	//c.press(robotgo.KeyF, 1)
	robotgo.Sleep(30)
}

func (c *chaoJiDou) migongyuqiutuMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.WeiXianHuiZhuan[1]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1042, 65, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(436, 89, 3, 4)
	c.press(robotgo.KeyD, 1)
	c.move(563, 30, 3, 10)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1169, 693, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 1)
	c.move(1067, 140, 3, 8)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyT, 4)

	// 第6张怪物图： boss
	c.clickButton(fuben.SmallMap[5], 8)
	robotgo.Sleep(20)

	// 下一张图
	c.move(719, 162, 1, 4)
	c.press(robotgo.KeyF, 3)
	c.move(1024, 521, 1, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	robotgo.Sleep(30)
}

func (c *chaoJiDou) denglulalaiyeMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.WeiXianHuiZhuan[2]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(52, 282, 3, 4)
	c.press(robotgo.KeyD, 1)
	c.multiMove(210, 564, 1, 1, 6)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 6)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(138, 145, 3, 4)
	c.press(robotgo.KeyD, 3)
	c.move(233, 230, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 11)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(926, 44, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 1)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(602, 52, 3, 4)
	c.press(robotgo.KeyD, 1)
	c.move(1319, 487, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第7张怪物图
	c.clickButton(fuben.SmallMap[6], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1101, 137, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第8张怪物图： boss
	c.clickButton(fuben.SmallMap[7], 10)
	robotgo.Sleep(20)

	// 下一张图
	c.move(966, 269, 1, 4)
	c.press(robotgo.KeyF, 3)
	c.move(367, 317, 1, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	robotgo.Sleep(30)
}

func (c *chaoJiDou) luosudedixiashijieMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.WeiXianHuiZhuan[3]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1084, 44, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1142, 376, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 6)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(998, 209, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(700, 104, 1, 1, 6)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 3)
	c.clickButton(fuben.SmallMap[5], 3)
	c.clickButton(fuben.SmallMap[5], 6)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(818, 53, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyT, 4)

	// 第7张怪物图: boss
	c.clickButton(fuben.SmallMap[6], 3)
	c.clickButton(fuben.SmallMap[6], 3)
	c.clickButton(fuben.SmallMap[6], 6)
	robotgo.Sleep(20)
}

func (c *chaoJiDou) haishangqiaoliangMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.WangMingTuDeDaoLu[0]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1095, 109, 1, 1, 3)
	robotgo.Sleep(2)
	robotgo.MoveSmooth(1095, 209, mouseSpeedX, mouseSpeedY)
	c.press(robotgo.KeyD, 1)
	c.multiMove(1095, 109, 1, 1, 3)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1140, 80, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1268, 363, 1, 1, 3)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1155, 85, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图： boss
	c.press(robotgo.KeyT, 4)
	c.clickButton(fuben.SmallMap[4], 8)
	robotgo.Sleep(20)

	// 下一张图
	c.move(1117, 392, 1, 4)
	c.press(robotgo.KeyF, 3)
	//c.move(1024, 521, 1, 1)
	//c.press(robotgo.KeyF, 1)
	//c.press(robotgo.KeyF, 1)
	//c.press(robotgo.KeyF, 1)
	robotgo.Sleep(25)
}

func (c *chaoJiDou) heihaiMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.WangMingTuDeDaoLu[1]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1150, 67, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1255, 401, 1, 1, 3)
	robotgo.Sleep(2)
	robotgo.MoveSmooth(1255, 501, mouseSpeedX, mouseSpeedY)
	c.press(robotgo.KeyD, 1)
	c.multiMove(1255, 401, 1, 1, 2)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1116, 115, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1136, 67, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 1)
	c.move(1239, 381, 3, 8)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第6张怪物图： boss
	c.press(robotgo.KeyT, 4)
	c.clickButton(fuben.SmallMap[5], 8)
	robotgo.Sleep(20)

	// 下一张图
	c.move(1121, 408, 1, 4)
	c.press(robotgo.KeyF, 3)
	c.move(670, 214, 1, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	robotgo.Sleep(25)
}

func (c *chaoJiDou) langzimoluMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.WangMingTuDeDaoLu[2]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1129, 77, 3, 4)
	c.press(robotgo.KeyD, 1)
	c.move(830, 49, 3, 4)
	c.press(robotgo.KeyD, 1)
	c.multiMove(918, 60, 1, 1, 3)
	robotgo.Sleep(2)
	c.move(1158, 154, 3, 5)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1250, 150, 1, 1, 3)
	robotgo.Sleep(2)
	robotgo.MoveSmooth(1250, 250, mouseSpeedX, mouseSpeedY)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 6)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(896, 52, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1281, 256, 1, 1, 4)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	c.move(1020, 102, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1152, 107, 1, 1, 6)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第6张怪物图: boss
	c.press(robotgo.KeyT, 4)
	c.clickButton(fuben.SmallMap[5], 8)
	robotgo.Sleep(20)
}

func (c *chaoJiDou) hongseshamoMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.ShenMiDongWuXue[0]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(540, 47, 3, 4)
	c.press(robotgo.KeyD, 1)
	c.move(127, 152, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(54, 421, 3, 4)
	c.press(robotgo.KeyD, 1)
	c.multiMove(166, 645, 1, 1, 4)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(103, 425, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(919, 623, 1, 1, 2)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	c.multiMove(233, 623, 1, 1, 2)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	c.multiMove(110, 444, 1, 1, 2)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	c.move(368, 182, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(82, 371, 1, 1, 5)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	c.multiMove(82, 371, 1, 1, 5)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(108, 349, 1, 1, 4)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	c.multiMove(653, 611, 1, 1, 2)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyT, 4)

	// 第7张怪物图: boss
	c.clickButton(fuben.SmallMap[6], 8)
	robotgo.Sleep(20)

	// 下一张图
	c.move(755, 190, 1, 4)
	c.press(robotgo.KeyF, 3)
	//c.move(367, 317, 1, 1)
	//c.press(robotgo.KeyF, 1)
	//c.press(robotgo.KeyF, 1)
	//c.press(robotgo.KeyF, 1)
	robotgo.Sleep(30)
}

func (c *chaoJiDou) xunzhaochuluMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.ShenMiDongWuXue[1]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1165, 360, 1, 1, 2)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1149, 558, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(897, 235, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.move(691, 99, 1, 4)
	c.press(robotgo.KeyF, 1)
	robotgo.Sleep(8)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(876, 101, 1, 1, 5)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	c.multiMove(742, 640, 1, 1, 6)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1105, 70, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第7张怪物图
	c.clickButton(fuben.SmallMap[6], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1028, 579, 1, 1, 2)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	c.multiMove(91, 694, 1, 1, 2)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	c.multiMove(1169, 621, 1, 1, 4)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	c.multiMove(1080, 86, 1, 1, 2)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	c.multiMove(1022, 592, 1, 1, 4)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第8张怪物图
	c.clickButton(fuben.SmallMap[7], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1220, 628, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第9张怪物图
	c.clickButton(fuben.SmallMap[8], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	robotgo.Sleep(3)
	c.move(1014, 117, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyT, 4)

	// 第10张怪物图: boss
	c.clickButton(fuben.SmallMap[9], 8)
	robotgo.Sleep(20)

	// 下一张图
	c.move(933, 277, 1, 4)
	c.press(robotgo.KeyF, 3)
	//c.move(367, 317, 1, 1)
	//c.press(robotgo.KeyF, 1)
	//c.press(robotgo.KeyF, 1)
	//c.press(robotgo.KeyF, 1)
	robotgo.Sleep(30)
}

func (c *chaoJiDou) yanhanxueshanMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.ShenMiDongWuXue[2]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(187, 93, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1195, 248, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1245, 693, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(719, 57, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1281, 80, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyT, 4)

	// 第7张怪物图: boss
	c.clickButton(fuben.SmallMap[6], 8)
	robotgo.Sleep(20)

	// 下一张图
	c.move(785, 159, 1, 4)
	c.press(robotgo.KeyF, 3)
	c.move(942, 579, 1, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	robotgo.Sleep(30)
}

func (c *chaoJiDou) bingxuefugaizhiqiangMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.ShenMiDongWuXue[3]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1011, 189, 1, 1, 15)
	robotgo.Sleep(3)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 5)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1118, 425, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 6)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(754, 42, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1192, 626, 1, 1, 2)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(664, 656, 1, 1, 2)
	robotgo.Sleep(8)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1214, 452, 1, 1, 2)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第7张怪物图: boss
	c.press(robotgo.KeyT, 4)
	c.clickButton(fuben.SmallMap[6], 8)
	robotgo.Sleep(20)
}

func (c *chaoJiDou) dixiajidijinjizuozhanMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.HuoZhiJiDian[0]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1106, 359, 1, 1, 3)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(613, 146, 1, 1, 3)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(892, 155, 1, 1, 2)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1020, 215, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1107, 133, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1090, 654, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第7张怪物图: boss
	c.press(robotgo.KeyT, 4)
	c.clickButton(fuben.SmallMap[6], 8)
	robotgo.Sleep(20)

	// 下一张图
	c.move(842, 202, 1, 4)
	c.press(robotgo.KeyF, 3)
	//c.move(942, 579, 1, 1)
	//c.press(robotgo.KeyF, 1)
	//c.press(robotgo.KeyF, 1)
	//c.press(robotgo.KeyF, 1)
	robotgo.Sleep(30)
}

func (c *chaoJiDou) huoyanzhuisuizheMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.HuoZhiJiDian[1]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1111, 82, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(842, 128, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(275, 149, 1, 1, 2)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1063, 75, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1260, 252, 1, 1, 4)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)

	// 第6张怪物图: boss
	c.press(robotgo.KeyT, 4)
	c.clickButton(fuben.SmallMap[5], 8)
	robotgo.Sleep(20)

	// 下一张图
	c.move(1064, 347, 1, 4)
	c.press(robotgo.KeyF, 3)
	c.move(555, 219, 1, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	robotgo.Sleep(30)
}

func (c *chaoJiDou) feiqirongluMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.HuoZhiJiDian[2]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1111, 82, 3, 4)
	c.press(robotgo.KeyD, 1)
	c.move(1061, 134, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1147, 540, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.multiMove(783, 59, 1, 1, 30)
	robotgo.Sleep(2)
	c.press(robotgo.KeyT, 4)
	c.press(robotgo.KeyW, 1)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1042, 154, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(725, 102, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(6)

	// 第6张怪物图: 无boss
	c.clickButton(fuben.SmallMap[5], 8)
	c.multiMove(408, 194, 1, 1, 12)
	c.multiMove(851, 150, 1, 1, 10)
	c.press(robotgo.F1, 1)
	c.multiMove(987, 627, 1, 1, 18)
	c.press(robotgo.F1, 1)
	c.multiMove(1167, 189, 1, 1, 10)
	c.press(robotgo.F1, 1)
	c.multiMove(483, 35, 1, 1, 12)
	c.multiMove(881, 63, 1, 1, 3)
	robotgo.Sleep(10)

	// 下一张图
	c.move(541, 590, 1, 4)
	c.press(robotgo.KeyF, 3)
	c.move(541, 590, 1, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	c.press(robotgo.KeyF, 1)
	robotgo.Sleep(30)
}

func (c *chaoJiDou) huoyanzhinvMeiri() {
	fuben := c.JuQingMap.FuBens[c.MeiRiMap.HuoZhiJiDian[3]]
	c.prestart()

	// 第1张怪物图
	c.clickButton(fuben.SmallMap[0], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1067, 250, 1, 1, 10)
	robotgo.Sleep(1)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第2张怪物图
	c.clickButton(fuben.SmallMap[1], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(702, 637, 1, 1, 3)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第3张怪物图
	c.clickButton(fuben.SmallMap[2], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(1226, 620, 1, 1, 3)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	c.move(989, 185, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第4张怪物图
	c.clickButton(fuben.SmallMap[3], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1273, 644, 3, 4)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第5张怪物图
	c.clickButton(fuben.SmallMap[4], 3)
	c.clickButton(fuben.SmallMap[4], 3)
	c.clickButton(fuben.SmallMap[4], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.move(1054, 130, 3, 4)
	c.press(robotgo.KeyS, 3)
	c.multiMove(1054, 130, 1, 1, 3)
	robotgo.Sleep(2)
	c.press(robotgo.KeyD, 1)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(2)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.KeyPress(robotgo.KeyS)
	robotgo.Sleep(4)

	// 第6张怪物图
	c.clickButton(fuben.SmallMap[5], 8)
	c.press(robotgo.KeyD, 1)
	c.press(robotgo.Key3, 2)
	c.multiMove(977, 49, 1, 1, 4)
	robotgo.Sleep(3)
	c.press(robotgo.KeyD, 1)
	robotgo.Sleep(4)

	// 第7张怪物图: boss
	c.press(robotgo.KeyT, 4)
	c.clickButton(fuben.SmallMap[6], 8)
	robotgo.Sleep(20)
}
