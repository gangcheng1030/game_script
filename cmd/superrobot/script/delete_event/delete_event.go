package main

import (
	"flag"
	"github.com/gangcheng1030/game_script/cmd/superrobot/client"
	"github.com/gangcheng1030/game_script/cmd/superrobot/model"
	"log"
)

var controllerAddr = flag.String("addr", "127.0.0.1:8080", "controller addr")
var eventId = flag.Int64("e", 1, "event id")
var nodeId = flag.String("n", "", "node id")
var groupId = flag.String("g", "", "group id")
var eventType = flag.Int("t", 1, "event type")

func main() {
	flag.Parse()

	err := client.DeleteEvent(*controllerAddr, *eventId, *nodeId, *groupId, model.EventType(*eventType))
	if err != nil {
		log.Printf("deleteEvent err: %v", err)
		return
	}
	log.Println("deleteEvent success.")
}
