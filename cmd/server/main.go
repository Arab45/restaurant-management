package main

import (
	"context"
	"os"
	"time"

	"RESTAURANT-MANAGEMENT/docs"
	"RESTAURANT-MANAGEMENT/internal/database"
	"RESTAURANT-MANAGEMENT/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

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
	router.Group("/api/v1")
	// REGISTER ROUTES
	docs.RegisterDocs(router)

	routes.UserRouter(router)
	routes.FoodRouter(router)
	routes.MenuRouter(router)
	routes.OrderRouter(router)
	routes.TableRouter(router)
	routes.InvoiceRouter(router)
	routes.OrderItemRouter(router)
	routes.NoteRouter(router)

	router.Run(":" + port)
}
