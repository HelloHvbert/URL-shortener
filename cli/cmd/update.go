/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	u "net/url"
	"os"
	"strconv"
	"unicode/utf8"

	"example.com/cli/db"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
)

var newUrl string
var newShortUrl string

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update existing URL",
	Long: `Update existing URL by providing the short URL and the new URL.
	User need to provide at least one of the new URL or new short URL.
	Usage: 
		url-shortener update [current-short-url] -s [new-short-url] -u [new-url]
	Example:
		url-shortener update XYZ123 -s 123XYZ
		url-shortener update XYZ123 -u https://example.com/new-url
		url-shortener update XYZ123 -s 123XYZ -u https://example.com/new-url`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// check if new URL or new short URL is provided
		if newUrl == "" && newShortUrl == "" {
			fmt.Println("Please provide the new URL or new short URL using flag")
			return
		}

		// parameters validation
		if startWithDash(newShortUrl) || startWithDash(newUrl) {
			fmt.Println("Short URL and URL cannot start with a '-', '!'. Try again.")
			return
		}
		// add also
		oldShortUrl := args[0]

		// Check if the short URL exists
		isReal, _ := db.GetOriginalUrl(oldShortUrl)
		if isReal == "" {
			fmt.Printf("This short URL does not exist. Add it first using 'add' command.\n")
			return
		}

		var data bson.M 

		err := godotenv.Load(".env")
		if err != nil {
				fmt.Println("Error loading .env file")
				return
		}
		// Get the max length of the short URL
		length, _ := strconv.ParseInt(os.Getenv("SHORT_URL_MAX_LEN"), 10, 64)

		if newShortUrl != ""{
			if utf8.RuneCountInString(newShortUrl) > int(length) {
				fmt.Println("Short URL is too long. Max length is", length)
				return
			}
	
			// Check if the shortened URL already exists		
			if ifUrlExists, _ := db.GetOriginalUrl(newShortUrl); ifUrlExists != ""{
				fmt.Printf("http://localhost:%s/%s is already used for %s.\nPlease choose another one.\n", PORT,u.PathEscape(newShortUrl),ifUrlExists)
				return
			}
		}

		// Check if the URL already exists
		if ifShortUrlExists, _ := db.CheckIfUrlExists(newUrl); newUrl != "" {	
			if ifShortUrlExists != "" {
				fmt.Printf("This URL is already used for http://localhost:%s/%s.\nPlease choose another one.\n", PORT, u.PathEscape(ifShortUrlExists))
				return
			}
		}
		

		if newUrl != "" && newShortUrl != "" {
			data = bson.M{"$set": bson.M{"url": newUrl, "shortUrl": newShortUrl}}
		} else if newShortUrl != "" {
			data = bson.M{"$set": bson.M{"shortUrl": newShortUrl}}
		} else {
			data = bson.M{"$set": bson.M{"url": newUrl}}
		}
		fmt.Println("Updating URL...")

		// Update the URL
		err = db.UpdateOne(oldShortUrl, data)
		if err != nil {
			fmt.Println("Error updating URL:", err)
			return
		}
		// Print new data
		fmt.Println("URL updated successfully!")
		if newUrl != "" {
			fmt.Printf("New URL: %s\n", newUrl)
		}
		if newShortUrl != "" {
			fmt.Printf("New shortened URL: http://localhost:%s/%s\n", PORT, u.PathEscape(newShortUrl))
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&newUrl, "url", "u", "", "New URL")
	updateCmd.Flags().StringVarP(&newShortUrl, "short-url", "s", "", "New short URL")
}
