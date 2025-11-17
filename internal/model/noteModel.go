package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type NoteModel struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Content string `bson:"content" json:"content"`
	UserID  int `bson:"user_id" json:"user_id"`
}