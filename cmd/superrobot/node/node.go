package main

import (
	"flag"
	"fmt"
	"github.com/gangcheng1030/game_script/chaojidou"
	autorobot_client "github.com/gangcheng1030/game_script/cmd/autorobot/client"
	"github.com/gangcheng1030/game_script/cmd/autorobot/server"
	"github.com/gangcheng1030/game_script/cmd/superrobot/client"
	"github.com/gangcheng1030/game_script/cmd/superrobot/model"
	"github.com/gangcheng1030/game_script/cmd/superrobot/node/global"
	"github.com/gangcheng1030/game_script/utils/robotgoutil"
	"github.com/go-vgo/robotgo"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var controllerAddr = flag.String("addr", "127.0.0.1:8080", "controller address")
var id = flag.String("id", "node2", "node id")
var port = flag.Int("p", 6688, "listen port")

var cfg *model.NodeConfig
var startGameButton robotgo.Rect
var autoRobotFollowerHandler *server.FollowerHandler
var captain chaojidou.ChaoJiDou
var event *model.Event

func initComponent() {
	flag.Parse()
	var err error

	robotgo.KeySleep = 100
	robotgo.MouseSleep = 100

	cfg, err = client.GetNodeConfig(*controllerAddr, *id)
	if err != nil {
		panic(err)
	}

	err = os.Chdir(cfg.DirName)
	if err != nil {
		panic(err)
	}

	startGameButton = robotgo.Rect{
		Point: robotgo.Point{X: 1250, Y: 695},
		Size:  robotgo.Size{W: 40, H: 10},
	}

	autoRobotFollowerHandler = &server.FollowerHandler{CjdDir: cfg.DirName, StartGameButton: startGameButton, FullScreenMode: cfg.FullScreenMode}
	go func() {
		http.Handle("/follower/autorobot", autoRobotFollowerHandler)
		http.Handle("/follower", &chaojidou.FollowerHandler{})
		http.ListenAndServe(":"+strconv.Itoa(*port), nil)
	}()

	err = client.OnlineNode(*controllerAddr, *id, strconv.Itoa(*port), cfg)
	if err != nil {
		panic(err)
	}

	//创建监听退出chan
	c := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				log.Println("Program Exit...", s)
				if event != nil {
					err = client.FinishEvent(*controllerAddr, event)
					if err != nil {
						log.Printf("finishEvent %d err: %v", event.Id, err)
					}
				}
				err = client.OfflineNode(*controllerAddr, *id, cfg)
				if err != nil {
					log.Printf("offlineNode err: %v", err)
				}
				os.Exit(0)
			default:
				log.Println("other signal", s)
			}
		}
	}()
}

func main() {
	initComponent()
	quitErrMsg := fmt.Sprintf("node %s 不存在", *id)
	follower := false

	for {
		if autoRobotFollowerHandler.IsRunning {
			follower = true
			time.Sleep(60 * time.Second)
			continue
		}

		if follower {
			follower = false
			time.Sleep(60 * time.Second)
			continue
		}

		var err error
		event, err = client.PickEvent(*controllerAddr, *id, cfg)
		if err != nil {
			log.Printf("pickEvent err: %v", err)
			if err.Error() == quitErrMsg && !autoRobotFollowerHandler.IsRunning {
				break
			}
			time.Sleep(30 * time.Second)
			continue
		}

		if event != nil {
			for {
				res := handleEvent(event)
				if res {
					break
				}
				captain.ForceQuit(true)
			}
			for i := 0; i < 5; i++ {
				err = client.FinishEvent(*controllerAddr, event)
				if err != nil {
					log.Printf("finishEvent %d err: %v", event.Id, err)
					time.Sleep(500 * time.Millisecond)
					continue
				}
				break
			}
			event = nil
		}

		time.Sleep(30 * time.Second)
	}
}

func handleEvent(event *model.Event) bool {
	account := &event.Account
	chaojidou.Follwers = account.FollowerAddrs

	if len(event.FollowerNodeIds) > 0 {
		for j := range account.FollowerAddrs {
			accountTmp := server.Account{
				AccountName: account.FollowerAccountNames[j],
				Password:    account.FollowerPasswords[j],
			}
			for {
				err := autorobot_client.SendEvent(account.FollowerAddrs[j], "signin", accountTmp)
				if err == nil {
					break
				}
				robotgo.Sleep(3)
			}
		}
		robotgo.Sleep(6)
	}

	cjd := exec.Command(cfg.DirName + "\\Launcher.exe")
	err := cjd.Run()
	if err != nil {
		panic(err)
	}
	robotgo.Sleep(45)

	// 开始游戏
	robotgoutil.ClickButton(startGameButton, 92)

	captain, err = chaojidou.Build(chaojidou.CLIENT_TYPE_OFFICIAL)
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

		isOnline := handleOneRole(&account.Roles[j], event.MeiRi, first, last)
		first = false
		if !isOnline {
			account.Roles = account.Roles[j:]
			return false
		}
	}

	return true
}

