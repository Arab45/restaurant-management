package main 

import (
	"os";
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"RESTAURANT-MANAGEMENT/internal/routes"
	"RESTAURANT-MANAGEMENT/internal/database"
	"github.com/joho/godotenv"
)

var foodCollection *mongo.Collection

func main () {
	    // LOAD .env FILE
	    godotenv.Load()

	    // CONNECT DATABASE FIRST
		database.ConnectDB()

		// INITIALIZE COLLECTION
		foodCollection = database.OpenCollection(database.Client, "food")
		
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