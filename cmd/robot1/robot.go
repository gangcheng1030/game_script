package main

/**
  cjdmj用的自动挂机脚本
*/

import (
	"errors"
	"flag"
	"fmt"
	"github.com/gangcheng1030/game_script/chaojidou"
	"github.com/gangcheng1030/game_script/utils/iputil"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"net/http"
	"strconv"
	"time"
)

var captainStr = flag.String("c", "official_client", "client of captain")
var meiriStr = flag.String("m", "xxcd", "meiritiaozhan")
var zhuisuStr = flag.String("z", "ddh", "zhuisu type")
var rule = flag.Int("r", 1, "leader or follower")
var leaderAddr = flag.String("l", "192.168.1.8:6688", "leader address")
var port = flag.Int("p", 6688, "listen port")

var captain chaojidou.ChaoJiDou
var localIp string

func initComponent() {
	flag.Parse()
	var err error
	captain, err = chaojidou.Build(chaojidou.ClientType(*captainStr))
	if err != nil {
		panic(err)
	}

	if *rule == chaojidou.RULE_TYPE_SLAVE {
		localIp, err = iputil.GetLocalIP()
		if err != nil {
			panic(err)
		}
		err = chaojidou.Register(*leaderAddr, localIp+":"+strconv.Itoa(*port))
		if err != nil {
			panic(err)
		}

		go func() {
			http.Handle("/follower", &chaojidou.FollowerHandler{})
			http.ListenAndServe(":"+strconv.Itoa(*port), nil)
		}()
	} else if *rule == chaojidou.RULE_TYPE_LEADER {
		go func() {
			http.Handle("/leader", &chaojidou.LeaderHandler{})
			http.ListenAndServe(":"+strconv.Itoa(*port), nil)
		}()
	} else {
		err = errors.New("invalid rule type")
		panic(err)
	}
}

func main() {
	initComponent()
	add()
}

