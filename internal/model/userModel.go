package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id" example:"60c5baaf4f1a2563f8e4d2b5"`
	FirstName string `bson:"first_name" json:"first_name" example:"John"`
	LastName  string `bson:"last_name" json:"last_name" example:"Doe"`
	Username  string `bson:"username" json:"username" example:"johndoe"`
	Email     string `bson:"email" json:"email"`
	Password  string `bson:"password" json:"password"`
	Role      string `bson:"role" json:"role" example:"admin"`
}