package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `json:"name" validation:"required"`
	Category   string             `json:"category" validation:"required"`
	Start_date *time.Time         `json:"start_date"`
	End_date   *time.Time         `json:"end_date"`
	Create_at  time.Time          `json:"create_at"`
	Update_at  time.Time          `json:"update_at"`
	Menu_id    time.Time          `json:"menu_id"`
}
