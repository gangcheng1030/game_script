package chaojidou

import (
	"github.com/gangcheng1030/game_script/utils"
	"github.com/go-vgo/robotgo"
)

type OfficialClientChaoJiDou struct {
	chaoJiDou
}

func NewOfficialClientChaoJiDou() ChaoJiDou {
	pss, _ := utils.FindProcessByName(OFFICIALCLIENT_PROCESS_NAME)
	c := chaoJiDou{
		Pid: pss.Pid,
		GameWindow: robotgo.Rect{
			Point: robotgo.Point{X: 0, Y: 0},
			Size:  robotgo.Size{W: 1360, H: 768},
		},
		BigMap: BigMap{
			ZhuiSu: robotgo.Rect{
				Point: robotgo.Point{X: 635, Y: 322},
				Size:  robotgo.Size{W: 7, H: 7},
			},
			LiuLangTuan: robotgo.Rect{
				Point: robotgo.Point{X: 765, Y: 515},
				Size:  robotgo.Size{W: 7, H: 7},
			},
			Hermosi: robotgo.Rect{
				Point: robotgo.Point{X: 1121, Y: 433},
				Size:  robotgo.Size{W: 7, H: 7},
			},
			MeiRi: robotgo.Rect{
				Point: robotgo.Point{X: 736, Y: 461},
				Size:  robotgo.Size{W: 4, H: 4},
			},
		},
		RenWuMap: RenWuMap{
			ZhuYe: robotgo.Rect{
				Point: robotgo.Point{X: 631, Y: 126},
				Size:  robotgo.Size{W: 20, H: 2},
			},
			FirstXunLu: robotgo.Rect{
				Point: robotgo.Point{X: 901, Y: 162},
				Size:  robotgo.Size{W: 2, H: 2},
			},
		},
		JuQingMap: JuQingMap{
			TongHuaZhenButton: robotgo.Rect{
				Point: robotgo.Point{X: 458, Y: 367},
				Size:  robotgo.Size{W: 20, H: 20},
			},
			FuBens: map[JuQingType]FuBen{
				JUQING_TYPE_YISHENGYUZHAOZE: {
					Window: robotgo.Rect{
						Point: robotgo.Point{X: 930, Y: 312},
						Size:  robotgo.Size{W: 20, H: 10},
					},
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1207, Y: 137},
							Size:  robotgo.Size{W: 6, H: 6},
						},
						{
							Point: robotgo.Point{X: 1237, Y: 137},
							Size:  robotgo.Size{W: 6, H: 6},
						},
						{
							Point: robotgo.Point{X: 1237, Y: 107},
							Size:  robotgo.Size{W: 6, H: 6},
						},
						{
							Point: robotgo.Point{X: 1267, Y: 107},
							Size:  robotgo.Size{W: 6, H: 6},
						},
						{
							Point: robotgo.Point{X: 1267, Y: 137},
							Size:  robotgo.Size{W: 6, H: 6},
						},
						{
							Point: robotgo.Point{X: 1297, Y: 137},
							Size:  robotgo.Size{W: 6, H: 6},
						},
						{
							Point: robotgo.Point{X: 1297, Y: 107},
							Size:  robotgo.Size{W: 6, H: 6},
						},
					},
				},
				JUQING_TYPE_SHENHAIYIZU: {
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1234, Y: 170},
							Size:  robotgo.Size{W: 2, H: 2},
						},
						{
							Point: robotgo.Point{X: 1234, Y: 149},
							Size:  robotgo.Size{W: 2, H: 2},
						},
						{
							Point: robotgo.Point{X: 1234, Y: 128},
							Size:  robotgo.Size{W: 2, H: 2},
						},
						{
							Point: robotgo.Point{X: 1255, Y: 128},
							Size:  robotgo.Size{W: 2, H: 2},
						},
						{
							Point: robotgo.Point{X: 1276, Y: 128},
							Size:  robotgo.Size{W: 2, H: 2},
						},
						{
							Point: robotgo.Point{X: 1276, Y: 107},
							Size:  robotgo.Size{W: 2, H: 2},
						},
						{
							Point: robotgo.Point{X: 1276, Y: 86},
							Size:  robotgo.Size{W: 2, H: 2},
						},
						{
							Point: robotgo.Point{X: 1276, Y: 65},
							Size:  robotgo.Size{W: 2, H: 2},
						},
					},
				},
				JUQING_TYPE_SHIJUNZHIZUI: {
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1270, Y: 141},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1270, Y: 111},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1270, Y: 81},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1240, Y: 81},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1240, Y: 111},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1240, Y: 141},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1240, Y: 171},
							Size:  robotgo.Size{W: 4, H: 4},
						},
					},
				},
				JUQING_TYPE_TONGXIANGWEIYIJIE: {
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1225, Y: 127},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1225, Y: 157},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1255, Y: 157},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1285, Y: 157},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1285, Y: 127},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1285, Y: 97},
							Size:  robotgo.Size{W: 4, H: 4},
						},
					},
				},
				JUQING_TYPE_YEXINZHONGHE: {
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1255, Y: 97},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1225, Y: 97},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1225, Y: 127},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1225, Y: 157},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1255, Y: 157},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1285, Y: 157},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1285, Y: 127},
							Size:  robotgo.Size{W: 4, H: 4},
						},
					},
				},
				JUQING_TYPE_FENGYINZHISHI: {
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1255, Y: 81},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1225, Y: 81},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1225, Y: 111},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1225, Y: 141},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1255, Y: 141},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1285, Y: 141},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1285, Y: 171},
							Size:  robotgo.Size{W: 4, H: 4},
						},
					},
				},
				JUQING_TYPE_YINGXIONGDEZIGE: {
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1270, Y: 171},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1270, Y: 141},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1270, Y: 111},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1270, Y: 81},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1240, Y: 81},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1240, Y: 111},
							Size:  robotgo.Size{W: 4, H: 4},
						},
					},
				},
				JUQING_TYPE_TIANHE: {
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1225, Y: 158},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1255, Y: 158},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1255, Y: 128},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1255, Y: 98},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1255, Y: 68},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1285, Y: 68},
							Size:  robotgo.Size{W: 4, H: 4},
						},
					},
				},
				JUQING_TYPE_TIANKONGDAPINGYUAN: {
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1227, Y: 165},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1227, Y: 140},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1252, Y: 140},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1277, Y: 140},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1277, Y: 115},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1277, Y: 90},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1277, Y: 65},
							Size:  robotgo.Size{W: 4, H: 4},
						},
					},
				},
				JUQING_TYPE_XUNNANGUANGCHANGRUKOU: {
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1225, Y: 158},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1225, Y: 128},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1255, Y: 128},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1285, Y: 128},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1285, Y: 98},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1285, Y: 68},
							Size:  robotgo.Size{W: 4, H: 4},
						},
					},
				},
				JUQING_TYPE_BANGONGXIUXISHI: {
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1255, Y: 158},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1285, Y: 158},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1285, Y: 128},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1255, Y: 128},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1255, Y: 98},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1225, Y: 98},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1225, Y: 68},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1255, Y: 68},
							Size:  robotgo.Size{W: 4, H: 4},
						},
					},
				},
				JUQING_TYPE_TIANBIANYUNXIA: {
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1225, Y: 141},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1225, Y: 111},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1225, Y: 81},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1255, Y: 81},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1255, Y: 111},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1285, Y: 111},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1285, Y: 81},
							Size:  robotgo.Size{W: 4, H: 4},
						},
					},
				},
				JUQING_TYPE_ZHONGJIBINGQI: {
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1270, Y: 141},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1270, Y: 111},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1300, Y: 111},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1300, Y: 81},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1270, Y: 81},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1240, Y: 81},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1210, Y: 81},
							Size:  robotgo.Size{W: 4, H: 4},
						},
					},
				},
				JUQING_TYPE_QUANJIANZUOZHAN: {
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1270, Y: 171},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1270, Y: 141},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1240, Y: 141},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1240, Y: 111},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1240, Y: 81},
							Size:  robotgo.Size{W: 4, H: 4},
						},
					},
				},
				JUQING_TYPE_QTSHAMOKAIFABU: {
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1225, Y: 141},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1225, Y: 111},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1255, Y: 111},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1255, Y: 141},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1285, Y: 141},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1285, Y: 111},
							Size:  robotgo.Size{W: 4, H: 4},
						},
					},
				},
				JUQING_TYPE_TIANKONGZHIZHUAN: {
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1255, Y: 111},
							Size:  robotgo.Size{W: 4, H: 4},
						},
					},
				},
				JUQING_TYPE_XINTIANEBAO: {
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1225, Y: 158},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1255, Y: 158},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1285, Y: 158},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1285, Y: 128},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1255, Y: 128},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1255, Y: 98},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1225, Y: 98},
							Size:  robotgo.Size{W: 4, H: 4},
						},
					},
				},
				JUQING_TYPE_BINGHE: {
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1225, Y: 141},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1255, Y: 141},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1285, Y: 141},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1285, Y: 111},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1285, Y: 81},
							Size:  robotgo.Size{W: 4, H: 4},
						},
					},
				},
				JUQING_TYPE_BINGXUEBULUOMUBACUN: {
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1255, Y: 158},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1255, Y: 128},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1255, Y: 98},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1225, Y: 98},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1225, Y: 68},
							Size:  robotgo.Size{W: 4, H: 4},
						},
					},
				},
				JUQING_TYPE_BUGUIXIAGU: {
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1205, Y: 155},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1235, Y: 155},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1265, Y: 155},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1265, Y: 125},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1265, Y: 95},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1265, Y: 65},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1295, Y: 65},
							Size:  robotgo.Size{W: 4, H: 4},
						},
					},
				},
				JUQING_TYPE_BINGFENGCHENGBAO: {
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1300, Y: 141},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1300, Y: 111},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1270, Y: 111},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1270, Y: 81},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1240, Y: 81},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1240, Y: 111},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1210, Y: 111},
							Size:  robotgo.Size{W: 4, H: 4},
						},
					},
				},
				JUQING_TYPE_JIANRUIXIAGU: {
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1285, Y: 158},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1255, Y: 158},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1225, Y: 158},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1225, Y: 128},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1225, Y: 98},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1255, Y: 98},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1285, Y: 98},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1285, Y: 68},
							Size:  robotgo.Size{W: 4, H: 4},
						},
					},
				},
				JUQING_TYPE_YUEGUANGDADAO: {
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1240, Y: 161},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1240, Y: 136},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1240, Y: 111},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1240, Y: 86},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1265, Y: 86},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1265, Y: 61},
							Size:  robotgo.Size{W: 4, H: 4},
						},
					},
				},
				JUQING_TYPE_SHIXIANZHENGYI: {
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1285, Y: 158},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1285, Y: 128},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1285, Y: 98},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1255, Y: 98},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1255, Y: 128},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1225, Y: 128},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1195, Y: 128},
							Size:  robotgo.Size{W: 4, H: 4},
						},
					},
				},
				JUQING_TYPE_SHIJIANTINGYUAN: {
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1295, Y: 125},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1295, Y: 95},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1265, Y: 95},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1265, Y: 125},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1235, Y: 125},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1205, Y: 125},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1205, Y: 95},
							Size:  robotgo.Size{W: 4, H: 4},
						},
					},
				},
				JUQING_TYPE_HUWEIDUIXIONGMEI: {
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1255, Y: 158},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1225, Y: 158},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1225, Y: 128},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1255, Y: 128},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1255, Y: 98},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1285, Y: 98},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1285, Y: 68},
							Size:  robotgo.Size{W: 4, H: 4},
						},
					},
				},
				JUQING_TYPE_TIANCAISHAONV: {
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1205, Y: 155},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1235, Y: 155},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1235, Y: 125},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1235, Y: 95},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1265, Y: 95},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1295, Y: 95},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1295, Y: 65},
							Size:  robotgo.Size{W: 4, H: 4},
						},
					},
				},
				JUQING_TYPE_QINGZHULIN: {
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1270, Y: 171},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1270, Y: 141},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1270, Y: 111},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1240, Y: 111},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1240, Y: 141},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1210, Y: 141},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1210, Y: 111},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1210, Y: 81},
							Size:  robotgo.Size{W: 4, H: 4},
						},
					},
				},
				JUQING_TYPE_EMEIBIEYUAN: {
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1255, Y: 158},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1285, Y: 158},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1285, Y: 128},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1255, Y: 128},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1255, Y: 98},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1255, Y: 68},
							Size:  robotgo.Size{W: 4, H: 4},
						},
					},
				},
				JUQING_TYPE_GAIYAN: {
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1270, Y: 171},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1270, Y: 141},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1240, Y: 141},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1240, Y: 111},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1240, Y: 81},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1270, Y: 81},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1300, Y: 81},
							Size:  robotgo.Size{W: 4, H: 4},
						},
					},
				},
				JUQING_TYPE_BEIMANGSHAN: {
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1225, Y: 158},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1225, Y: 128},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1255, Y: 128},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1285, Y: 128},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1315, Y: 128},
							Size:  robotgo.Size{W: 4, H: 4},
						},
						{
							Point: robotgo.Point{X: 1315, Y: 98},
							Size:  robotgo.Size{W: 4, H: 4},
						},
					},
				},
				JUQING_TYPE_BIWUDASAIJUESAI: {
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1255, Y: 111},
							Size:  robotgo.Size{W: 4, H: 4},
						},
					},
				},
			},
		},
		ZhuiSuMap: ZhuiSuMap{
			EnterButton: robotgo.Rect{
				Point: robotgo.Point{X: 76, Y: 714},
				Size:  robotgo.Size{W: 100, H: 10},
			},
			JiuYunDong: FuBen{
				Window: robotgo.Rect{
					Point: robotgo.Point{X: 440, Y: 410},
					Size:  robotgo.Size{W: 180, H: 50},
				},
				DifficultyTypePoses: []robotgo.Rect{
					{
						Point: robotgo.Point{X: 480, Y: 355},
						Size:  robotgo.Size{W: 40, H: 10},
					},
					{
						Point: robotgo.Point{X: 545, Y: 355},
						Size:  robotgo.Size{W: 40, H: 10},
					},
				},
				SmallMap: []robotgo.Rect{
					{
						Point: robotgo.Point{X: 1250, Y: 160},
						Size:  robotgo.Size{W: 10, H: 10},
					},
					{
						Point: robotgo.Point{X: 1275, Y: 160},
						Size:  robotgo.Size{W: 10, H: 10},
					},
					{
						Point: robotgo.Point{X: 1275, Y: 135},
						Size:  robotgo.Size{W: 10, H: 10},
					},
					{
						Point: robotgo.Point{X: 1275, Y: 110},
						Size:  robotgo.Size{W: 10, H: 10},
					},
					{
						Point: robotgo.Point{X: 1250, Y: 110},
						Size:  robotgo.Size{W: 10, H: 10},
					},
					{
						Point: robotgo.Point{X: 1225, Y: 110},
						Size:  robotgo.Size{W: 10, H: 10},
					},
					{
						Point: robotgo.Point{X: 1225, Y: 85},
						Size:  robotgo.Size{W: 10, H: 10},
					},
					{
						Point: robotgo.Point{X: 1225, Y: 60},
						Size:  robotgo.Size{W: 10, H: 10},
					},
				},
			},
		},
		MeiRiMap: MeiRiMap{
			Window: robotgo.Rect{
				Point: robotgo.Point{X: 629, Y: 319},
				Size:  robotgo.Size{W: 100, H: 20},
			},
			ShenHaiZhuKe:    []JuQingType{JUQING_TYPE_SHENHAIYIZU, JUQING_TYPE_SHIJUNZHIZUI, JUQING_TYPE_TONGXIANGWEIYIJIE, JUQING_TYPE_YEXINZHONGHE},
			ZuiQiangJianShi: []JuQingType{JUQING_TYPE_FENGYINZHISHI, JUQING_TYPE_YINGXIONGDEZIGE, JUQING_TYPE_TIANHE, JUQING_TYPE_TIANKONGDAPINGYUAN},
			TieBiShouWei:    []JuQingType{JUQING_TYPE_XUNNANGUANGCHANGRUKOU, JUQING_TYPE_BANGONGXIUXISHI, JUQING_TYPE_TIANBIANYUNXIA, JUQING_TYPE_ZHONGJIBINGQI},
			YanHua:          []JuQingType{JUQING_TYPE_QUANJIANZUOZHAN, JUQING_TYPE_QTSHAMOKAIFABU, JUQING_TYPE_TIANKONGZHIZHUAN, JUQING_TYPE_XINTIANEBAO},
			BingFengWangGuo: []JuQingType{JUQING_TYPE_BINGHE, JUQING_TYPE_BINGXUEBULUOMUBACUN, JUQING_TYPE_BUGUIXIAGU, JUQING_TYPE_BINGFENGCHENGBAO},
			XiaoXinChuDian:  []JuQingType{JUQING_TYPE_JIANRUIXIAGU, JUQING_TYPE_YUEGUANGDADAO, JUQING_TYPE_SHIXIANZHENGYI},
			RenXiaoGuiDa:    []JuQingType{JUQING_TYPE_SHIJIANTINGYUAN, JUQING_TYPE_HUWEIDUIXIONGMEI, JUQING_TYPE_TIANCAISHAONV},
			JiuYunZhiDian:   []JuQingType{JUQING_TYPE_QINGZHULIN, JUQING_TYPE_EMEIBIEYUAN, JUQING_TYPE_GAIYAN, JUQING_TYPE_BEIMANGSHAN, JUQING_TYPE_BIWUDASAIJUESAI},
		},
		LiuLangTuanMap: LiuLangTuanMap{
			EnterButton: robotgo.Rect{
				Point: robotgo.Point{X: 76, Y: 714},
				Size:  robotgo.Size{W: 100, H: 10},
			},
			FuBenArray: []FuBen{
				{
					Window: robotgo.Rect{
						Point: robotgo.Point{X: 445, Y: 335},
						Size:  robotgo.Size{W: 180, H: 50},
					},
					DifficultyTypePoses: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 420, Y: 280},
							Size:  robotgo.Size{W: 40, H: 10},
						},
						{
							Point: robotgo.Point{X: 480, Y: 280},
							Size:  robotgo.Size{W: 40, H: 10},
						},
						{
							Point: robotgo.Point{X: 540, Y: 280},
							Size:  robotgo.Size{W: 40, H: 10},
						},
					},
				},
			},
		},
		JinBenMap: JinBenMap{
			EnterButton: robotgo.Rect{
				Point: robotgo.Point{X: 76, Y: 714},
				Size:  robotgo.Size{W: 100, H: 10},
			},
			ChuanShaoButton: robotgo.Rect{
				Point: robotgo.Point{X: 900, Y: 375},
				Size:  robotgo.Size{W: 200, H: 70},
			},
			FuBenArray: []FuBen{
				{
					Window: robotgo.Rect{
						Point: robotgo.Point{X: 515, Y: 385},
						Size:  robotgo.Size{W: 180, H: 45},
					},
					DifficultyTypePoses: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 490, Y: 330},
							Size:  robotgo.Size{W: 40, H: 10},
						},
						{
							Point: robotgo.Point{X: 553, Y: 330},
							Size:  robotgo.Size{W: 40, H: 10},
						},
						{
							Point: robotgo.Point{X: 616, Y: 330},
							Size:  robotgo.Size{W: 40, H: 10},
						},
						{
							Point: robotgo.Point{X: 679, Y: 330},
							Size:  robotgo.Size{W: 40, H: 10},
						},
					},
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1250, Y: 90},
							Size:  robotgo.Size{W: 10, H: 10},
						},
						{
							Point: robotgo.Point{X: 1220, Y: 90},
							Size:  robotgo.Size{W: 10, H: 10},
						},
						{
							Point: robotgo.Point{X: 1220, Y: 120},
							Size:  robotgo.Size{W: 10, H: 10},
						},
						{
							Point: robotgo.Point{X: 1250, Y: 120},
							Size:  robotgo.Size{W: 10, H: 10},
						},
						{
							Point: robotgo.Point{X: 1280, Y: 120},
							Size:  robotgo.Size{W: 10, H: 10},
						},
						{
							Point: robotgo.Point{X: 1280, Y: 150},
							Size:  robotgo.Size{W: 10, H: 10},
						},
						{
							Point: robotgo.Point{X: 1250, Y: 150},
							Size:  robotgo.Size{W: 10, H: 10},
						},
						{
							Point: robotgo.Point{X: 1250, Y: 180},
							Size:  robotgo.Size{W: 10, H: 10},
						},
						{
							Point: robotgo.Point{X: 1280, Y: 180},
							Size:  robotgo.Size{W: 10, H: 10},
						},
						{
							Point: robotgo.Point{X: 1250, Y: 180},
							Size:  robotgo.Size{W: 10, H: 10},
						},
						{
							Point: robotgo.Point{X: 1220, Y: 180},
							Size:  robotgo.Size{W: 10, H: 10},
						},
					},
				},
				{
					Window: robotgo.Rect{
						Point: robotgo.Point{X: 820, Y: 290},
						Size:  robotgo.Size{W: 180, H: 45},
					},
					DifficultyTypePoses: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 795, Y: 233},
							Size:  robotgo.Size{W: 40, H: 10},
						},
						{
							Point: robotgo.Point{X: 858, Y: 330},
							Size:  robotgo.Size{W: 40, H: 10},
						},
						{
							Point: robotgo.Point{X: 921, Y: 330},
							Size:  robotgo.Size{W: 40, H: 10},
						},
						{
							Point: robotgo.Point{X: 984, Y: 330},
							Size:  robotgo.Size{W: 40, H: 10},
						},
					},
					SmallMap: []robotgo.Rect{
						{
							Point: robotgo.Point{X: 1290, Y: 180},
							Size:  robotgo.Size{W: 10, H: 10},
						},
						{
							Point: robotgo.Point{X: 1290, Y: 150},
							Size:  robotgo.Size{W: 10, H: 10},
						},
						{
							Point: robotgo.Point{X: 1260, Y: 150},
							Size:  robotgo.Size{W: 10, H: 10},
						},
						{
							Point: robotgo.Point{X: 1260, Y: 120},
							Size:  robotgo.Size{W: 10, H: 10},
						},
						{
							Point: robotgo.Point{X: 1230, Y: 120},
							Size:  robotgo.Size{W: 10, H: 10},
						},
						{
							Point: robotgo.Point{X: 1230, Y: 150},
							Size:  robotgo.Size{W: 10, H: 10},
						},
						{
							Point: robotgo.Point{X: 1230, Y: 120},
							Size:  robotgo.Size{W: 10, H: 10},
						},
						{
							Point: robotgo.Point{X: 1230, Y: 90},
							Size:  robotgo.Size{W: 10, H: 10},
						},
						{
							Point: robotgo.Point{X: 1200, Y: 90},
							Size:  robotgo.Size{W: 10, H: 10},
						},
						{
							Point: robotgo.Point{X: 1200, Y: 60},
							Size:  robotgo.Size{W: 10, H: 10},
						},
						{
							Point: robotgo.Point{X: 1230, Y: 60},
							Size:  robotgo.Size{W: 10, H: 10},
						},
					},
				},
			},
		},
		StartBattleButton: robotgo.Rect{
			Point: robotgo.Point{X: 620, Y: 510},
			Size:  robotgo.Size{W: 110, H: 10},
		},
		EnterButton: robotgo.Rect{
			Point: robotgo.Point{X: 76, Y: 714},
			Size:  robotgo.Size{W: 100, H: 10},
		},
		EnterDButton: robotgo.Rect{
			Point: robotgo.Point{X: 585, Y: 455},
			Size:  robotgo.Size{W: 80, H: 10},
		},
	}
	return &OfficialClientChaoJiDou{chaoJiDou: c}
}
