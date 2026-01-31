package routes

import(
	"github.com/gin-gonic/gin"
	controller"RESTAURANT-MANAGEMENT/internal/controller"
)
func FoodRouter(incomingRouter *gin.RouterGroup){
	// CreateFood godoc
// @Summary      Create a new food item
// @Description  Create a food item and attach it to a menu
// @Tags         Foods
// @Accept       json
// @Produce      json
// @Param        food body model.FoodModel true "Food payload"
// @Success      200 {object} mongo.InsertOneResult
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /api/v1/food [post]
	incomingRouter.POST("/food", controller.CreateFood())
	incomingRouter.GET("/foods", controller.GetFoods())
	incomingRouter.GET("/food/:id", controller.GetFood())
	incomingRouter.PUT("/food/:id", controller.UpdateFood())
}