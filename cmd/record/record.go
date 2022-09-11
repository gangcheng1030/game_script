package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	hook "github.com/robotn/gohook"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var eventFileName = flag.String("o", "event.txt", "file name which save events.")

var eventFile *os.File
var eventWriter *bufio.Writer

func initComponent() {
	flag.Parse()
	var err error

	eventFile, err = os.OpenFile(*eventFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	eventWriter = bufio.NewWriter(eventFile)
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
				eventWriter.Flush()
				eventFile.Close()
				fmt.Println("录制结束.")
				os.Exit(0)
			default:
				fmt.Println("other signal", s)
			}
		}
	}()

	fmt.Println("5秒后开始录制.")
	time.Sleep(5 * time.Second)
	fmt.Println("开始录制.")
	defer func() {
		eventWriter.Flush()
		eventFile.Close()
		fmt.Println("录制结束.")
	}()

	evChan := hook.Start()
	defer hook.End()

	for ev := range evChan {
		data, _ := json.Marshal(ev)
		eventWriter.WriteString(string(data))
		eventWriter.WriteString("\n")
	}
}
