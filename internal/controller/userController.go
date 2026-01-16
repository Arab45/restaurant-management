package controller

import (
	"RESTAURANT-MANAGEMENT/internal/database"
	"RESTAURANT-MANAGEMENT/internal/model"
	"context"
	"log"
	"time"
	"strconv"
	"net/http"


	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"

)


var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}

		 page, err1 := strconv.Atoi(c.Query("page"))
		 if err1 != nil || page < 1 {
			page = 1
		 }

		 startIndex := (page-1) * recordPerPage

		startIndex, err := strconv.Atoi(c.Query("startIndex"))

		matchStage := bson.D{{"$match", bson.D{{}}}}
		projectStage := bson.D{
			{"$project", bson.D{
				{"id", 0},
				{"total_count", 1},
				{"user_items", bson.D{{"$slice", []interface{}{"$date", startIndex, recordPerPage}}}}
			}}
		}

		result, err := userCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage,
			projectStage
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing to user item"})
		}

		var allUsers []bson.M 
		if err = result.All(ctx, &allUsers); err != nil{
			log.Fatal(err)
		}

		c.JSON(http.statusOk, allUsers[0])

		//either pass an error

		//identify want to return all users based on the various query parsed
}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		
		userId := c.Param("user_id")
		var user models.user

		err := userCollection.findOne(ctx, bson.M{"user_id": userId}).Decode(&user)

		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing to user items"})
		}

		c.JSON(http.statusOk, user)
	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Update User",
	})	
	}
}


func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
//convert the JSON data coming from postman to something that golang understand

//validate the data based on user struct

//you'll check if the email has already been used by another user

//hash password

//you'll also check if the phone no. has already been used by another

//create some extra details for the user object - created_at, updated_at, ID

//generate token and refresh token (generate all token function from helper )

//if all ok, then you insert this new user into the user collection

//return status ok and send the result back
	}
}

func LogIn() gin.HandlerFunc {
	return func(c *gin.Context) {
//convert the login data from postman which is in JSON to golang readable format

//find a user with that email and see if that user exists

//then you will verify the password

//if all goes well, then you'll generate tokens

//update tokens - token and refresh token

//return statusOk
	}
}

func HashPasssword(password string) string{
}

func VerifyPassword(userPassword string, providePassword string)(bool, string){

}