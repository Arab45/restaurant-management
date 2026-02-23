package main

import (
	"context"
	"os"
	"time"

	"RESTAURANT-MANAGEMENT/docs" // Import without underscore to use
	"RESTAURANT-MANAGEMENT/internal/database"
	"RESTAURANT-MANAGEMENT/internal/routes"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Restaurant Management API
// @version 1.0
// @description A comprehensive Restaurant Management System API for managing users, menus, food items, orders, tables, invoices, and notes.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /api/v1
// @schemes http

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

	//CORS CONFIG
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
			"https://restaurant-management-f9kx.onrender.com",
		},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Programmatically set Swagger info
	// Support both local and online servers
	docs.SwaggerInfo.Host = "localhost:3000"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Title = "Restaurant Management API"
	docs.SwaggerInfo.Description = "Restaurant Management System API - Available on both local and online servers"
	docs.SwaggerInfo.Version = "1.0"

	// Swagger UI - This uses the embedded swagger.json automatically!
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
