package main

import (
	"github.com/ryanabraham/urserver/handlers"
)

func main() {
	a := handlers.App{
		MatchmakingQueue: make(chan string),
	}
	a.Initialize()
	a.Run(":8080")
}
