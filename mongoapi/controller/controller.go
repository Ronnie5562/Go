package controller

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://rabimbola:yN1VqwUgdJQph8sj@cluster0.9qglrgw.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

// Most Important
var colection *mongo.Collection

// connect with mongoDB

func init() {
	//Client option
	clientOption := options.Client().ApplyURI(connectionString)
	mongo.Connect(context.TODO(), clientOption)
}
