package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type InvoiceModel struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OrderID     int `bson:"order_id" json:"order_id"`
	TotalAmount float64 `bson:"total_amount" json:"total_amount"`
	PaidAmount  float64 `bson:"paid_amount" json:"paid_amount"`
	DueAmount   float64 `bson:"due_amount" json:"due_amount"`
	Status      string `bson:"status" json:"status"`
}