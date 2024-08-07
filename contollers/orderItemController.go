package contoller

import (
	"context"
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
)

type OrderItemPack struct {
	Table_id    *string
	Order_items []models.OrderItem
}

var orderItemCollection *mongo.Collection = datebase.OpenCollection(datebase.Client, "orderItem")

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		result, err := orderItemCollection.Find(context.TODO(), bson.M{})
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing order items"})
			return
		}
		var allOrderItem []bson.M
		if err = result.All(ctx, &allOrderItem); err != nil {
			log.Fatal(err)
			return
		}
		c.JSON(http.StatusOK, allOrderItem)
	}
}
func GetOrderItemsBYOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderId := c.Param("order_id")

		allOrderItems, err := ItemByOrder(orderId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "no order found with respect to this order id"})
			return
		}
		c.JSON(http.StatusOK, allOrderItems)

	}
}
func ItemByOrder(id string) (OrderItems []primitive.M, err error) {
var ctx, cancel= context.WithTimeout(context.Background(), 100*time.Second)

matchStage:=bson.D{{"$match", bson.D{{"order)id", id}}}}

lookupStage :=bson.D{{"$lookup",bson.D{{"from","food"},{"localField","food_id"},{"foreignField","food_id"},{"as","food"}}}}

unwindStage:=bson.D{{"$unwind",bson.D{{"path","$food"},{"preserveNUllAndEmptyArrays",true}}}}

lookupOderStage:=bson.D{{"$lookup", bson.D{{"food", "order"},{"localField","order_id"},{"foreignField","order"},{"as","order"}}}}
unwindOrderStage:=bson.D{{"$unwind",bson.D{{"path","$order"},{"preserveNUllAndEmptyArrays",true}}}}

}
func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		orderItemId := c.Param("order_item_id")

		var orderItem models.OrderItem

		err := orderItemCollection.FindOne(ctx, bson.M{"orderItem_id": orderItemId}).Decode(&orderItem)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": " error occured while listing order item"})
			return
		}
		c.JSON(http.StatusOK, orderItem)

	}
}
func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var orderItemPack OrderItemPack
		var order models.Order

		if err := c.BindJSON(&orderItemPack); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		}
		order.Order_date, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		orderItmesToBeInserted := []interface{}{}
		order.Table_id = orderItemPack.Table_id
		order_id := orderItemOrderCreator(order)

		for _, orderItem := range orderItemPack.Order_items {
			orderItem.Order_id = order.Order_id
			validationErr := validate.Struct(orderItem)
			if validationErr != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
				return
			}

			orderItem.ID = primitive.NewObjectID()
			orderItem.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			orderItem.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			orderItem.Order_item_id = orderItem.ID.Hex()
			var num = toFixed(*orderItem.Unit_price, 2)
			orderItem.Unit_price = &num
			orderItmesToBeInserted = append(orderItmesToBeInserted, orderItem)

		}
		insertedOrderItem, err := orderItemCollection.InsertMany(ctx, orderItmesToBeInserted)

		if err != nil {
			log.Fatal(err)

		}
		defer cancel()
		c.JSON(http.StatusOK, insertedOrderItem)
	}
}
func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var orderItem models.OrderItem
		orderItemId := c.Param("order_item_id")
		filter := bson.M{"order_item_id": orderItemId}
		var updateObj primitive.D
		if orderItem.Unit_price != nil {
			updateObj = append(updateObj, bson.E{"unit_price", *&orderItem.Unit_price})
		}
		if len(orderItem.Quantity) == 0 {
			updateObj = append(updateObj, bson.E{"quantity", *&orderItem.Quantity})

		}
		if orderItem.Food_id != nil {
			updateObj = append(updateObj, bson.E{"food_id", *&orderItem.Food_id})
		}
		orderItem.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"update_at", orderItem.Updated_at})

		upsert := true

		opt := options.UpdateOptions{
			Upsert: &upsert,
		}
		result, err := orderItemCollection.UpdateOne(
			ctx, filter, bson.D{{"$set", updateObj}}, &opt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"erroe": "Order is not updated"})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}
