package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"message": "Hello World!",
		})
	})
	server.Run("localhost:8000")
	fmt.Println("Hello World!")
}