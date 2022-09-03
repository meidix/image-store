package router

import (
	"github.com/gin-gonic/gin"
)


func Router() *gin.Engine {
	instance := gin.Default()
	return instance
}