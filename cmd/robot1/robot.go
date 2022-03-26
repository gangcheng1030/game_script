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
		captain.JiuYunDong(chaojidou.DIFFICULTY_TYPE_SHULIAN, 1)
		fmt.Println("shift-s end")
		endTime = time.Now()
	})

	fmt.Println("--- Please press shift + t to confirm ---")
	hook.Register(hook.KeyDown, []string{robotgo.Shift, robotgo.KeyT}, func(e hook.Event) {
		fmt.Println("shift-t")
		if endTime.Add(time.Second).After(time.Now()) {
			fmt.Println("shift-t 相隔时间太短.")
			return
		}
		captain.LiuLangTuan(chaojidou.LIULANGTUAN_TYPE_1, chaojidou.DIFFICULTY_TYPE_YINGXIONG)
		fmt.Println("shift-t end")
		endTime = time.Now()
	})

	s := hook.Start()
	<-hook.Process(s)
}
