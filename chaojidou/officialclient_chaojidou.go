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
		StartBattleButton: robotgo.Rect{
			Point: robotgo.Point{X: 620, Y: 510},
			Size:  robotgo.Size{W: 110, H: 10},
		},
		EnterDButton: robotgo.Rect{
			Point: robotgo.Point{X: 585, Y: 455},
			Size:  robotgo.Size{W: 80, H: 10},
		},
	}
	return &OfficialClientChaoJiDou{chaoJiDou: c}
}

//func (oc *OfficialClientChaoJiDou) JiuYunDong() {
//	robotgo.KeyPress(robotgo.KeyM)
//
//	var tmpPoint robotgo.Point
//	tmpPoint = utils.GetRandomPointInRect(oc.BigMap.ZhuiSu)
//	robotgo.MoveSmooth(tmpPoint.X+oc.GameWindow.X, tmpPoint.Y+oc.GameWindow.Y, 5, 10)
//	robotgo.Click()
//	robotgo.Sleep(15)
//
//	robotgo.KeyPress(robotgo.KeyF)
//	tmpPoint = utils.GetRandomPointInRect(oc.ZhuiSuMap.JiuYunDong)
//	robotgo.MoveSmooth(tmpPoint.X+oc.GameWindow.X, tmpPoint.Y+oc.GameWindow.Y, 5, 10)
//	robotgo.Click()
//
//}
