package controller

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const connectionString = "mongodb+srv://ih0608899:ih060899@nodetuts.iujkb.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

// MOST IMPORTANT
var collection *mongo.Collection

// connect with mongoDB
func init() {
	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")
}
