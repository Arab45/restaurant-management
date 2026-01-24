package routes

import(
	"github.com/gin-gonic/gin"
	controller"RESTAURANT-MANAGEMENT/internal/controller"
)

func MenuRouter(incomingRouter *gin.Engine) {
	incomingRouter.POST("/menu", controller.CreateMenu())
	incomingRouter.GET("/menus", controller.GetMenus())
	incomingRouter.GET("/menu/:id", controller.GetMenu())
	incomingRouter.PUT("/menu/:id", controller.UpdateMenu())
	// incomingRouter.DELETE("/menu/:id", controller.DeleteMenu())
}