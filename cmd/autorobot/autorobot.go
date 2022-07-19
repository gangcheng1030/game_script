package main

/**
  cjdmj用的自动挂机脚本
*/

import (
	"flag"
	"github.com/gangcheng1030/game_script/chaojidou"
	"github.com/gangcheng1030/game_script/cmd/autorobot/client"
	"github.com/gangcheng1030/game_script/cmd/autorobot/config"
	"github.com/gangcheng1030/game_script/cmd/autorobot/server"
	"github.com/gangcheng1030/game_script/utils/robotgoutil"
	"github.com/go-vgo/robotgo"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

var captainStr = flag.String("c", "official_client", "client of captain")
var mode = flag.Int("m", 1, "模式：1 表示 单人模式；2 表示 组队模式")
var rule = flag.Int("r", 1, "leader or follower")
var cjdDir = flag.String("d", "G:\\Netease\\CJDMJ", "cjd dir")
var configPath = flag.String("conf", "autorobot.json", "config")
var port = flag.Int("p", 6688, "listen port")
var fullScreenMode = flag.Int("f", 0, "全屏模式： 0 - 1176 * 664; 1 - 1920 * 1080")
var begin = flag.String("s", "", "开始时间，格式: 2006-01-02 15:04:05")

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

	if *mode == 1 || *rule == chaojidou.RULE_TYPE_LEADER {
		cfg, err = config.InitConfig(robotDir + "\\" + *configPath)
		if err != nil {
			panic(err)
		}
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
	if len(*begin) > 0 {
		beginTime, err := time.ParseInLocation("2006-01-02 15:04:05", *begin, time.Local)
		if err != nil {
			panic(err)
		}
		for {
			if beginTime.Before(time.Now()) {
				break
			}
			time.Sleep(time.Minute)
		}
	}
	if *mode != 1 && *rule == chaojidou.RULE_TYPE_SLAVE {
		robotgo.KeySleep = 100
		robotgo.MouseSleep = 100

		http.Handle("/follower/autorobot", &server.FollowerHandler{CjdDir: *cjdDir, StartGameButton: startGameButton, FullScreenMode: *fullScreenMode})
		http.Handle("/follower", &chaojidou.FollowerHandler{})
		http.ListenAndServe(":"+strconv.Itoa(*port), nil)
	} else {
		start()
	}
}

func start() {
	robotgo.KeySleep = 100
	robotgo.MouseSleep = 100

	for i := 0; i < len(cfg.Accounts); i++ {
		account := cfg.Accounts[i]

		if *mode != 1 {
			chaojidou.Follwers = account.FollowerAddrs
			for j := range account.FollowerAddrs {
				accountTmp := server.Account{
					AccountName: account.FollowerAccountNames[j],
					Password:    account.FollowerPasswords[j],
				}
				for {
					err := client.SendEvent(account.FollowerAddrs[j], "signin", accountTmp)
					if err == nil {
						break
					}
					robotgo.Sleep(3)
				}
			}
			robotgo.Sleep(6)
		}

		cjd := exec.Command(*cjdDir + "\\Launcher.exe")
		err := cjd.Run()
		if err != nil {
			panic(err)
		}
		robotgo.Sleep(45)

		// 开始游戏
		robotgoutil.ClickButton(startGameButton, 92)

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
	}

}

func handleOneRole(role config.Role, first bool, last bool) {
	if *mode != 1 {
		for i := range role.FollowerRoleIds {
			roleTmp := server.Role{
				Id:    role.FollowerRoleIds[i],
				First: first,
				Last:  last,

				PostClearBag: role.PostClearBag,
			}
			for {
				err := client.SendEvent(chaojidou.Follwers[i], "select_role", roleTmp)
				if err == nil {
					break
				}
				robotgo.Sleep(3)
			}
		}
		robotgo.Sleep(6)
	}

	captain.SelectRole(role.Id-1, first, *fullScreenMode)

	captain.RepairEquipment()
	chaojidou.NpcWaitSecs = 20
	if !role.DisablePreClearBag {
		captain.ClearBag(true)
	}
	chaojidou.NpcWaitSecs = 30
	captain.CardsUp()

	if *mode != 1 {
		captain.CreateGroup(role.FollowerJunTuanNames, first)
	}

	if len(role.Fubens) == 0 {
		chaojidou.NpcWaitSecs = 20
		if *mode == 1 {
			captain.MeiRiTiaoZhan(chaojidou.MeiRiType(cfg.MeiRi), chaojidou.DIFFICULTY_TYPE_MAOXIAN)
			captain.RepairEquipment()
			captain.ClearBag(true)
		}
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
				captain.MeiRiTiaoZhan(chaojidou.MeiRiType(cfg.MeiRi), chaojidou.DifficultyType(fuben.Difficulty))
				if i < len(role.Fubens)-1 {
					captain.ClearBag(true)
				}
			} else if fuben.Name == chaojidou.ZHUISU_TYPE_GUTU {
				captain.ZhuiSu(chaojidou.ZHUISU_TYPE_GUTU, chaojidou.DifficultyType(fuben.Difficulty))
			} else if fuben.Name == chaojidou.ZHUISU_TYPE_JIUYUNDONG {
				captain.ZhuiSu(chaojidou.ZHUISU_TYPE_JIUYUNDONG, chaojidou.DifficultyType(fuben.Difficulty))
			} else if fuben.Name == chaojidou.ZHUISU_TYPE_DADUHUI {
				captain.ZhuiSu(chaojidou.ZHUISU_TYPE_DADUHUI, chaojidou.DifficultyType(fuben.Difficulty))
			} else if fuben.Name == chaojidou.ZHUISU_TYPE_TONGHUAZHEN {
				captain.ZhuiSu(chaojidou.ZHUISU_TYPE_TONGHUAZHEN, chaojidou.DifficultyType(fuben.Difficulty))
			} else if fuben.Name == chaojidou.ZHUISU_TYPE_KAERJIAYIZHI {
				captain.ZhuiSu(chaojidou.ZHUISU_TYPE_KAERJIAYIZHI, chaojidou.DifficultyType(fuben.Difficulty))
			} else if fuben.Name == chaojidou.ZHUISU_TYPE_GELAXIYA {
				captain.ZhuiSu(chaojidou.ZHUISU_TYPE_GELAXIYA, chaojidou.DifficultyType(fuben.Difficulty))
			} else if fuben.Name == chaojidou.ZHUISU_TYPE_BULINDIXI {
				captain.ZhuiSu(chaojidou.ZHUISU_TYPE_BULINDIXI, chaojidou.DifficultyType(fuben.Difficulty))
			} else if fuben.Name == chaojidou.ZHUISU_TYPE_LALAIYE {
				captain.ZhuiSu(chaojidou.ZHUISU_TYPE_LALAIYE, chaojidou.DifficultyType(fuben.Difficulty))
			} else if fuben.Name == "llt1" {
				captain.LiuLangTuan(chaojidou.LIULANGTUAN_TYPE_1, chaojidou.DIFFICULTY_TYPE_YINGXIONG)
			} else if fuben.Name == "sxdcs" {
				captain.JinBen(chaojidou.JINBEN_TYPE_SUXING, chaojidou.DIFFICULTY_TYPE_MAOXIAN, 1)
			} else if fuben.Name == "haqszh" {
				captain.JinBen(chaojidou.JINBEN_TYPE_HEIAN, chaojidou.DIFFICULTY_TYPE_MAOXIAN, 1)
			} else if fuben.Name == "jzys" {
				captain.JiZhanYanSuan()
			}
		}
	}

	if *mode != 1 {
		for i := range role.FollowerRoleIds {
			roleTmp := server.Role{
				Id:    role.FollowerRoleIds[i],
				First: first,
				Last:  last,

				DisablePreClearBag: role.DisablePreClearBag,
				PostClearBag:       role.PostClearBag,
			}
			for {
				err := client.SendEvent(chaojidou.Follwers[i], "quit", roleTmp)
				if err == nil {
					break
				}
				robotgo.Sleep(3)
			}
		}
		robotgo.Sleep(6)
	}

	captain.RepairEquipment()
	if role.PostClearBag {
		chaojidou.NpcWaitSecs = 8
		captain.ClearBag(false)
		chaojidou.NpcWaitSecs = 30
	}
	captain.CardsDown()

	if !last {
		captain.QuitRole()
	} else {
		captain.Quit()
	}
}
