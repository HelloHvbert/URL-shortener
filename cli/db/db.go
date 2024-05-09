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

// Holds urls collection in the database
var collection *mongo.Collection


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
		// added
			fmt.Println("Error connecting to database")
			os.Exit(-1)
	}
	// Connection works
	// fmt.Println("Connected to MongoDB!")

	collection = client.Database("url-shortener").Collection("urls")
}

// Get the original URL from the shortened URL
func GetOriginalUrl(shortUrl string) (string, error) {
	var result bson.M
	// Find the shortened URL in the database
	err := collection.FindOne(context.Background(), bson.M{"shortUrl": shortUrl}).Decode(&result)
	if err != nil {
			return "", err
	}

	// If URL is not found
	if result == nil {
		return "",  errors.New("short URL not found")
	}

	// Return the original URL
	return result["url"].(string), nil
}

// Check if url already exists
func CheckIfUrlExists(url string) (string, error) {
	var result bson.M
	// Find the shortened URL in the database
	err := collection.FindOne(context.Background(), bson.M{"url": url}).Decode(&result)
	if err != nil {
			return "", err
	}
	// If shortened URL is not found
	if result == nil {
		return "",  errors.New("URL not found")
	}

	// Return the shortened URL
	return result["shortUrl"].(string), nil
}

// DisplayAllUrls: Display all the URLs in the database
func DisplayAllUrls() ([]bson.M, error) {
	// Find all the URLs in the database
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
			fmt.Println(err)
			return nil, err
	}
	var results []bson.M
	if err = cursor.All(context.Background(), &results); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return results, nil
}
// InsertOne: Insert a new URL into the database
func InsertOne(originalUrl string, shortUrl string) error{
	_, err := collection.InsertOne(context.Background(), bson.M{"url": originalUrl, "shortUrl": shortUrl})
	
	return err
}

// DeleteOne: Delete a URL from the database
func DeleteOne(shortUrl string) error {
	_, err := collection.DeleteOne(context.Background(), bson.M{"shortUrl": shortUrl})
	return err
}

// UpdateOne: Update a URL in the database
func UpdateOne(shortUrl string, data bson.M) error {
	_, err := collection.UpdateOne(context.Background(), bson.M{"shortUrl": shortUrl}, data)
	return err
}