func add() {
	robotgo.KeySleep = 100
	robotgo.MouseSleep = 100
	var endTime time.Time = time.Now()

	fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
	hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("ctrl-shift-q")
		hook.End()
	})

	fmt.Println("--- Please press shift + 1 to '2-llt + 5-zs' ---")
	hook.Register(hook.KeyDown, []string{robotgo.Shift, robotgo.Key1}, func(e hook.Event) {
		fmt.Println("shift-1")
		if endTime.Add(time.Second).After(time.Now()) {
			fmt.Println("shift-1 相隔时间太短.")
			return
		}

		chaojidou.NpcWaitSecs = 30
		chaojidou.ReadMapWaitSecs = 30
		captain.ZhuiSu(chaojidou.ZhuiSuType(*zhuisuStr), chaojidou.DIFFICULTY_TYPE_SHULIAN)
		chaojidou.NpcWaitSecs = 10
		chaojidou.ReadMapWaitSecs = 30
		for i := 0; i < 4; i++ {
			captain.ZhuiSu(chaojidou.ZhuiSuType(*zhuisuStr), chaojidou.DIFFICULTY_TYPE_SHULIAN)
		}
		chaojidou.NpcWaitSecs = 30
		chaojidou.ReadMapWaitSecs = 30
		captain.LiuLangTuan(chaojidou.LIULANGTUAN_TYPE_1, chaojidou.DIFFICULTY_TYPE_YINGXIONG)
		chaojidou.NpcWaitSecs = 10
		chaojidou.ReadMapWaitSecs = 30
		captain.LiuLangTuan(chaojidou.LIULANGTUAN_TYPE_1, chaojidou.DIFFICULTY_TYPE_YINGXIONG)
		fmt.Println("shift-1 end")
		endTime = time.Now()
	})

	fmt.Println("--- Please press shift + 2 to '2-llt + 4-zs' ---")
	hook.Register(hook.KeyDown, []string{robotgo.Shift, robotgo.Key2}, func(e hook.Event) {
		fmt.Println("shift-2")
		if endTime.Add(time.Second).After(time.Now()) {
			fmt.Println("shift-2 相隔时间太短.")
			return
		}

		chaojidou.NpcWaitSecs = 30
		chaojidou.ReadMapWaitSecs = 30
		captain.ZhuiSu(chaojidou.ZhuiSuType(*zhuisuStr), chaojidou.DIFFICULTY_TYPE_SHULIAN)
		chaojidou.NpcWaitSecs = 10
		chaojidou.ReadMapWaitSecs = 30
		for i := 0; i < 3; i++ {
			captain.ZhuiSu(chaojidou.ZhuiSuType(*zhuisuStr), chaojidou.DIFFICULTY_TYPE_SHULIAN)
		}
		chaojidou.NpcWaitSecs = 30
		chaojidou.ReadMapWaitSecs = 30
		captain.LiuLangTuan(chaojidou.LIULANGTUAN_TYPE_1, chaojidou.DIFFICULTY_TYPE_YINGXIONG)
		chaojidou.NpcWaitSecs = 10
		chaojidou.ReadMapWaitSecs = 30
		captain.LiuLangTuan(chaojidou.LIULANGTUAN_TYPE_1, chaojidou.DIFFICULTY_TYPE_YINGXIONG)
		fmt.Println("shift-2 end")
		endTime = time.Now()
	})

	fmt.Println("--- Please press shift + s to '1-zs + 2-llt + 4-jb' ---")
	hook.Register(hook.KeyDown, []string{robotgo.Shift, robotgo.KeyS}, func(e hook.Event) {
		fmt.Println("shift-s")
		if endTime.Add(time.Second).After(time.Now()) {
			fmt.Println("shift-s 相隔时间太短.")
			return
		}
		chaojidou.NpcWaitSecs = 30
		chaojidou.ReadMapWaitSecs = 30
		captain.ZhuiSu(chaojidou.ZhuiSuType(*zhuisuStr), chaojidou.DIFFICULTY_TYPE_SHULIAN)
		captain.LiuLangTuan(chaojidou.LIULANGTUAN_TYPE_1, chaojidou.DIFFICULTY_TYPE_YINGXIONG)
		chaojidou.NpcWaitSecs = 10
		chaojidou.ReadMapWaitSecs = 30
		captain.LiuLangTuan(chaojidou.LIULANGTUAN_TYPE_1, chaojidou.DIFFICULTY_TYPE_YINGXIONG)
		chaojidou.NpcWaitSecs = 50
		chaojidou.ReadMapWaitSecs = 30
		captain.JinBen(chaojidou.JINBEN_TYPE_SUXING, chaojidou.DIFFICULTY_TYPE_MAOXIAN, 1)
		chaojidou.NpcWaitSecs = 10
		captain.JinBen(chaojidou.JINBEN_TYPE_SUXING, chaojidou.DIFFICULTY_TYPE_MAOXIAN, 1)
		captain.JinBen(chaojidou.JINBEN_TYPE_HEIAN, chaojidou.DIFFICULTY_TYPE_MAOXIAN, 1)
		captain.JinBen(chaojidou.JINBEN_TYPE_HEIAN, chaojidou.DIFFICULTY_TYPE_MAOXIAN, 1)
		chaojidou.NpcWaitSecs = 30
		fmt.Println("shift-s end")
		endTime = time.Now()
	})

	fmt.Println("--- Please press shift + a to '2-llt + 4-jb' ---")
	hook.Register(hook.KeyDown, []string{robotgo.Shift, robotgo.KeyA}, func(e hook.Event) {
		fmt.Println("shift-a")
		if endTime.Add(time.Second).After(time.Now()) {
			fmt.Println("shift-a 相隔时间太短.")
			return
		}
		chaojidou.NpcWaitSecs = 30
		chaojidou.ReadMapWaitSecs = 30
		captain.LiuLangTuan(chaojidou.LIULANGTUAN_TYPE_1, chaojidou.DIFFICULTY_TYPE_YINGXIONG)
		chaojidou.NpcWaitSecs = 10
		chaojidou.ReadMapWaitSecs = 30
		captain.LiuLangTuan(chaojidou.LIULANGTUAN_TYPE_1, chaojidou.DIFFICULTY_TYPE_YINGXIONG)
		chaojidou.NpcWaitSecs = 50
		chaojidou.ReadMapWaitSecs = 30
		captain.JinBen(chaojidou.JINBEN_TYPE_SUXING, chaojidou.DIFFICULTY_TYPE_MAOXIAN, 1)
		chaojidou.NpcWaitSecs = 10
		captain.JinBen(chaojidou.JINBEN_TYPE_SUXING, chaojidou.DIFFICULTY_TYPE_MAOXIAN, 1)
		captain.JinBen(chaojidou.JINBEN_TYPE_HEIAN, chaojidou.DIFFICULTY_TYPE_MAOXIAN, 1)
		captain.JinBen(chaojidou.JINBEN_TYPE_HEIAN, chaojidou.DIFFICULTY_TYPE_MAOXIAN, 1)
		chaojidou.NpcWaitSecs = 30
		fmt.Println("shift-a end")
		endTime = time.Now()
	})

	fmt.Println("--- Please press shift + t to llt ---")
	hook.Register(hook.KeyDown, []string{robotgo.Shift, robotgo.KeyT}, func(e hook.Event) {
		fmt.Println("shift-t")
		if endTime.Add(time.Second).After(time.Now()) {
			fmt.Println("shift-t 相隔时间太短.")
			return
		}
		chaojidou.NpcWaitSecs = 30
		chaojidou.ReadMapWaitSecs = 30
		captain.LiuLangTuan(chaojidou.LIULANGTUAN_TYPE_1, chaojidou.DIFFICULTY_TYPE_YINGXIONG)
		fmt.Println("shift-t end")
		endTime = time.Now()
	})

	fmt.Println("--- Please press shift + e to haqszc ---")
	hook.Register(hook.KeyDown, []string{robotgo.Shift, robotgo.KeyE}, func(e hook.Event) {
		fmt.Println("shift-e")
		if endTime.Add(time.Second).After(time.Now()) {
			fmt.Println("shift-e 相隔时间太短.")
			return
		}
		chaojidou.NpcWaitSecs = 30
		chaojidou.ReadMapWaitSecs = 30
		captain.JinBen(chaojidou.JINBEN_TYPE_HEIAN, chaojidou.DIFFICULTY_TYPE_MAOXIAN, 1)
		fmt.Println("shift-e end")
		endTime = time.Now()
	})

	fmt.Println("--- Please press shift + d to sxdcs ---")
	hook.Register(hook.KeyDown, []string{robotgo.Shift, robotgo.KeyD}, func(e hook.Event) {
		fmt.Println("shift-d")
		if endTime.Add(time.Second).After(time.Now()) {
			fmt.Println("shift-d 相隔时间太短.")
			return
		}
		chaojidou.NpcWaitSecs = 30
		chaojidou.ReadMapWaitSecs = 30
		captain.JinBen(chaojidou.JINBEN_TYPE_SUXING, chaojidou.DIFFICULTY_TYPE_MAOXIAN, 1)
		fmt.Println("shift-d end")
		endTime = time.Now()
	})

	fmt.Println("--- Please press shift + r to mrtz ---")
	hook.Register(hook.KeyDown, []string{robotgo.Shift, robotgo.KeyR}, func(e hook.Event) {
		fmt.Println("shift-r")
		if endTime.Add(time.Second).After(time.Now()) {
			fmt.Println("shift-r 相隔时间太短.")
			return
		}

		var tp chaojidou.MeiRiType = chaojidou.MeiRiType(*meiriStr)
		chaojidou.NpcWaitSecs = 30
		chaojidou.ReadMapWaitSecs = 30
		captain.MeiRiTiaoZhan(tp, chaojidou.DIFFICULTY_TYPE_MAOXIAN)
		fmt.Println("shift-r end")
		endTime = time.Now()
	})

	fmt.Println("--- Please press shift + q to zs ---")
	hook.Register(hook.KeyDown, []string{robotgo.Shift, robotgo.KeyQ}, func(e hook.Event) {
		fmt.Println("shift-q")
		if endTime.Add(time.Second).After(time.Now()) {
			fmt.Println("shift-q 相隔时间太短.")
			return
		}

		chaojidou.NpcWaitSecs = 30
		chaojidou.ReadMapWaitSecs = 30
		captain.ZhuiSu(chaojidou.ZhuiSuType(*zhuisuStr), chaojidou.DIFFICULTY_TYPE_SHULIAN)
		fmt.Println("shift-q end")
		endTime = time.Now()
	})

	fmt.Println("--- Please press shift + w to ads ---")
	hook.Register(hook.KeyDown, []string{robotgo.Shift, robotgo.KeyW}, func(e hook.Event) {
		fmt.Println("shift-w")
		if endTime.Add(time.Second).After(time.Now()) {
			fmt.Println("shift-w 相隔时间太短.")
			return
		}

		chaojidou.NpcWaitSecs = 30
		chaojidou.ReadMapWaitSecs = 30
		captain.AoDeSai(chaojidou.DIFFICULTY_TYPE_MAOXIAN)
		fmt.Println("shift-w end")
		endTime = time.Now()
	})

	fmt.Println("--- Please press shift + 3 to jzys ---")
	hook.Register(hook.KeyDown, []string{robotgo.Shift, robotgo.Key3}, func(e hook.Event) {
		fmt.Println("shift-3")
		if endTime.Add(time.Second).After(time.Now()) {
			fmt.Println("shift-3 相隔时间太短.")
			return
		}

		chaojidou.NpcWaitSecs = 30
		chaojidou.ReadMapWaitSecs = 30
		captain.JiZhanYanSuan(0)
		fmt.Println("shift-3 end")
		endTime = time.Now()
	})

	s := hook.Start()
	<-hook.Process(s)
}
