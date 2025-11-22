package routes

import(
	"github.com/gin-gonic/gin"
	controller"RESTAURANT-MANAGEMENT/internal/controller"
)

func TableRouter(incomingRouter *gin.Engine) {
	incomingRouter.POST("/table", controller.CreateTable())
	incomingRouter.GET("/tables", controller.GetTables())
	incomingRouter.GET("/table/:id", controller.GetTable())
	incomingRouter.PUT("/table/:id", controller.UpdateTable())
	incomingRouter.DELETE("/table/:id", controller.DeleteTable())
}