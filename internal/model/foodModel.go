package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type FoodModel struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string `bson:"name" json:"name"`
	Calories    int `bson:"calories" json:"calories"`
	Ingredients []string `bson:"ingredients" json:"ingredients"`
	Category    string `bson:"category" json:"category"`
}