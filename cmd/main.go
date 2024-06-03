package main

import (
	"Perico6255/task-go-rest/internal/config"
	"Perico6255/task-go-rest/internal/router"
)

func main() {
  config.ConnectDatabase()
	server := router.SetUp()

	server.Run(":4000")
}
