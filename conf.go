// Package main provides ...
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var AppConf Config

type Config struct {
	DBID       string
	DBUsername string
	DBPassword string
}

func init() {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	AppConf = Config{}
	err := decoder.Decode(&AppConf)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(AppConf)
}
