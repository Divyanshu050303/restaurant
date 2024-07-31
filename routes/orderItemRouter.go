package routes

import (
	"github.com/gin-gonic/gin"

	controller "golang-restaurant-management/controllers"
)

func OrderItemRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/orderItems", controller.GetItems())
	incommingRoutes.GET("/orderItems/:orderItem_id", controller.GetItem())
	incommingRoutes.GET("/orderItems-order/:orderItem_id", controller.GetOrderItemByOrder())
	incommingRoutes.POST("/orderItems", controller.CreateItem())
	incommingRoutes.PATCH("/orderItems/:orderItem_id", controller.UpdateItem())

}
