package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type OrderItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OrderID  int `bson:"order_id" json:"order_id"`
	FoodID   int `bson:"food_id" json:"food_id"`
	Quantity int `bson:"quantity" json:"quantity"`
	Price    float64 `bson:"price" json:"price"`
}