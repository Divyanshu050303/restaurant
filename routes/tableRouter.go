package routes

import (
	"github.com/gin-gonic/gin"

	controller "golang-restaurant-management/contollers"
	"golang-restaurant-management/middleware"
)

func TableRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.Use(middleware.Authentication())
	incommingRoutes.GET("/tables", controller.GetTables())
	incommingRoutes.GET("/tables/:tables_id", controller.GetTable())
	incommingRoutes.POST("/tables", controller.CreateTable())
	incommingRoutes.PATCH("/tables/:tables_id", controller.UpdateTable())

}
