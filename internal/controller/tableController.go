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
	"go.mongodb.org/mongo-driver/bson/options"
	"github.com/go-playground/validator/v10"
)

var tableCollection *mongo.Collection = database.OpenCollection(database.Client, "table")
var validate = validator.New()


func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context){
		var err, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var table models.Table 

		if err := c.BindJSON(&table); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(table)

		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		table.Created_at = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		table.Updated_at = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		table.ID = primitive.NewObjectID()
		table.Order_id = table.ID.Hex()

		result, insertErr := tableCollection.InsertOne(ctx, table)

		if insertErr != nil {
			msg: fmt.Sprintf("Table item was not created")
			c.JSON(http.StatusInternalServerError. gin.H{"error": msg})
			return
		}
		defer cancel()

		c,JSON(http.StatusOK, result)

}
}

func GetTable() gin.HandlerFunc {
	return func(c *gin.Context){
		var err, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		tableId := c.Param("table_id")
		var table model.Table

		err := tableCollection.FindOne(ctx, bson.M{"table_id", tableId}).Decode(&table)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occure while fetching the tables"})
		}

		c.JSON(http.StatusOK, table)
}
}

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.second)

		result, err  := tableCollection.Find(context.TODO(), bson.M{})
		defer cancel()

		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occur"})
		}

		var allTables []bson.M 
		if err != result.All(ctx, &allTables); err != nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOk, allTables)
}
}

func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context){
		var err, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var table models.Table 

		tableId := c.Param("table_id")

		if err := c.BindJSON(&table); err != nil {
			c.JSON(http.StatusBadRequest. gin.H{"error": err.Error()})
			return
		}

		var updatedObj primitive.D  

		if table.Number_of_guests != nil {
			updatedObj = append(updateObj, bson.E{"number_of_guests", table.Number_of_guests})
		}

		if table.Table_number != nil {
			updatedObj = append(updateObj, bson.E{"table_number", table.Table_number})

		}

		table.Updated_at = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		upsert := true
		opt := option.UpdateOptios{
			Upsert: &upsert
		}

		filter := bson.M{"table_id": tableId}

		result, err := tableCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{"$set", updatedObj}
			},
			&opt
		)

		if err != nil {
			msg := fmt.Sprintf("table item update failed")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, result)

}
}

func DeleteTable() gin.HandlerFunc {
	return func(c *gin.Context){
	c.JSON(200, gin.H{
		"result": "Delete Table",
	})	
}		
}