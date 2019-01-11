package main

import (
	"fmt"

	"github.com/michalnov/routerTest/server"
)

func main() {
	fmt.Println("Hello Server")

	running := make(chan int, 2)
	go func() {
		server := server.Server{}
		server.SetupServer(running)
		server.StartServer()
	}()

	x := <-running
	fmt.Println("one" + string(x))
}
