package databases

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database


func InitDB(uri string, dbName string) error {
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
		if err != nil {
			return err
		}
		db = client.Database(dbName)
		return nil
}

func GetDB() *mongo.Database {
	return db
}