package router

import "github.com/gin-gonic/gin"

func HelloWorld(c *gin.Context) {
	c.IndentedJSON(200, gin.H{"message": "Hello world"})
}

func SetUp() *gin.Engine{
	server := gin.Default()
	server.GET("/hello", HelloWorld)
  return server

}
