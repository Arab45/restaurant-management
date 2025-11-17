package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type OrderModel struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID     int `bson:"user_id" json:"user_id"`
	TableID    int 	`bson:"table_id" json:"table_id"`
	FoodItems  []int `bson:"food_items" json:"food_items"`
	TotalPrice float64 `bson:"total_price" json:"total_price"`
	Status     string `bson:"status" json:"status"`
}