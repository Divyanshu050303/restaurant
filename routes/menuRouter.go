package routes

import (
	"github.com/gin-gonic/gin"

	controller "golang-restaurant-management/contollers"
	"golang-restaurant-management/middleware"
)

func MenuRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.Use(middleware.Authentication())
	incommingRoutes.GET("/menu", controller.GetMenu())
	incommingRoutes.GET("/menu/:menu_id", controller.GetMenu())
	incommingRoutes.POST("/menu", controller.CreateMenu())
	incommingRoutes.PATCH("/menu/:menu_id", controller.UpdateMenu())

}
