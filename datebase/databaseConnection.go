package datebase

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	MongoDb := "mongodb+srv://divya2003a:admin@cluster0.5tn37xn.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	// fmt.Println(MongoDb)
	// client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancle()
	// err = client.Connect(ctx)
	// if err != nil {
	// 	log.Fatal(err)

	// }
	clientOption := options.Client().ApplyURI(MongoDb)

	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")
	fmt.Println("Connected to MongoDB")
	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant").Collection(collectionName)
	return collection
}
