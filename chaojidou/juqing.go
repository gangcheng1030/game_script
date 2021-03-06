package chaojidou

import "C"
import "github.com/go-vgo/robotgo"

type JuQingType string

const (
	JUQING_TYPE_YISHENGYUZHAOZE          = "ysyzz"    // 医生与沼泽
	JUQING_TYPE_SHENHAIYIZU              = "shyz"     // 深海遗族
	JUQING_TYPE_SHIJUNZHIZUI             = "sjzz"     // 弑君之罪
	JUQING_TYPE_TONGXIANGWEIYIJIE        = "txwyj"    // 同相位异界
	JUQING_TYPE_YEXINZHONGHE             = "yxzh"     // 野心重荷
	JUQING_TYPE_FENGYINZHISHI            = "fyzs"     // 封印之石
	JUQING_TYPE_YINGXIONGDEZIGE          = "yxdzg"    // 英雄的资格
	JUQING_TYPE_TIANHE                   = "th"       // 天河
	JUQING_TYPE_TIANKONGDAPINGYUAN       = "tkdpy"    // 天空大平原
	JUQING_TYPE_XUNNANGUANGCHANGRUKOU    = "xngcrk"   // 殉难广场入口
	JUQING_TYPE_BANGONGXIUXISHI          = "bgxxs"    // 办公休息室
	JUQING_TYPE_TIANBIANYUNXIA           = "tbyx"     // 天边云霞
	JUQING_TYPE_ZHONGJIBINGQI            = "zjbq"     // 终极兵器
	JUQING_TYPE_QUANJIANZUOZHAN          = "qjzz"     // 全歼作战
	JUQING_TYPE_QTSHAMOKAIFABU           = "qtsmkfb"  // QT沙漠开发部
	JUQING_TYPE_TIANKONGZHIZHUAN         = "tkzz"     // 天空之钻
	JUQING_TYPE_XINTIANEBAO              = "xteb"     // 新天鹅堡
	JUQING_TYPE_BINGHE                   = "bh"       // 冰河
	JUQING_TYPE_BINGXUEBULUOMUBACUN      = "bxblmbc"  // 冰雪部落姆巴村
	JUQING_TYPE_BUGUIXIAGU               = "bgxg"     // 不归峡谷
	JUQING_TYPE_BINGFENGCHENGBAO         = "bfcb"     // 冰封城堡
	JUQING_TYPE_JIANRUIXIAGU             = "jrxg"     // 尖锐峡谷
	JUQING_TYPE_YUEGUANGDADAO            = "ygdd"     // 月光大道
	JUQING_TYPE_SHIXIANZHENGYI           = "sxzy"     // 实现正义
	JUQING_TYPE_SHIJIANTINGYUAN          = "sjty"     // 时间庭院
	JUQING_TYPE_HUWEIDUIXIONGMEI         = "hwdxm"    // 护卫队兄妹
	JUQING_TYPE_TIANCAISHAONV            = "tcsn"     // 天才少女
	JUQING_TYPE_QINGZHULIN               = "qzl"      // 青竹林
	JUQING_TYPE_EMEIBIEYUAN              = "emby"     // 峨眉别院
	JUQING_TYPE_GAIYAN                   = "gy"       // 丐岩
	JUQING_TYPE_BEIMANGSHAN              = "bms"      // 北邙山
	JUQING_TYPE_BIWUDASAIJUESAI          = "bwdsjs"   // 比武大赛决赛
	JUQING_TYPE_KONGZHONGANXI            = "kzax"     // 空中暗袭
	JUQING_TYPE_DIXIATINGCHECHANG        = "dxtcc"    // 地下停车场
	JUQING_TYPE_NIZONGSHENSHESHOU        = "nzsss"    // 匿踪神射手
	JUQING_TYPE_YOULING                  = "yl"       // 幽灵
	JUQING_TYPE_DIXIAJIDISHIYANSHI       = "dxjdsys"  // 地下基地：实验室
	JUQING_TYPE_SHALEDUONAIHUOSHANKOU    = "jldnhsk"  // 沙勒多奈火山口
	JUQING_TYPE_HAIDELUSHIYANSHI         = "hdlsys"   // 海德鲁实验室
	JUQING_TYPE_KAISALIYAGANG            = "kslyg"    // 凯撒里亚港
	JUQING_TYPE_ZHUQIAO                  = "zq"       // 主桥
	JUQING_TYPE_YOULINGCANGKU            = "ylck"     // 幽灵仓库
	JUQING_TYPE_DIXIAJIDIZUIZHONGJUEZHAN = "dxjdzzjz" // 地下基地：最终决战
	JUQING_TYPE_ZAICIMIANDUIWEIXIAN      = "zcmdwx"   // 再次面对危险
	JUQING_TYPE_HEIMEIGUITINGYUAN        = "hmgty"    // 黑玫瑰庭院
	JUQING_TYPE_BOLIZHIDI                = "blzd"     // 剥离之地
	JUQING_TYPE_RANHEIDEWANXIA1          = "rhdwx1"   // 染黑的晚霞1
	JUQING_TYPE_SENLINBIANJIE            = "slbj"     // 森林边界
	JUQING_TYPE_MIGONGYUQIUTU            = "mgyqt"    // 迷宫与囚徒
	JUQING_TYPE_DENGLULALAIYE            = "dllly"    // 登陆拉莱耶
	JUQING_TYPE_LUOSUDEDIXIASHIJIE       = "lsddxsj"  // 罗素的地下世界
	JUQING_TYPE_HAISHANGQIAOLIANG        = "hsql"     // 海上桥梁
	JUQING_TYPE_HEIHAI                   = "hh"       // 黑海
	JUQING_TYPE_LANGZIMOLU               = "lzml"     // 浪子末路
	JUQING_TYPE_HONGSESHAMO              = "hssm"     // 红色沙漠
	JUQING_TYPE_XUNZHAOCHULU             = "xzcl"     // 寻找出路
	JUQING_TYPE_YANHANXUESHAN            = "yhxs"     // 严寒雪山
	JUQING_TYPE_BINGXUEFUGAIZHIQIANG     = "bxfgzq"   // 冰雪覆盖之墙
	JUQING_TYPE_DIXIAJIDIJINJIZUOZHAN    = "dxjdjjzz" // 地下基地：紧急作战
	JUQING_TYPE_HUOYANZHUISUIZHE         = "hyzsz"    // 火焰追随者
	JUQING_TYPE_FEIQIRONGLU              = "fqrl"     // 废弃熔炉
	JUQING_TYPE_HUOYANZHINV              = "hyzn"     // 火焰之女
	JUQING_TYPE_DAYANGSHOUHUZHE          = "dyshz"    // 大洋守护者
	JUQING_TYPE_QIUZHUXINHAO             = "qzxh"     // 求助信号
	JUQING_TYPE_MENGMAMUDI               = "mmmd"     // 猛犸墓地
	JUQING_TYPE_DUNYUESHENGDI            = "dysd"     // 钝月圣地
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
