package controller

import (
	"context"
	"fmt"
	"RESTAURANT-MANAGEMENT/internal/database"
	"RESTAURANT-MANAGEMENT/internal/model"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InvoiceViewFormat struct {
	Invoice_id        string
	Payment_method    string
	Order_id          string
	Payment_status    *string
	Payment_due       interface{}
	Table_number      interface{}
	Payment_due_date  time.time
	Order_details     interface{}

}

var invoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "invoice")

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.second)

		var invoice models.Invoice 

		if err := c.BindJSON(&invoice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var order models.Order 

		err := orderCollection.FindOne(ctx, bson.M{"order_id": invoice.Order_id}).Decode(&order)
		defer cancel()

		if err != nil {
			msg := fmt.Sprintf("message: Order was not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		}

		status := "PENDING"
		if invoice.Payment_status == nil {
			invoice.Payment_status = &status
		}

		invoice.Payment_due_date, _ = time.Parse(time.RFC3339, time.Now()AddDate(0, 0, 1).Format(time.RFC3339))
		invoice.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		invoice.updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		invoice.ID = primitive.NewObjectID
		invoice.Invoice_id = invoice.ID.Hex()

		validationErr := validate.Struct(invoice)
		if validationErr != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := invoiceCollection.InsertOne(ctx, invoice)

		if err != nil{
			msg: fmt.Sprintf("invoice item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.second)
		invoiceId := c.Param("invoice_id")

		var invoice models.Invoice 

		err := invoiceCollection.FindOne(ctx, bson.M{"invoice_id": invoiceId}).Decode(&invoice)
		defer cancel()

		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing invoice item"})
		}

		var invoiceView InvoiceViewFormat 

		allOrderItem, err := ItemByOrder(invoice.Order_id)
		invoiceView.Order_id = invoice.Order_id
		invoiceView.Payment_due_date = invoice.Payment_due_date

		invoiceView.Payment_method = "null"
		if invoice.Payment_method != nil {
			invoiceView.Payment_method = *invoice.Payment_method
		}

		invoiceView.Invoice_id = invoice.Invoice_id
		invoiceView.Payment_status = *&invoice.Payment_status
		invoiceView.Payment_due = allOrderItem[0]["payment_due"]
		invoiceView.Table_number = allOrderItem[0]["table_number"]
		invoiceView.Order_details = allOrderItem[0]["order_items"]

		c.JSON(http.StatusOk, invoiceView)
}
}

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.second)

		result, err  := invoiceCollection.Find(context.TODO(), bson.M{})
		defer cancel()

		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occur"})
		}

		var allInvoice []bson.M 
		if err != result.All(ctx, &allInvoice); err != nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOk, allInvoice)
}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.second)

		var invoice models.Invoice 
		invoiceId := c.Param("invoice_id")

		if err := c.BindJSON(&invoice); err != nill {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		filter := bson.M{"invoice_id": invoiceId}

		var updatedObj primitive.D 

		if invoice.Payment_method != nil {
			updatedObj = append(updatedObj, bson.E{"payment_method", invoice.Payment_method})
		}

		if invoice.Payment_status != nill{
			updatedObj = append(updatedObj, bson.E{"payment_status", invoice.Payment_status})

		}

		invoice.Updated_at = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339));
		updatedObj = append(updatedObj, bson.E{"updated_at", inoice.Updated_at})

		upsert := true
		opt := options.UpdateOptions{
			Upsert: &upsert
		}

		status = "PENDING"
		if invoice.Payment_status == nil {
			invoice.Payment_status = &status
		}

		result, err := invoiceCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{"$set", updatedObj}
			}, 
			&opt
		)

		if err != nil{
			msg := fmt.Sprintf("invoice item update filed")
			c.JSON(status.StatusInternalServerError, gin.H{"error": msg})
		}

		defer cancel()
		c.Json(http.StatusOk, result)
}
}

func DeleteInvoice() gin.HandlerFunc {
	return func(c *gin.Context){
	c.JSON(200, gin.H{
		"result": "Delete Invoice",
	})
}
}