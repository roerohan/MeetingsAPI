package models

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


// MongoConnect returns a client through which you can connect to MongoDB
func MongoConnect(mongoURL string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(mongoURL)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[INFO] Connected to MongoDB!")
	return client
}
