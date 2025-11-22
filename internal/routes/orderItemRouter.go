package routes
import(
	"github.com/gin-gonic/gin"
	controller"RESTAURANT-MANAGEMENT/internal/controller"
)

func OrderItemRouter(incomingRouter *gin.Engine) {
	incomingRouter.POST("/orderItem", controller.CreateOrderItem())
	incomingRouter.GET("/orderItems", controller.GetOrderItems())
	incomingRouter.GET("/orderItem/:id", controller.GetOrderItemByOrder())
	incomingRouter.PUT("/orderItem/:id", controller.UpdateOrderItem())
	incomingRouter.DELETE("/orderItem/:id", controller.DeleteOrderItem())
}