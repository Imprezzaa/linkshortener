package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB takes the connection string environment variable and opens a connection to the DB and returns a mongo.Client
func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

var DB *mongo.Client = ConnectDB()

// GetCollections allows you pull a collection from the DB
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("linkshortener").Collection(collectionName)
	return collection
}
