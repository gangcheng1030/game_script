package main

import (
	"encoding/json"
	"flag"
	"github.com/gangcheng1030/game_script/cmd/superrobot/client"
	"github.com/gangcheng1030/game_script/cmd/superrobot/model"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var controllerAddr = flag.String("addr", "127.0.0.1:8080", "controller addr")
var configPath = flag.String("conf", "superrobot.json", "config")

func main() {
	flag.Parse()
	var err error
	var binDir string

	binDir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	file, err := ioutil.ReadFile(filepath.Join(binDir, *configPath))
	if err != nil {
		panic(err)
	}

	events := &Events{}
	err = json.Unmarshal(file, events)
	if err != nil {
		panic(err)
	}

	for _, event := range events.Events {
		err = client.AddEvent(*controllerAddr, event)
		if err != nil {
			log.Printf("addEvent err: %v", err)
		}
	}
}

type Events struct {
	Events []*model.Event
}
