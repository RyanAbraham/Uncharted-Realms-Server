package main

import (
	"fmt"

	"github.com/ryanabraham/urserver/handlers"
)

func main() {
	a := handlers.App{
		MatchmakingQueue: make(chan string),
	}
	a.Initialize()
	fmt.Println("Serving on port 8080")
	a.Run(":8080")
}
