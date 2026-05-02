package controller

import (
	"RESTAURANT-MANAGEMENT/internal/database"
	"RESTAURANT-MANAGEMENT/internal/helper"
	"RESTAURANT-MANAGEMENT/internal/model"
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// GetUsers godoc
// @Summary Get all users with pagination
// @Description Retrieve a list of all users with pagination support
// @Tags User
// @Produce json
// @Param recordPerPage query int false "Number of records per page (default: 10)"
// @Param page query int false "Page number (default: 1)"
// @Success 200 {object} map[string]interface{} "List of users"
// @Router /users [get]
func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var userCollection = database.Collection("users")

		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}

		page, err1 := strconv.Atoi(c.Query("page"))
		if err1 != nil || page < 1 {
			page = 1
		}

		startIndex := (page - 1) * recordPerPage

		matchStage := bson.D{{Key: "$match", Value: bson.D{}}}
		projectStage := bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "id", Value: 0},
				{Key: "total_count", Value: 1},
				{Key: "user_items", Value: bson.D{{Key: "$slice", Value: []interface{}{"$date", startIndex, recordPerPage}}}},
			}},
		}

		result, err := userCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage,
			projectStage,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing to user item"})
			return
		}

		var allUsers []bson.M
		if err = result.All(ctx, &allUsers); err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, allUsers[0])

		//either pass an error

		//identify want to return all users based on the various query parsed
	}
}

// GetUser godoc
// @Summary Get a specific user by ID
// @Description Retrieve user details by providing the user ID
// @Tags User
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} model.UserModel "User details"
// @Router /user/{id} [get]
func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var userCollection = database.Collection("users")

		userId := c.Param("user_id")
		var user model.UserModel

		err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing to user items"})
		}

		c.JSON(http.StatusOK, user)
	}
}


// SignUp godoc
// @Summary User registration/sign up
// @Description Create a new user account with email, password, phone and personal information
// @Tags User
// @Accept json
// @Produce json
// @Param user body model.UserModel true "User data (first_name, last_name, email, password, phone required)"
// @Success 200 {object} map[string]interface{} "User created successfully with ID"
// @Router /user [post]
func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var userCollection = database.Collection("users")
		var user model.UserModel

		//convert the JSON data coming from postman to something that golang understand
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		//validate the data based on user struct
		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		//you'll check if the email has already been used by another user
		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the email"})
			return
		}

		//hash password
		password := HashPasssword(user.Password)
		user.Password = password

		//you'll also check if the phone no. has already been used by another
		count, err = userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})

		defer cancel()

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the phone"})
			return
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "This email or phone number already exists"})
		}
		//create some extra details for the user object - created_at, updated_at, ID
		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		user.ID = primitive.NewObjectID()
		hexID := user.ID.Hex()
		user.UserID = hexID

		//generate token and refresh token (generate all token function from helper )
		token, refreshToken, err := helper.GenerateAllToken(
			user.Email,
			user.FirstName,
			user.LastName,
			user.UserID,
		)

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while generating token"})
			return
		}

		user.Token = token
		user.RefreshToken = refreshToken
		//if all ok, then you insert this new user into the user collection
		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			msg := "user item was not created"
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()

		//return status ok and send the result back
		c.JSON(http.StatusOK, resultInsertionNumber)
	}
}

