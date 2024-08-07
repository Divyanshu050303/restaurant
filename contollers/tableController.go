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

var tableCollection *mongo.Collection = datebase.OpenCollection(datebase.Client, "table")

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		result, err := orderCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error while listing table item"})
			return
		}
		var allTabels []bson.M
		if err = result.All(ctx, &allTabels); err != nil {
			log.Fatal(err)

		}
		c.JSON(http.StatusOK, allTabels)

	}
}
func GetTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		tableId := c.Param("table_id")

		var table models.Table

		err := tableCollection.FindOne(ctx, bson.M{"table_id": tableId}).Decode(&table)

		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetichng the table item"})
			return
		}
		c.JSON(http.StatusOK, table)

	}
}
func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var table models.Table
		if err := c.BindJSON(&table); err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(table)

		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		table.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		table.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		table.ID = primitive.NewObjectID()

		table.Table_id = table.ID.Hex()

		result, insertErr := tableCollection.InsertOne(ctx, table)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Table item is not created"})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)

	}
}
func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var table models.Table

		tableId := c.Param("table_id")
		if err := c.BindJSON(&table); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var updateObj primitive.D

		if table.Number_of_guest != nil {
			updateObj = append(updateObj, bson.E{"number_of_guests", table.Number_of_guest})

		}

		if table.Table_number != nil {
			updateObj = append(updateObj, bson.E{"table_number", table.Table_number})

		}

		table.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		upsert := true
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}
		filter := bson.M{"table_id": tableId}

		result, err := tableCollection.UpdateOne(ctx, filter, bson.D{{"$set", updateObj}}, &opt)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Table item is not updated"})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)

	}
}
