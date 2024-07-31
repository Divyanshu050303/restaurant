package contoller

import "github.com/gin-gonic/gin"

func GetUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
func SignUp() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func HashPassword() string {
}

func VerifyPassword(userPassword string, providePassword string) (bool, string) {
}
