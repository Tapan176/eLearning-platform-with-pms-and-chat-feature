package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

// InitMongoDB initializes the MongoDB connection
func InitMongoDB() *mongo.Database {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017") // Replace with your MongoDB connection string
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	db = client.Database("myelearningdb") // Replace with your database name

	return db
}

// GetDB returns the MongoDB database instance
func GetDB() *mongo.Database {
	return db
}
