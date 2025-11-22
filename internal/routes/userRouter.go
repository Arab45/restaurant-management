package routes

import(
	"github.com/gin-gonic/gin"
	controller"RESTAURANT-MANAGEMENT/internal/controller"
)

func UserRouter(incomingRouter *gin.Engine) {
	incomingRouter.POST("/user", controller.SignUp())
	incomingRouter.GET("/users", controller.GetUsers())
	incomingRouter.GET("/user/:id", controller.GetUser())
	incomingRouter.PUT("/user/:id", controller.UpdateUser())
	incomingRouter.DELETE("/user/:id", controller.LogIn())
}