package router

import (
	"Perico6255/task-go-rest/internal/handlers"

	"github.com/gin-gonic/gin"
)


func SetUp() *gin.Engine{
	server := gin.Default()

	server.GET("/hello", handlers.HelloWorld)

  return server
}
