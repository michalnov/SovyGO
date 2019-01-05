package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type config struct {
	Server   server   `json:"server,omitempty"`
	Database database `json:"database,omitempty"`
	Log      log      `json:"log,omitempty"`
}

//loadConfig load configuration file and return new config structure
func loadConfig() {
	fmt.Println("loading config file")
}

type server struct {
	Folder string `json:"folder,omitempty"`
	Port   string `json:"port,omitempty"`
}

type database struct {
	Master db `json:"master,omitempty"`
	Slave  db `json:"slave,omitempty"`
}

type db struct {
	Active   bool   `json:"active,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	Port     string `json:"port,omitempty"`
	Name     string `json:"name,omitempty"`
	Address  string `json:"address,omitempty"`
}

type log struct {
	Logging  bool     `json:"logging,omitempty"`
	External external `json:"external,omitempty"`
	Database database `json:"database,omitempty"`
}

type external struct {
	Active  bool   `json:"active,omitempty"`
	Address string `json:"address,omitempty"`
}

//InitializeDb produce strings for acces to database
func InitializeDb() (string, string, error) {
	absPath, _ := filepath.Abs("config/config.json")
	fmt.Println("Database Opening configuration File")
	temFile, err := ioutil.ReadFile(absPath)
	if err != nil {
		fmt.Println("Reading Db config failed")
		return "", "", err
	}
	conf := DbConf{}
	err = json.Unmarshal(temFile, &conf)
	if err != nil {
		fmt.Println("Unmarshal Db config failed")
		return "", "", err
	}
	primaryString := ""
	logString := ""
	primaryString = conf.DbUser + ":" + conf.DbPassword + "@tcp(" + conf.DbAddress + ":" + conf.DbPort + ")/" + conf.DbName
	if conf.LogDb {
		fmt.Println("LOG status: active")
		logString = conf.LogDbUser + ":" + conf.LogDbPassword + "@tcp(" + conf.LogDbAddress + ":" + conf.LogDbPort + ")/" + conf.LogDbName
	}
	fmt.Println("Database setup: succes")
	return primaryString, logString, nil
}
