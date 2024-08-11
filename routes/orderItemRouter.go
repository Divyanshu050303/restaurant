package routes

import (
	"github.com/gin-gonic/gin"

	controller "golang-restaurant-management/contollers"
	"golang-restaurant-management/middleware"
)

func OrderItemRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.Use(middleware.Authentication())
	incommingRoutes.GET("/orderItems", controller.GetOrderItems())
	incommingRoutes.GET("/orderItems/:orderItem_id", controller.GetOrderItem())
	incommingRoutes.GET("/orderItems-order/:orderItem_id", controller.GetOrderItemsBYOrder())
	incommingRoutes.POST("/orderItems", controller.CreateOrderItem())
	incommingRoutes.PATCH("/orderItems/:orderItem_id", controller.UpdateOrderItem())

}
