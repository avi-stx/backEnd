package main

import (
	"Backend_side/config"
	"Backend_side/utils"
)

func init() {
	err := config.ReadConfig("config/local_server.json")
	if err != nil {
		panic(err)
	}

}

func main() {
	router := utils.InitServer()
	router.Run("localhost:8080")
}
