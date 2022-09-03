package main

import "github.com/gin-gonic/gin"


func Router(socket string) *gin.Engine {
	instance := gin.Default()
	instance.GET("/", helloWorld)
	return instance
}