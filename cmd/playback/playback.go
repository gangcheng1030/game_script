package main

import (
	"bufio"
	"chenggang.idea/game_script/utils"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var eventFileName = flag.String("-i", "event.txt", "input: event file name")

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

	fmt.Println("5秒后开始回放.")
	time.Sleep(5 * time.Second)
	fmt.Println("开始回放.")
	defer func() {
		eventFile.Close()
		fmt.Println("回放结束.")
	}()

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
		case hook.KeyUp:
			err = robotgo.KeyUp(utils.CodeToKey[event.Keycode])
			if err != nil {
				fmt.Printf("KeyUp err: %v", event)
			}
		case hook.KeyDown:
			if event.Keychar == 65535 {
				err = robotgo.KeyDown(utils.CodeToKey[event.Keycode])
			} else {
				fmt.Println(event)
				fmt.Println(string(event.Keychar))
				//err = robotgo.KeyDown(string(event.Keychar))
				err = robotgo.KeyToggle("enter")
			}

			if err != nil {
				fmt.Printf("KeyDown err: %v", event)
			}
		default:
			// do nothing
		}
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
