package database

import (
    "context"
    "log"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// Connect to MongoDB
func ConnectDB() *mongo.Client {
    mongoURL := os.Getenv("MONGODB_URL")

    if mongoURL == "" {
        mongoURL = "mongodb://localhost:27017" 
    }

    client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    Client = client
    return client
}

// Exported function — THIS is what your main.go uses
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
    return client.Database("restaurant").Collection(collectionName)
}
