package main

import (
	"fmt"

	"./server"
)

func main() {
	fmt.Println("Hello Server")

	running := make(chan int)
	go func() {
		server := server.Server{}
		server.SetupServer(running)
		server.StartServer()
	}()

	for i := 0; i < 1; i++ {
		_ = <-running
	}
}
