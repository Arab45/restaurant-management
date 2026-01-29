package database

import (
	"context"
	"log"
	"os"
	"time"


	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectMongo(ctx context.Context) {
	uri := os.Getenv("MONGODB_URL")

	opts := options.Client().
		ApplyURI(uri).
		SetMaxPoolSize(20).
		SetMinPoolSize(5).
		SetConnectTimeout(10 * time.Second)

	c, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal("Mongo connect error:", err)
	}

	if err := c.Ping(ctx, nil); err != nil {
		log.Fatal("Mongo ping failed:", err)
	}

	client = c
	log.Println("MongoDB connected")
}

func Client() *mongo.Client {
	if client == nil {
		log.Fatal("Mongo client is not initialized")
	}
	return client
}
