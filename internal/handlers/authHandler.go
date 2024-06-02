package handlers

import "github.com/gin-gonic/gin"


func HelloWorld(c *gin.Context) {
	c.IndentedJSON(200, gin.H{"message": "Hello world"})
}

