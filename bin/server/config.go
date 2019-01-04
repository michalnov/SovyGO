package server

import "fmt"

//config model of config.json file
type config struct {
}

//loadConfig load configuration file and return new config structure
func loadConfig() {
	fmt.Println("loading config file")
}
