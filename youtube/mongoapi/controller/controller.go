package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"
)

const connectionString = "mongodb+srv://rabimbola:yN1VqwUgdJQph8sj@cluster0.9qglrgw.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

// Most Important
var collection *mongo.Collection

// connect with mongoDB

func init() {
	//Client option
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance created")
}
