package routes

import (
	controller "golang-restaurant-management/contollers"
	"golang-restaurant-management/middleware"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authentication())
	incomingRoutes.GET("/foods", controller.GetFoods())
	incomingRoutes.GET("/foods/:food_id", controller.GetFood())
	incomingRoutes.POST("/foods", controller.CreatFood())
	incomingRoutes.PATCH("/foods/:food_id", controller.UpdateFood())
}
