package main

import (
	"fmt"
	"github.com/vcaesar/keycode"

	hook "github.com/robotn/gohook"
)

func main() {
	add()
	//low()
	//event()
}

func add() {
	fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
	hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("ctrl-shift-q")
		hook.End()
	})

	for k := range keycode.Keycode {
		hook.Register(hook.KeyDown, []string{k}, func(e hook.Event) {
			c := hook.RawcodetoKeychar(e.Rawcode)
			fmt.Println(c)
			//fmt.Printf("KeyCode: %d, Raw: %d, KeyChar: %s\n", e.Keycode, e.Rawcode, hook.RawcodetoKeychar(e.Rawcode))
		})
	}

	//fmt.Println("--- Please press w---")
	//hook.Register(hook.KeyDown, []string{"w"}, func(e hook.Event) {
	//	fmt.Println("w")
	//})

	hook.Register(hook.MouseDown, []string{}, func(e hook.Event) {
		fmt.Println(e)
	})

	s := hook.Start()
	<-hook.Process(s)
}

func low() {
	evChan := hook.Start()
	defer hook.End()

	for ev := range evChan {
		fmt.Println("hook: ", ev)
	}
}

func event() {
	ok := hook.AddEvents("q", "ctrl", "shift")
	if ok {
		fmt.Println("add events...")
	}

	//keve := hook.AddEvent("k")
	//if keve {
	//	fmt.Println("you press... ", "k")
	//}

	mleft := hook.AddEvent("mleft")
	if mleft {
		fmt.Println("you press... ", "mouse left button")
	}
}
