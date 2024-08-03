package contoller

import (
	"context"
	"golang-restaurant-management/datebase"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

 


func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func GetTable() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
func CreateTable() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
func UpdateTable() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
