package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gangcheng1030/game_script/utils"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var eventFileName = flag.String("i", "event.txt", "input: event file name")

var eventFile *os.File
var eventReader *bufio.Reader

func initComponent() {
	flag.Parse()
	var err error

	eventFile, err = os.OpenFile(*eventFileName, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	eventReader = bufio.NewReader(eventFile)
}

func main() {
	initComponent()

	//创建监听退出chan
	c := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(c, syscall.SIGINT)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGINT:
				eventFile.Close()
				fmt.Println("回放结束.")
				os.Exit(0)
			default:
				fmt.Println("other signal", s)
			}
		}
	}()

	fmt.Println("10秒后开始回放.")
	time.Sleep(10 * time.Second)
	fmt.Println("开始回放.")
	defer func() {
		eventFile.Close()
		fmt.Println("回放结束.")
	}()

	var preTime time.Time
	first := true
	for {
		event, err := getNextEvent(eventReader)
		if err != nil {
			fmt.Printf("read event err: %v", err)
			break
		}
		if event == nil {
			break
		}

		switch event.Kind {
		case hook.KeyDown:
			k, ok := utils.Raw2key[event.Rawcode]
			if !ok {
				log.Printf("wrong rawCode: %d", event.Rawcode)
				continue
			}
			fmt.Println(k)
			if !first {
				time.Sleep(event.When.Sub(preTime))
			}
			robotgo.KeyDown(k)
			preTime = event.When
		case hook.KeyUp:
			k, ok := utils.Raw2key[event.Rawcode]
			if !ok {
				log.Printf("wrong rawCode: %d", event.Rawcode)
				continue
			}
			if !first {
				time.Sleep(event.When.Sub(preTime))
			}
			robotgo.KeyUp(k)
			preTime = event.When
		//case hook.MouseMove:
		//	if !first {
		//		dur := event.When.Sub(preTime)
		//		if dur.Milliseconds() > 100 {
		//			time.Sleep(event.When.Sub(preTime))
		//			robotgo.Move(int(event.X), int(event.Y))
		//			preTime = event.When
		//		}
		//	} else {
		//		robotgo.Move(int(event.X), int(event.Y))
		//		preTime = event.When
		//	}
		case hook.MouseDown:
			robotgo.Move(int(event.X), int(event.Y))
			if !first {
				time.Sleep(event.When.Sub(preTime))
			}
			if event.Button == 2 {
				robotgo.MouseDown("right")
			} else {
				robotgo.MouseDown()
			}
			preTime = event.When
		case hook.MouseUp:
			robotgo.Move(int(event.X), int(event.Y))
			if !first {
				time.Sleep(event.When.Sub(preTime))
			}
			if event.Button == 2 {
				robotgo.MouseUp("right")
			} else {
				robotgo.MouseUp()
			}
			preTime = event.When
		default:
			// do nothing
		}

		first = false
	}
}

func getNextEvent(reader *bufio.Reader) (*hook.Event, error) {
	for {
		line, _, errTmp := reader.ReadLine()
		if errTmp == nil {
			// 文件中可能有空白行
			if len(line) == 0 {
				continue
			}
			event := hook.Event{}
			err := json.Unmarshal(line, &event)
			if err != nil {
				return nil, err
			}
			return &event, nil
		} else if errTmp == io.EOF {
			return nil, nil
		} else {
			return nil, errTmp
		}
	}
}
