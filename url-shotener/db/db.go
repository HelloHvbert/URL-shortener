package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

// DB:
// - Run: Connect to the database
// - GetOriginalUrl: Get the original URL from the short URL

func Run() {
	// Get value from .env
	err := godotenv.Load(".env")
	if err != nil {
			log.Fatalf("Error loading .env file: %s", err)
	}
	MONGO_URI := os.Getenv("DB_URI")

	// Connect to the database.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MONGO_URI))
	if err != nil {
			log.Fatal(err)
	}
	// Connection works
	fmt.Println("Connected to MongoDB!")

	collection = client.Database("url-shortener").Collection("urls")
}

// GetOriginalUrl: Get the original URL from the short URL
func GetOriginalUrl(shortUrl string) (string, error) {
	var result bson.M
	// Find the short URL in the database
	err := collection.FindOne(context.Background(), bson.M{"shortUrl": shortUrl}).Decode(&result)
	if err != nil {
			return "", err
	}

	if result == nil {
		return "",  errors.New("short URL not found")
	}

	return result["url"].(string), nil
}
