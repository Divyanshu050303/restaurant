package contoller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOrderItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
func GetOrderItemsBYOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
func ItemByOrder(id string)(OrderItems[]primitive.M, err){
}
func GetOrderItems() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
func CreateOrderItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
func UpdateOrderItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
