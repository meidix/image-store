package databases

import (
	"mediastore/config"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestDBIsInitilized(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Errorf("the Following error has occured:\n%v\n", err)
	}
	uri := os.Getenv(config.MONGO_URI_NAME_ENV)
	err = InitDB(uri, config.MONGO_DB_NAME)
	if err != nil {
		t.Fatalf("database is not initilized the way it should:\n%v", err)
	}
}