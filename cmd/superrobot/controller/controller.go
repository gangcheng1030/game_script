package main

import (
	"flag"
	"github.com/gangcheng1030/game_script/cmd/superrobot/controller/core"
	"github.com/gangcheng1030/game_script/cmd/superrobot/controller/server"
	"net/http"
)

func initComponent() {
	flag.Parse()
}

func main() {
	initComponent()
	controller := core.NewController()
	http.Handle("/controller", &server.ControllerHandler{Controller: controller})
	http.ListenAndServe(":8080", nil)
}
