package contoller

import (
	"golang-restaurant-management/datebase"
	"golang-restaurant-management/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

var orderCollection *mongo.Collection = datebase.OpenCollection(datebase.Client, "order")

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := orderCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing order items"})
		}
		var allOrders []bson.M
		if err = result.All(ctx, &allOrders); err != nil {
			log.Fatal((err))

		}
		c.JSON(http.StatusOK, allOrders)
	}
}
func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		orderid := c.Param("order_id")
		var order models.Food
		err := orderCollection.FindOne(ctx, bson.M{"order_id": orderid}).Decode(&order)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "errot occured while featching the order"})
		}

		c.JSON(http.StatusOK, order)
	}
}
func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var table models.Table
		var order models.Order
		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(order)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		if order.Table_id != nil {
			err := tableCollection.FindOne(ctx, bson.M{"table_id": order.Table_id}).Decode(&table)

			defer cancel()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Table not found"})

				return
			}

		}
		order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.ID = primitive.NewObjectID()
		order.Order_id = order.ID.Hex()
		result, insertErr := orderCollection.InsertOne(ctx, order)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "order itme is not created"})
			return

		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}
func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var table models.Table
		var order models.Order
		var updateObj primitive.D
		orderId := c.Param("order_id")
		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if order.Table_id != nil {
			err := menuCollection.FindOne(ctx, bson.M{"table_id": table.Table_id}).Decode(&table)
			defer cancel()

			if err != nil {

				c.JSON(http.StatusInternalServerError, gin.H{"error": "message:menu not found"})
				return

			}
			updateObj = append(updateObj, bson.E{"menu", order.Order_id})

		}
		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at", order.Updated_at})
		upsert := true

		filter := bson.M{"order_id": orderId}
		opt := options.UpdateOptions{
			Upsert: &upsert}
		result, err := orderCollection.UpdateOne(ctx, filter, bson.D{{"$set", updateObj}}, &opt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Order item failed to update"})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func OrderItemOrderCreator(order models.Order) string {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.ID = primitive.NewObjectID()
	order.Order_id = order.ID.Hex()

	orderCollection.InsertOne(ctx, order)
	defer cancel()
	return order.Order_id

}
