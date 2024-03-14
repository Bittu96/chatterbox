package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGO_URI     = "" //keys.GetKeys().MONGO_URI
	MONGO_DB_NAME = "" //eys.GetKeys().MONGO_DB_NAME
)

func MongoConnect() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(MONGO_URI),
	)

	if err != nil {
		fmt.Println("mongo connection failed")
		return nil, err
	}

	fmt.Println("mongo connection success")
	return client, nil
	// return client.Database(MONGO_DB_NAME)
}
