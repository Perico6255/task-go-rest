package main

import (
	"Perico6255/task-go-rest/internal/router"
)

func main() {
	server := router.SetUp()
	server.Run(":4000")
}
