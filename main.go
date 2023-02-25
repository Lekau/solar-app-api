package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func helloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func main() {
	route := gin.Default()
	route.GET("/", helloWorld)
	route.Run("localhost:5001")
}