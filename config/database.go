package config

import (
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func NewDatabase() (*mongo.Client, *mongo.Database) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
	)
	serveApi := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serveApi)
	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}
	db := client.Database(os.Getenv("DB_NAME"))
	return client, db
}
