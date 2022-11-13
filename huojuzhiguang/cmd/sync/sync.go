package main

import (
	"flag"
	"github.com/gangcheng1030/game_script/huojuzhiguang/cmd/sync/follower_server"
	"github.com/gangcheng1030/game_script/huojuzhiguang/cmd/sync/global"
	"github.com/gangcheng1030/game_script/huojuzhiguang/cmd/sync/synccore"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"net/http"
	"strconv"
	"strings"
)

var rule = flag.Int("r", 1, "leader or follower")
var followerStr = flag.String("f", "", "followers")
var mode = flag.Int("m", 1, "mode")
var port = flag.Int("p", 6688, "listen port")

func initComponent() {
	flag.Parse()
	global.Mode = *mode
	global.Port = *port

	robotgo.KeySleep = 10
	robotgo.MouseSleep = 10

	if *rule == 1 {
		global.Leader = true
	} else {
		global.Leader = false
	}

	if len(*followerStr) > 0 {
		global.Followers = strings.Split(*followerStr, ",")
	}
}

func main() {
	initComponent()

	if global.Leader {
		runLeader()
	} else {
		runFollower()
	}
}

func runLeader() {
	evChan := hook.Start()
	defer hook.End()

	for e := range evChan {
		ev := e
		if ev.Kind == hook.KeyDown || ev.Kind == hook.MouseDown || ev.Kind == hook.MouseMove {
			//脚本启动之后，只响应“鼠标移动”、"v"、"s"这几种事件
			//if !global.Open || ev.Kind == hook.MouseMove || ev.Kind == hook.KeyDown {
			go synccore.HandleFollowers(&ev, 0, 0)
			//}
		}

		if ev.Kind == hook.KeyDown {
			if ev.Rawcode == 110 || ev.Rawcode == 40 { // decimal point
				break
			}

			if ev.Rawcode == 83 { // s
				synccore.HandleS()
				continue
			}
		}
	}
}

func runFollower() {
	http.Handle("/follower", &follower_server.FollowerHandler{})
	http.ListenAndServe(":"+strconv.Itoa(*port), nil)
}
