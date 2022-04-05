package main

import (
	"flag"
	"fmt"
	"github.com/gangcheng1030/game_script/chaojidou"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"time"
)

var captainStr = flag.String("c", "official_client", "client of captain")
var start = flag.String("s", "ysyzz", "开始剧情")

var captain chaojidou.ChaoJiDou
var juqingArray []chaojidou.JuQingType

func initComponent() {
	flag.Parse()
	var err error
	captain, err = chaojidou.Build(chaojidou.ClientType(*captainStr))
	if err != nil {
		panic(err)
	}

	juqingArray = []chaojidou.JuQingType{chaojidou.JUQING_TYPE_YISHENGYUZHAOZE}
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
		startIndex := chaojidou.SearchJuqing(juqingArray, chaojidou.JuQingType(*start))
		if startIndex == -1 {
			fmt.Println("没有该剧情.")
			return
		}
		for i := startIndex; i < len(juqingArray); i++ {
			captain.JuQing(juqingArray[i])
		}
	})

	s := hook.Start()
	<-hook.Process(s)
}
