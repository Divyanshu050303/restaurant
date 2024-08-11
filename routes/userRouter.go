package routes

import (
	"github.com/gin-gonic/gin"

	controller "golang-restaurant-management/contollers"
)

func UserRoutes(incomintRoutes *gin.Engine) {
	incomintRoutes.GET("/users", controller.GetUsers())
	incomintRoutes.GET("/users/:user_id", controller.GetUser())
	incomintRoutes.POST("/users/signUp", controller.Signup())
	incomintRoutes.POST("/users/login", controller.Login())
}
