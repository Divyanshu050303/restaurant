package routes

import (
	"github.com/gin-gonic/gin"

	controller "golang-restaurant-management/contollers"
	"golang-restaurant-management/middleware"
)

func InvoiceRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.Use(middleware.Authentication())
	incommingRoutes.GET("/invoices", controller.GetInvoices())
	incommingRoutes.GET("/invoices/:invoice_id", controller.GetInvoices())
	incommingRoutes.POST("/invoices", controller.CreateInvoice())
	incommingRoutes.PATCH("/invoices/:invoice_id", controller.UpdateInvoice())

}
