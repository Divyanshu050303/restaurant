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

type InvoiceViewFormat struct {
	Invoice_id       string
	Payment_method   string
	Order_id         string
	Payment_status   *string
	Payment_due      interface{}
	Table_number     interface{}
	Paymnet_due_date time.Time
	Order_details    interface{}
}

var invoiceCollection *mongo.Collection = datebase.OpenCollection(datebase.Client, "invoice")

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := invoiceCollection.Find(context.TODO(), bson.M{})
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing invoice"})

		}
		var allInvoice []bson.M
		if err = result.All(ctx, &allInvoice); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allInvoice)
	}
}
func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
var ctx, cancel =context.WithTimeout(context.Background(), 100*time.Second)

	}
}
func CreateInvoice() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
func UpdateInvoice() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
