package main 

import (
	"os"
	"context"
	"time"


	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"RESTAURANT-MANAGEMENT/internal/routes"
	"RESTAURANT-MANAGEMENT/internal/database"
	"github.com/joho/godotenv"
)

var foodCollection *mongo.Collection

func main () {
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	    // LOAD .env FILE
	    godotenv.Load()

	    // CONNECT DATABASE FIRST
		database.ConnectMongo(ctx)

		
	port := os.Getenv("PORT") 

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRouter(router)
	routes.FoodRouter(router)
	routes.OrderRouter(router)
	routes.TableRouter(router)
	routes.InvoiceRouter(router)
	routes.OrderItemRouter(router)
	routes.NoteRouter(router)

	router.Run(":" + port)
}