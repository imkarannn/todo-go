package db

import (
    "context"
    "log"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Database and collection names
const (
    dbName         = "todo"
    collectionName = "userMaster"
)

// Function to establish connection with MongoDB
func init() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    var err error
    client, err = mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }
}

// GetCollection returns the MongoDB collection
func GetCollection() *mongo.Collection {
    return client.Database(dbName).Collection(collectionName)
}
