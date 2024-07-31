package routes

import (
	"github.com/gin-gonic/gin"

	controller "golang-restaurant-management/controllers"
)

func InvoiceRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/invoices", controller.GetInvoices())
	incommingRoutes.GET("/invoices/:invoice_id", controller.GetInvoices())
	incommingRoutes.POST("/invoices", controller.CreateInvoice())
	incommingRoutes.PATCH("/invoices/:invoice_id", controller.UpdateInvoice())

}
