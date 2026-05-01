package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type UserModel struct {
	ID            primitive.ObjectID `bson:"_id"`
	FirstName     string             `json:"first_name"`
	LastName      string             `json:"last_name"`
	Password      string             `json:"password"`
	Email         string             `json:"email"`
	Avatar        string             `json:"avatar"`
	Phone         string             `json:"phone"`
	Token         string             `json:"token"`
	RefreshToken  string             `json:"refresh_token"`
	CreatedAt     time.Time          `json:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at"`
	UserID        string             `json:"user_id"`
}

// type UserModel struct {
// 	ID            primitive.ObjectID `bson:"_id"`
// 	First_name    *string            `json:"first_name" validate:"required,min=2,max=100"`
// 	Last_name     *string            `json:"last_name" validate:"required,min=2,max=100"`
// 	Password      *string            `json:"password" validate:"required,min=6"`
// 	Email         *string            `json:"email" validate:"required"`
// 	Avatar        *string            `json:"avatar"`
// 	Phone         *string            `json:"phone" validate:"required,min=10,max=15"`
// 	Token         *string            `json:"token"`
// 	Refresh_Token *string            `json:"refresh_token"`
// 	Created_at    time.Time          `json:"created_at"`
// 	Updated_at    time.Time          `json:"updated_at"`
// 	User_id       *string            `json:"user_id"`
// }

// LoginRequest is the JSON body for POST /user-login (email and password only).
type LoginRequest struct {
	Email     string `json:"email" example:"user@example.com"`
	Password  string `json:"password" example:"yourpassword"`
}
