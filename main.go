package main

import (
	"fmt"
	"mediastore/config"
	"mediastore/controllers"
	"mediastore/router"
)

func main() {
	app := router.Router()
	app.GET("/", controllers.HelloWorld)
	app.Run(config.ServerAddress)
	fmt.Println("Hello World!")
}