package routes

import (
	"github.com/gin-gonic/gin"

	controller "golang-restaurant-management/controllers"
)

func UserRoutes(incomintRoutes *gin.Engine) {
	incomintRoutes.GET("/users", controller.GetUser())
	incomintRoutes.GET("/users/:user_id", controller.GetUser())
	incomintRoutes.POST("/users/signUp", controller.SignUp())
	incomintRoutes.POST("/users/login", controller.Login())
}
