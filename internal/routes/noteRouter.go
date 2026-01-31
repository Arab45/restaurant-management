package routes

import(
	"github.com/gin-gonic/gin"
	controller"RESTAURANT-MANAGEMENT/internal/controller"
)

func NoteRouter(incomingRouter *gin.RouterGroup){
	incomingRouter.POST("/note", controller.CreateNote())
	incomingRouter.GET("/notes", controller.GetNotes())
	incomingRouter.GET("/note/:id", controller.GetNote())
	incomingRouter.PUT("/note/:id", controller.UpdateNote())
	incomingRouter.DELETE("/note/:id", controller.DeleteNote())
}