package controller

import (
	"RESTAURANT-MANAGEMENT/internal/database"
	"RESTAURANT-MANAGEMENT/internal/model"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
var validate = validator.New()

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var food model.FoodModel
	var menu model.MenuModel

	if err := c.BindJSON(&food); err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{"error": err.Error()},
		)
		return
	}
	validationErr := validate.Struct(food)

	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	err := menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)

	if err != nil {
		msg := fmt.Sprintf("menu was not found")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	food.Created_at = time.Now()
	food.Updated_at = time.Now()
	food.ID = primitive.NewObjectID()
	food.Food_id = food.ID.Hex()
	var num = toFixed(*food.Price, 2)
	food.Price = &num

	result, insertErr := foodCollection.InsertOne(ctx, food)

	if insertErr != nil {
		msg := fmt.Sprintf("food item was not created")
		c.JSON(
			http.StatusInternalServerError, gin.H{"error": msg},
		)
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, result)
}

}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
      var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	  foodId := c.Param("food_id")
	  var food model.FoodModel

	  err := foodCollection.FindOne(ctx, bson.M("food_id", foodId)).Decode(&food)
	  defer cancel()
	  if err != nil {
		c.JSON(
			http.StatusInternalServerError, gin.H{"error": "error occured while fetching the food item"},
		)
	  }
	  c.JSON(http.StatusOK, food)
}
}

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Get Foods",
	})
}
}

func UpdateFood( ) gin.HandlerFunc{
	return func(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Update Food",
	})
}
}

func DeleteFood( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Delete Food",
	})
}

func round( num float64) int{
	return int(num + 0.5)
}

func toFixed( num float64, precision int) float64{
	return 0.0234
}
