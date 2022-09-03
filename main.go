package main

import (
	"log"
	"mediastore/config"
	"mediastore/controllers"
	"mediastore/databases"
	"mediastore/router"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if config.USE_DB {
		dbURI := os.Getenv(config.MONGO_URI_NAME_ENV)
		dbName := os.Getenv(config.MONGO_DB_NAME)
		err = databases.InitDB(dbURI, dbName)
		if err != nil {
			log.Fatalf("there was an error while initlizing the db:\n%v", err)
		}
	}
	app := router.Router()
	app.GET("/", controllers.HelloWorld)
	app.Run(config.SERVER_ADDRESS)
}