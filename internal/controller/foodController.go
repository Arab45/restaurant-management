package controller

import (
	"RESTAURANT-MANAGEMENT/internal/database"
	"RESTAURANT-MANAGEMENT/internal/model"
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

	  err := foodCollection.FindOne(ctx, bson.M{"food_id", foodId}).Decode(&food)
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
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		//Pagination
		recordPerpage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerpage < 1 {
			recordPerpage = 10

		}
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil || page < 1 {
			page = 1
		}
		startIndex := (page - 1) * recordPerpage
		startIndex, err = strconv.Atoi(c.Query("startIndex"))

		matchStage := bson.D.D{{"$match", bson.D{{}}}}
		groupStage := bson.D{{"$group", bson.D{{"_id", bson.D{{"_id", "nil"}}}, {"total_count", bson.D{{"$sum", 1}}}, {"data", bson.D{{"$push", "$$ROOT"}}}}}}}
		projectStage := bson.D{{"$project", bson.D{{
			{"_id", 0},
			{"total_count", 1},
			{"food_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerpage}}}},
		}}}}

		result, err := foodCollection.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage, projectStage})
		defer cancel()
		if err != nil {
			c.JSON(
				http.StatusInternalServerError, gin.H{"error": "Error occured while listing food items"},
			)
		}
		var allFoods []bson.M
		if err = result.All(ctx, &allfoods); err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, allFoods[0])
}


func UpdateFood( ) gin.HandlerFunc{
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second);
		var menu models.Menu
		var food models.Food

		foodId := c.param("food_id");

		if err := c.BindJSON(&food); err != nil {
			c.JSON(hrrp.StatusBadRequest, gin.H("error", err.Error()));
			return
		}

		var updatedObj primitive.Decode

		if food.Name != nil{
			updatedObj = append(updatedObj, bson.E{"name", food.Name})
		}

		if food.Price != nil {
			updatedObj = append(updatedObj, bson.E{"price", food.Price})
		}

		if food.Food_image != nil {
			updatedObj = append(updatedObj, bson.E{"food_image", food.Food_image})
		} 

		if food.Menu_id != nil {
			err := menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)
			defer cancel()
			if err != nil {
				msg := fmt.Sprintf("message: Menu not found")
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}
			updateObj = append(updateObj, bson.E{"menu", food.Price})
		}

		food.Update_at = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339));
		updateObj = append(updateObj, bson.E("update_at", food.Updated_at));

		upsert := true
		filter := bson.M{"food_id", foodID}

		opt := options.UpdateOptions{
			Update: &upsert
		}

		res, err := foodCollection.Update(
			ctx,
			filter,
			bson.D{
				{"$set", updateObj}
			},
			&opt
		)

		if err != nil {
			msg := fmt.Sprint("foot item update failed")
			c.JSON(hhtp.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		c.JSON(http.StatusOK, res)
}
}

func DeleteFood( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Delete Food",
	})
}

func round( num float64) int{
	return int(num + math.Copysign(0.5, num))
}

func toFixed( num float64, precision int) float64{
	output := mat.Pow(10, float64(precision))
	return float64(round(num*output))
}