// LogIn godoc
// @Summary User login/authentication
// @Description Authenticate user with email and password, and return access and refresh tokens
// @Tags User
// @Accept json
// @Produce json
// @Param credentials body model.LoginRequest true "Email and password"
// @Success 200 {object} model.UserModel "Login successful with user details and tokens"
// @Router /user-login [post]
func LogIn() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		userCollection := database.Collection("users")

		var loginData model.LoginRequest
		var foundUser model.UserModel

		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid input data",
			})
			return
		}

		err := userCollection.FindOne(ctx, bson.M{"email": loginData.Email}).Decode(&foundUser)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid email or password",
			})
			return
		}

		passwordIsValid, msg := VerifyPassword(loginData.Password, foundUser.Password)
		if !passwordIsValid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": msg,
			})
			return
		}

		// Generate tokens
		token, refreshToken, err := helper.GenerateAllToken(
			foundUser.Email,
			foundUser.FirstName,
			foundUser.LastName,
			foundUser.UserID,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to generate tokens",
			})
			return
		}

		// Save tokens
		helper.UpdateAllTokens(token, refreshToken, foundUser.UserID)

		// SET COOKIE (THIS IS WHAT YOU WERE MISSING)
		c.SetCookie(
			"authentication", // name
			token,          // value
			3600,           // maxAge (seconds)
			"/",            // path
			"",             // domain (empty = current domain)
			false,          // secure (true in production HTTPS)
			true,           // httpOnly (VERY IMPORTANT)
		)

		// Optional: refresh token cookie
		c.SetCookie(
			"refresh_token",
			refreshToken,
			7*24*3600, // 7 days
			"/",
			"",
			false,
			true,
		)

		foundUser.Password = ""

		c.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
			"user":    foundUser,
		})
	}
}


// func LogIn() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 		defer cancel()

// 		userCollection := database.Collection("users")

// 		var loginData model.LoginRequest
// 		var foundUser model.UserModel

// 		// Bind request
// 		if err := c.ShouldBindJSON(&loginData); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": "Invalid input data",
// 			})
// 			return
// 		}

// 		// Find user by email
// 		err := userCollection.FindOne(ctx, bson.M{"email": loginData.Email}).Decode(&foundUser)
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{
// 				"error": "Invalid email or password",
// 			})
// 			return
// 		}

// 		log.Printf("User found: %s", foundUser.Email)

// 		// Verify password
// 		passwordIsValid, msg := VerifyPassword(loginData.Password, foundUser.Password)	
// 		if !passwordIsValid {
// 			c.JSON(http.StatusUnauthorized, gin.H{
// 				"error": msg,
// 			})
// 			return
// 		}

// 		// Generate tokens
// 		token, refreshToken, err := helper.GenerateAllToken(
// 			foundUser.Email,
// 			foundUser.FirstName,
// 			foundUser.LastName,
// 			foundUser.UserID,
// 		)

// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"error": "Failed to generate tokens",
// 			})
// 			return
// 		}

// 		// Save tokens
// 		helper.UpdateAllTokens(token, refreshToken, foundUser.UserID)
// 		// if err != nil {
// 		// 	c.JSON(http.StatusInternalServerError, gin.H{
// 		// 		"error": "Failed to update tokens",
// 		// 	})
// 		// 	return
// 		// }

// 		// OPTIONAL: remove password before sending response
// 		foundUser.Password = ""

// 		// Success response
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "Login successful",
// 			"user":    foundUser,
// 			// "token":   token,
// 		})
// 	}
// }


// func LogIn() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 		defer cancel()
// 		var userCollection = database.Collection("users")
// 		var user model.UserModel
// 		var foundUser model.UserModel

// 		log.Printf("Attempting to find user with email: %s", *user.Email)

// 		//convert the login data from postman which is in JSON to golang readable format
// 		if err := c.BindJSON(&user); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

		

// 		//find a user with that email and see if that user exists
// 		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
// 		// err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&user)
// 		// defer cancel()
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found, login seems to be incorrect"})
// 			return
// 		}

// 		log.Printf("User found: %v", user)
// 		//then you will verify the password
// 		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
// 		log.Printf("Password verification result: %v, Message: %s", passwordIsValid, msg)
// 		if passwordIsValid != true {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
// 			return
// 		}
// 		//if all goes well, then you'll generate tokens
// 		token, refreshToken, _ := helper.GenerateAllToken(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, *foundUser.User_id)

// 		//update tokens - token and refresh token
// 		helper.UpdateAllTokens(token, refreshToken, *foundUser.User_id)

// 		//return statusOk
// 		c.JSON(http.StatusOK, foundUser)
// 	}
// }

func HashPasssword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

func VerifyPassword(userPassword string, providePassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providePassword), []byte(userPassword))
	check := true
	msg := ""
	if err != nil {
		msg = "login or password is incorrect"
		check = false
	}
	return check, msg
}
