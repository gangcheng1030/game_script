package main

/**
  cjdmj用的自动挂机脚本
*/

import (
	"flag"
	"github.com/gangcheng1030/game_script/chaojidou"
	"github.com/gangcheng1030/game_script/cmd/autorobot/config"
	"github.com/gangcheng1030/game_script/utils/robotgoutil"
	"github.com/go-vgo/robotgo"
	"github.com/shirou/gopsutil/v3/process"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var captainStr = flag.String("c", "official_client", "client of captain")
var cjdDir = flag.String("d", "G:\\Netease\\CJDMJ", "cjd dir")
var configPath = flag.String("conf", "autorobot.json", "config")

var cfg *config.AutoRobotConfig
var captain chaojidou.ChaoJiDou
var robotDir string
var startGameButton robotgo.Rect

func initComponent() {
	flag.Parse()
	var err error

	robotDir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	cfg, err = config.InitConfig(robotDir + "\\" + *configPath)
	if err != nil {
		panic(err)
	}

	err = os.Chdir(*cjdDir)
	if err != nil {
		panic(err)
	}

	startGameButton = robotgo.Rect{
		Point: robotgo.Point{X: 1250, Y: 695},
		Size:  robotgo.Size{W: 40, H: 10},
	}
}

func main() {
	initComponent()
	start()
}

func start() {
	robotgo.KeySleep = 100
	robotgo.MouseSleep = 100

	for i := 0; i < len(cfg.Accounts); i++ {
		cjd := exec.Command(*cjdDir + "\\Launcher.exe")
		err := cjd.Run()
		if err != nil {
			panic(err)
		}
		robotgo.Sleep(45)

		account := cfg.Accounts[i]

		// 开始游戏
		robotgoutil.ClickButton(startGameButton, 110)

		captain, err = chaojidou.Build(chaojidou.ClientType(*captainStr))
		if err != nil {
			panic(err)
		}

		// 将窗口置于前台
		robotgo.MoveSmooth(946, 190, 0.9, 0.9)
		robotgo.Click()
		robotgo.Sleep(3)

		// 登录
		captain.SignIn(account.AccountName, account.Password)

		// 角色
		first := true
		last := false
		for j := 0; j < len(account.Roles); j++ {
			if j == len(account.Roles)-1 {
				last = true
			}

			handleOneRole(account.Roles[j], first, last)
			first = false
		}

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

}

func handleOneRole(role config.Role, first bool, last bool) {
	captain.SelectRole(role.Id-1, first)

	captain.RepairEquipment()
	captain.ClearBag()
	captain.CardsUp()

	if len(role.Fubens) == 0 {
		chaojidou.NpcWaitSecs = 20
		captain.MeiRiTiaoZhan(chaojidou.MeiRiType(cfg.MeiRi), chaojidou.DIFFICULTY_TYPE_MAOXIAN)
		captain.ClearBag()
		captain.LiuLangTuan(chaojidou.LIULANGTUAN_TYPE_1, chaojidou.DIFFICULTY_TYPE_YINGXIONG)
		chaojidou.NpcWaitSecs = 10
		captain.LiuLangTuan(chaojidou.LIULANGTUAN_TYPE_1, chaojidou.DIFFICULTY_TYPE_YINGXIONG)
		chaojidou.NpcWaitSecs = 45
		captain.JinBen(chaojidou.JINBEN_TYPE_SUXING, chaojidou.DIFFICULTY_TYPE_MAOXIAN, 1)
		chaojidou.NpcWaitSecs = 10
		captain.JinBen(chaojidou.JINBEN_TYPE_SUXING, chaojidou.DIFFICULTY_TYPE_MAOXIAN, 1)
		captain.JinBen(chaojidou.JINBEN_TYPE_HEIAN, chaojidou.DIFFICULTY_TYPE_MAOXIAN, 1)
		captain.JinBen(chaojidou.JINBEN_TYPE_HEIAN, chaojidou.DIFFICULTY_TYPE_MAOXIAN, 1)
		chaojidou.NpcWaitSecs = 30
	} else {
		chaojidou.NpcWaitSecs = 30
		for i, fuben := range role.Fubens {
			if fuben.Name == "meiri" {
				captain.MeiRiTiaoZhan(chaojidou.MeiRiType(cfg.MeiRi), chaojidou.DIFFICULTY_TYPE_MAOXIAN)
				if i < len(role.Fubens)-1 {
					captain.ClearBag()
				}
			} else if fuben.Name == chaojidou.ZHUISU_TYPE_JIUYUNDONG {
				captain.ZhuiSu(chaojidou.ZHUISU_TYPE_JIUYUNDONG, chaojidou.DIFFICULTY_TYPE_SHULIAN)
			} else if fuben.Name == chaojidou.ZHUISU_TYPE_DADUHUI {
				captain.ZhuiSu(chaojidou.ZHUISU_TYPE_DADUHUI, chaojidou.DIFFICULTY_TYPE_SHULIAN)
			} else if fuben.Name == "llt1" {
				captain.LiuLangTuan(chaojidou.LIULANGTUAN_TYPE_1, chaojidou.DIFFICULTY_TYPE_YINGXIONG)
			} else if fuben.Name == "sxdcs" {
				captain.JinBen(chaojidou.JINBEN_TYPE_SUXING, chaojidou.DIFFICULTY_TYPE_MAOXIAN, 1)
			} else if fuben.Name == "haqszh" {
				captain.JinBen(chaojidou.JINBEN_TYPE_HEIAN, chaojidou.DIFFICULTY_TYPE_MAOXIAN, 1)
			}
		}
	}

	captain.RepairEquipment()
	captain.ClearBag()
	captain.CardsDown()

	if !last {
		captain.QuitRole()
	} else {
		captain.Quit()
	}
}
