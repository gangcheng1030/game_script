package main

/**
  cjdmj用的自动挂机脚本
*/

import (
	"flag"
	"fmt"
	"github.com/gangcheng1030/game_script/chaojidou"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"time"
)

var captainStr = flag.String("c", "official_client", "client of captain")
var meiriStr = flag.String("m", "xxcd", "meiritiaozhan")

var captain chaojidou.ChaoJiDou

func initComponent() {
	flag.Parse()
	var err error
	captain, err = chaojidou.Build(chaojidou.ClientType(*captainStr))
	if err != nil {
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

	fmt.Println("--- Please press shift + s to confirm ---")
	hook.Register(hook.KeyDown, []string{robotgo.Shift, robotgo.KeyS}, func(e hook.Event) {
		fmt.Println("shift-s")
		if endTime.Add(time.Second).After(time.Now()) {
			fmt.Println("shift-s 相隔时间太短.")
			return
		}
		chaojidou.NpcWaitSecs = 30
		chaojidou.ReadMapWaitSecs = 120
		captain.JiuYunDong(chaojidou.DIFFICULTY_TYPE_SHULIAN, 1)
		captain.LiuLangTuan(chaojidou.LIULANGTUAN_TYPE_1, chaojidou.DIFFICULTY_TYPE_YINGXIONG)
		chaojidou.NpcWaitSecs = 10
		chaojidou.ReadMapWaitSecs = 30
		captain.LiuLangTuan(chaojidou.LIULANGTUAN_TYPE_1, chaojidou.DIFFICULTY_TYPE_YINGXIONG)
		chaojidou.NpcWaitSecs = 50
		chaojidou.ReadMapWaitSecs = 60
		captain.JinBen(chaojidou.JINBEN_TYPE_SUXING, chaojidou.DIFFICULTY_TYPE_MAOXIAN, 2)
		chaojidou.NpcWaitSecs = 10
		captain.JinBen(chaojidou.JINBEN_TYPE_HEIAN, chaojidou.DIFFICULTY_TYPE_MAOXIAN, 1)
		captain.JinBen(chaojidou.JINBEN_TYPE_HEIAN, chaojidou.DIFFICULTY_TYPE_MAOXIAN, 1)
		chaojidou.NpcWaitSecs = 30
		fmt.Println("shift-s end")
		endTime = time.Now()
	})

	fmt.Println("--- Please press shift + a to confirm ---")
	hook.Register(hook.KeyDown, []string{robotgo.Shift, robotgo.KeyA}, func(e hook.Event) {
		fmt.Println("shift-a")
		if endTime.Add(time.Second).After(time.Now()) {
			fmt.Println("shift-a 相隔时间太短.")
			return
		}
		chaojidou.NpcWaitSecs = 30
		chaojidou.ReadMapWaitSecs = 120
		captain.LiuLangTuan(chaojidou.LIULANGTUAN_TYPE_1, chaojidou.DIFFICULTY_TYPE_YINGXIONG)
		chaojidou.NpcWaitSecs = 10
		chaojidou.ReadMapWaitSecs = 30
		captain.LiuLangTuan(chaojidou.LIULANGTUAN_TYPE_1, chaojidou.DIFFICULTY_TYPE_YINGXIONG)
		chaojidou.NpcWaitSecs = 50
		chaojidou.ReadMapWaitSecs = 60
		captain.JinBen(chaojidou.JINBEN_TYPE_SUXING, chaojidou.DIFFICULTY_TYPE_MAOXIAN, 2)
		chaojidou.NpcWaitSecs = 10
		captain.JinBen(chaojidou.JINBEN_TYPE_HEIAN, chaojidou.DIFFICULTY_TYPE_MAOXIAN, 1)
		captain.JinBen(chaojidou.JINBEN_TYPE_HEIAN, chaojidou.DIFFICULTY_TYPE_MAOXIAN, 1)
		chaojidou.NpcWaitSecs = 30
		fmt.Println("shift-a end")
		endTime = time.Now()
	})

	fmt.Println("--- Please press shift + t to 流浪团 ---")
	hook.Register(hook.KeyDown, []string{robotgo.Shift, robotgo.KeyT}, func(e hook.Event) {
		fmt.Println("shift-t")
		if endTime.Add(time.Second).After(time.Now()) {
			fmt.Println("shift-t 相隔时间太短.")
			return
		}
		chaojidou.NpcWaitSecs = 30
		chaojidou.ReadMapWaitSecs = 60
		captain.LiuLangTuan(chaojidou.LIULANGTUAN_TYPE_1, chaojidou.DIFFICULTY_TYPE_YINGXIONG)
		fmt.Println("shift-t end")
		endTime = time.Now()
	})

	fmt.Println("--- Please press shift + e to 黑暗侵蚀之巢 ---")
	hook.Register(hook.KeyDown, []string{robotgo.Shift, robotgo.KeyE}, func(e hook.Event) {
		fmt.Println("shift-e")
		if endTime.Add(time.Second).After(time.Now()) {
			fmt.Println("shift-e 相隔时间太短.")
			return
		}
		chaojidou.NpcWaitSecs = 30
		chaojidou.ReadMapWaitSecs = 60
		captain.JinBen(chaojidou.JINBEN_TYPE_HEIAN, chaojidou.DIFFICULTY_TYPE_MAOXIAN, 1)
		fmt.Println("shift-e end")
		endTime = time.Now()
	})

	fmt.Println("--- Please press shift + d to 苏醒的传说 ---")
	hook.Register(hook.KeyDown, []string{robotgo.Shift, robotgo.KeyD}, func(e hook.Event) {
		fmt.Println("shift-d")
		if endTime.Add(time.Second).After(time.Now()) {
			fmt.Println("shift-d 相隔时间太短.")
			return
		}
		chaojidou.NpcWaitSecs = 30
		chaojidou.ReadMapWaitSecs = 60
		captain.JinBen(chaojidou.JINBEN_TYPE_SUXING, chaojidou.DIFFICULTY_TYPE_MAOXIAN, 1)
		fmt.Println("shift-d end")
		endTime = time.Now()
	})

	fmt.Println("--- Please press shift + r to 每日挑战 ---")
	hook.Register(hook.KeyDown, []string{robotgo.Shift, robotgo.KeyR}, func(e hook.Event) {
		fmt.Println("shift-r")
		if endTime.Add(time.Second).After(time.Now()) {
			fmt.Println("shift-r 相隔时间太短.")
			return
		}

		var tp chaojidou.MeiRiType = chaojidou.MeiRiType(*meiriStr)
		chaojidou.NpcWaitSecs = 30
		chaojidou.ReadMapWaitSecs = 80
		captain.MeiRiTiaoZhan(tp, chaojidou.DIFFICULTY_TYPE_MAOXIAN)
		fmt.Println("shift-r end")
		endTime = time.Now()
	})

	fmt.Println("--- Please press shift + z to 追溯 ---")
	hook.Register(hook.KeyDown, []string{robotgo.Shift, robotgo.KeyZ}, func(e hook.Event) {
		fmt.Println("shift-z")
		if endTime.Add(time.Second).After(time.Now()) {
			fmt.Println("shift-z 相隔时间太短.")
			return
		}

		chaojidou.NpcWaitSecs = 30
		chaojidou.ReadMapWaitSecs = 90
		captain.JiuYunDong(chaojidou.DIFFICULTY_TYPE_SHULIAN, 1)
		fmt.Println("shift-z end")
		endTime = time.Now()
	})

	s := hook.Start()
	<-hook.Process(s)
}
