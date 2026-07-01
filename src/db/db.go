package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	// Getting environment variables
	mongodbURI := os.Getenv("MONGODB_URI")
	dbName := os.Getenv("DB_NAME")
	
	if mongodbURI == "" || dbName == "" {
		log.Fatal("Error loading database environment variables")
	}

	// Creating context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 24 * time.Hour)
	defer cancel()

	// Creating client options
	clientOptions := options.Client().ApplyURI(mongodbURI)

	// Connecting to mongodb
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Connection to mongodb failed: ", err)
	}

	// Checking the connection with a ping
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("MongoDB ping failed: ", err)
	}

	fmt.Println("Connected to MongoDB")

	// Getting the database
	DB = client.Database(dbName)
}

// function to get Collection
func GetCollection(coll string) *mongo.Collection {
	return DB.Collection(coll)
}