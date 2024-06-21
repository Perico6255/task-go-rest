package router

import (
	"Perico6255/task-go-rest/internal/config"
	"Perico6255/task-go-rest/internal/handlers"
	"Perico6255/task-go-rest/internal/repository"
	"Perico6255/task-go-rest/internal/services"

	"github.com/gin-gonic/gin"
)


func SetUp() *gin.Engine{
	server := gin.Default()

  
  r:= repository.NewUserRepository(config.DB)
  s:= services.NewUserService(r)
  hUser := handlers.NewUserHandler(s)


	server.GET("/hello", handlers.HelloWorld)
	server.GET("/users", hUser.GetAllUsers)
	server.POST("/user", hUser.CreateUser)

  return server
}