func handleOneRole(role *model.Role, meiri string, first bool, last bool) bool {
	if len(role.FollowerRoleIds) > 0 {
		for i := range role.FollowerRoleIds {
			roleTmp := server.Role{
				Id:    role.FollowerRoleIds[i],
				First: first,
				Last:  last,

				DisableCardsUp:     role.DisableCardsUp,
				DisablePreClearBag: role.DisablePreClearBag,
				PostClearBag:       role.PostClearBag,
			}
			for {
				err := autorobot_client.SendEvent(chaojidou.Follwers[i], "select_role", roleTmp)
				if err == nil {
					break
				}
				robotgo.Sleep(3)
			}
		}
		robotgo.Sleep(6)
	}

	captain.SelectRole(role.Id-1, first, cfg.FullScreenMode)

	captain.RepairEquipment()
	chaojidou.NpcWaitSecs = 20
	if !role.DisablePreClearBag {
		captain.ClearBag(true)
	}
	chaojidou.NpcWaitSecs = 30
	if !role.DisableCardsUp {
		captain.CardsUp()
	}

	if len(role.FollowerRoleIds) > 0 {
		captain.CreateGroup(role.FollowerJunTuanNames, first)
	}

	if len(role.Fubens) == 0 {
		chaojidou.NpcWaitSecs = 20
		captain.LiuLangTuan(chaojidou.LIULANGTUAN_TYPE_1, chaojidou.DIFFICULTY_TYPE_YINGXIONG)
		if !captain.IsOnline(true) {
			role.Fubens = global.DefaultFubens[1:]
			role.DisableCardsUp = true
			role.DisablePreClearBag = true
			return false
		}

		chaojidou.NpcWaitSecs = 10
		captain.LiuLangTuan(chaojidou.LIULANGTUAN_TYPE_1, chaojidou.DIFFICULTY_TYPE_YINGXIONG)
		if !captain.IsOnline(true) {
			role.Fubens = global.DefaultFubens[2:]
			role.DisableCardsUp = true
			role.DisablePreClearBag = true
			return false
		}

		chaojidou.NpcWaitSecs = 45
		captain.JinBen(chaojidou.JINBEN_TYPE_SUXING, chaojidou.DIFFICULTY_TYPE_MAOXIAN, 1)
		if !captain.IsOnline(true) {
			role.Fubens = global.DefaultFubens[3:]
			role.DisableCardsUp = true
			role.DisablePreClearBag = true
			return false
		}

		chaojidou.NpcWaitSecs = 10
		captain.JinBen(chaojidou.JINBEN_TYPE_SUXING, chaojidou.DIFFICULTY_TYPE_MAOXIAN, 1)
		if !captain.IsOnline(true) {
			role.Fubens = global.DefaultFubens[4:]
			role.DisableCardsUp = true
			role.DisablePreClearBag = true
			return false
		}

		captain.JinBen(chaojidou.JINBEN_TYPE_HEIAN, chaojidou.DIFFICULTY_TYPE_MAOXIAN, 1)
		if !captain.IsOnline(true) {
			role.Fubens = global.DefaultFubens[5:]
			role.DisableCardsUp = true
			role.DisablePreClearBag = true
			return false
		}

		captain.JinBen(chaojidou.JINBEN_TYPE_HEIAN, chaojidou.DIFFICULTY_TYPE_MAOXIAN, 1)
		if !captain.IsOnline(true) {
			role.Fubens = global.EmptyFubens
			role.DisableCardsUp = true
			role.DisablePreClearBag = true
			return false
		}
		chaojidou.NpcWaitSecs = 30
	} else {
		chaojidou.NpcWaitSecs = 30
		for i, fuben := range role.Fubens {
			if fuben.Name == "meiri" {
				captain.MeiRiTiaoZhan(chaojidou.MeiRiType(meiri), chaojidou.DifficultyType(fuben.Difficulty))
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
				if first {
					chaojidou.NpcWaitSecs = 40
				}
				captain.JinBen(chaojidou.JINBEN_TYPE_SUXING, chaojidou.DIFFICULTY_TYPE_MAOXIAN, 1)
				chaojidou.NpcWaitSecs = 30
			} else if fuben.Name == "haqszh" {
				if first {
					chaojidou.NpcWaitSecs = 40
				}
				captain.JinBen(chaojidou.JINBEN_TYPE_HEIAN, chaojidou.DIFFICULTY_TYPE_MAOXIAN, 1)
				chaojidou.NpcWaitSecs = 30
			} else if fuben.Name == "jzys" {
				captain.JiZhanYanSuan()
			} else if fuben.Name == "xyb" {
				captain.XingYunBi(fuben.Difficulty)
			}

			if !captain.IsOnline(true) {
				if i < len(role.Fubens)-1 {
					role.Fubens = role.Fubens[(i + 1):]
				} else {
					role.Fubens = global.EmptyFubens
				}
				role.DisableCardsUp = true
				role.DisablePreClearBag = true
				return false
			}
		}
	}

	if len(role.FollowerRoleIds) > 0 {
		for i := range role.FollowerRoleIds {
			roleTmp := server.Role{
				Id:    role.FollowerRoleIds[i],
				First: first,
				Last:  last,

				DisablePreClearBag: role.DisablePreClearBag,
				PostClearBag:       role.PostClearBag,
			}
			for {
				err := autorobot_client.SendEvent(chaojidou.Follwers[i], "quit", roleTmp)
				if err == nil {
					break
				}
				robotgo.Sleep(3)
			}
		}
		robotgo.Sleep(6)
	}

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

	return true
}
