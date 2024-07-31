package routes

import (
	"github.com/gin-gonic/gin"

	controller "golang-restaurant-management/controllers"
)

func FoodRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/foods", controller.GetFoods())
	incommingRoutes.GET("/foods/:food_id", controller.GetFoods())
	incommingRoutes.POST("/foods", controller.CreatFood())
	incommingRoutes.PATCH("/foods/:food_id", controller.UpdateFood())
}
