package database

import "go.mongodb.org/mongo-driver/mongo"

const DatabaseName = "restaurant"

func Collection(name string) *mongo.Collection {
	return Client().Database(DatabaseName).Collection(name)
}
