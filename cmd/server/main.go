package main

import (
	"context"
	"os"
	"time"

	"RESTAURANT-MANAGEMENT/internal/database"
	"RESTAURANT-MANAGEMENT/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"

	_ "RESTAURANT-MANAGEMENT/docs"
)

// @title Restaurant Management API
// @version 1.0
// @description Restaurant Management System API
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/v1

var foodCollection *mongo.Collection

func main() {
	// LOAD .env FILE
	godotenv.Load()

	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// CONNECT DATABASE FIRST
	database.ConnectMongo(ctx)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Swagger endpoint
	router.GET("/swagger/restaurant-management/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API ROUTES (with auth)
	apiGroup := router.Group("/api/v1")
	{
		routes.UserRouter(apiGroup)
		routes.FoodRouter(apiGroup)
		routes.MenuRouter(apiGroup)
		routes.OrderRouter(apiGroup)
		routes.TableRouter(apiGroup)
		routes.InvoiceRouter(apiGroup)
		routes.OrderItemRouter(apiGroup)
		routes.NoteRouter(apiGroup)
	}

	router.Run(":" + port)
}
