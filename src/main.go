package main

import "github.com/gin-gonic/gin"

func HelloWorld(c *gin.Context) {
	c.IndentedJSON(200, gin.H{"message": "Hello world"})
}
func main() {
	server := gin.Default()
	server.GET("/hello", HelloWorld)
	server.Run(":4000")
}
