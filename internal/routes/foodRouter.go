package routes

import(
	"github.com/gin-gonic/gin"
	controller"RESTAURANT-MANAGEMENT/internal/controller"
)
func FoodRouter(incomingRouter *gin.RouterGroup){
	incomingRouter.POST("/food", controller.CreateFood())
	incomingRouter.GET("/foods", controller.GetFoods())
	incomingRouter.GET("/food/:id", controller.GetFood())
	incomingRouter.PUT("/food/:id", controller.UpdateFood())
}