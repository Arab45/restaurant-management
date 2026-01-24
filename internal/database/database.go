// package database

// import (
//     "context"
//     "log"
//     "os"
//     "time"

//     "go.mongodb.org/mongo-driver/mongo"
//     "go.mongodb.org/mongo-driver/mongo/options"
// )

// var Client *mongo.Client

// // Connect to MongoDB
// func ConnectDB() *mongo.Client {
//     mongoURL := os.Getenv("MONGODB_URL")

//     if mongoURL == "" {
//         mongoURL = "mongodb://localhost:27017"
//     }

//     client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
//     if err != nil {
//         log.Fatal(err)
//     }

//     ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//     defer cancel()

//     err = client.Connect(ctx)
//     if err != nil {
//         log.Fatal(err)
//     }

//     Client = client
//     return client
// }

// // Exported function — THIS is what your main.go uses
// func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
//     return client.Database("restaurant").Collection(collectionName)
// }

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
	log.Println("✅ MongoDB connected")
}

func Client() *mongo.Client {
	if client == nil {
		log.Fatal("Mongo client is not initialized")
	}
	return client
}
