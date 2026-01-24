package controller

import (
	"RESTAURANT-MANAGEMENT/internal/database"
	"RESTAURANT-MANAGEMENT/internal/model"
	"context"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Validate = validator.New()

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
        var menuCollection = database.Collection("menus")
		var foodCollection = database.Collection("foods")
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
			msg := "menu was not found"
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
			msg := "food item was not created"
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
		var foodCollection = database.Collection("foods")
		foodId := c.Param("food_id")
		var food model.FoodModel

		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)
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
		defer cancel()

		var foodCollection = database.Collection("foods")
		// Pagination
		recordPerPage, err := strconv.Atoi(c.DefaultQuery("recordPerPage", "10"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}

		page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
		if err != nil || page < 1 {
			page = 1
		}

		startIndex := (page - 1) * recordPerPage

		matchStage := bson.D{{
			Key:   "$match",
			Value: bson.D{},
		}}

		groupStage := bson.D{{
			Key: "$group",
			Value: bson.D{
				{Key: "_id", Value: nil},
				{Key: "total_count", Value: bson.D{
					{Key: "$sum", Value: 1},
				}},
				{Key: "data", Value: bson.D{
					{Key: "$push", Value: "$$ROOT"},
				}},
			},
		}}

		projectStage := bson.D{{
			Key: "$project",
			Value: bson.D{
				{Key: "_id", Value: 0},
				{Key: "total_count", Value: 1},
				{Key: "food_items", Value: bson.D{
					{Key: "$slice", Value: []interface{}{"$data", startIndex, recordPerPage}},
				}},
			},
		}}

		result, err := foodCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage,
			groupStage,
			projectStage,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error occurred while listing food items",
			})
			return
		}
		defer result.Close(ctx)

		var allFoods []bson.M
		if err := result.All(ctx, &allFoods); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error decoding food items",
			})
			return
		}

		if len(allFoods) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"total_count": 0,
				"food_items":  []interface{}{},
			})
			return
		}

		c.JSON(http.StatusOK, allFoods[0])
	}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var foodCollection = database.Collection("foods")
        var menuCollection = database.Collection("menus")

		var menu model.MenuModel
		var food model.FoodModel

		foodId := c.Param("food_id")

		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		updateObj := bson.D{}

		if food.Name != nil {
			updateObj = append(updateObj, bson.E{
				Key:   "name",
				Value: food.Name,
			})
		}

		if food.Price != nil {
			updateObj = append(updateObj, bson.E{
				Key:   "price",
				Value: food.Price,
			})
		}

		if food.Food_image != nil {
			updateObj = append(updateObj, bson.E{
				Key:   "food_image",
				Value: food.Food_image,
			})
		}

		if food.Menu_id != nil {
			err := menuCollection.
				FindOne(ctx, bson.M{"menu_id": food.Menu_id}).
				Decode(&menu)

			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"error": "menu not found",
				})
				return
			}

			updateObj = append(updateObj, bson.E{
				Key:   "menu_id",
				Value: food.Menu_id,
			})
		}

		updateObj = append(updateObj, bson.E{
			Key:   "updated_at",
			Value: time.Now(),
		})

		filter := bson.M{"food_id": foodId}

		opts := options.Update().SetUpsert(true)

		res, err := foodCollection.UpdateOne(
			ctx,
			filter,
			bson.D{{
				Key:   "$set",
				Value: updateObj,
			}},
			opts,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "food item update failed",
			})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

func DeleteFood(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Delete Food",
	})
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num * output))
}
