package main

import hook "github.com/robotn/gohook"

func main() {
	evChan := hook.Start()
	defer hook.End()
}
