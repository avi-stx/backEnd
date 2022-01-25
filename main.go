package main

import (
	"Backend_side/config"
	"Backend_side/utils"
)

func init() {
	if err := config.ReadConfig("config/local_server.json"); err != nil {
		panic(err)
	}
}

func main() {
	router := utils.InitServer()
	router.Run(":8080")
}

//
