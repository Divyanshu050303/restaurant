package routes

import (
	"github.com/gin-gonic/gin"

	controller "golang-restaurant-management/controllers"
)

func TableRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/tables", controller.GetTables())
	incommingRoutes.GET("/tables/:tables_id", controller.GetTable())
	incommingRoutes.POST("/tables", controller.CreateTables())
	incommingRoutes.PATCH("/tables/:tables_id", controller.UpdateTables())

}
