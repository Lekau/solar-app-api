package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Yes the server is running",
	})
}

func main() {
	route := gin.Default()
	route.GET("/", ping)
	route.Run("localhost:5001")
}