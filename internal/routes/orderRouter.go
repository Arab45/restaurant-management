package routes

import(
	"github.com/gin-gonic/gin"
	controller"RESTAURANT-MANAGEMENT/internal/controller"
)

func OrderRouter(incomingRouter *gin.Engine) {
	incomingRouter.POST("/order", controller.CreateOrder())
	incomingRouter.GET("/orders", controller.GetOrders())
	incomingRouter.GET("/order/:id", controller.GetOrder())
	incomingRouter.PUT("/order/:id", controller.UpdateOrder())
	incomingRouter.DELETE("/order/:id", controller.DeleteOrder())
}