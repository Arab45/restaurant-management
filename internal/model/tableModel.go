package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type TableModel struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Number     int `bson:"number" json:"number" example:"12"`
	Capacity   int `bson:"capacity" json:"capacity" example:"4"`
	Location   string `bson:"location" json:"location" example:"Patio"`
	IsOccupied bool `bson:"is_occupied" json:"is_occupied" example:"false"`
}