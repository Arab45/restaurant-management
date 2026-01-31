package routes

import (
	controller "RESTAURANT-MANAGEMENT/internal/controller"
	_ "RESTAURANT-MANAGEMENT/internal/model"
	"github.com/gin-gonic/gin"
	_ "go.mongodb.org/mongo-driver/mongo"
)

func FoodRouter(incomingRouter *gin.RouterGroup) {
	incomingRouter.POST("/food", controller.CreateFood())
	incomingRouter.GET("/foods", controller.GetFoods())
	incomingRouter.GET("/food/:id", controller.GetFood())
	incomingRouter.PUT("/food-update/:id", controller.UpdateFood())
}
