package routes

import(
	"github.com/gin-gonic/gin"
	controller"RESTAURANT-MANAGEMENT/internal/controller"
)
func InvoiceRouter(incomingRouter *gin.RouterGroup){
	incomingRouter.POST("/invoice", controller.CreateInvoice())
	incomingRouter.GET("/invoices", controller.GetInvoices())
	incomingRouter.GET("/invoice/:id", controller.GetInvoice())
	incomingRouter.PUT("/invoice/:id", controller.UpdateInvoice())
	incomingRouter.DELETE("/invoice/:id", controller.DeleteInvoice())
}