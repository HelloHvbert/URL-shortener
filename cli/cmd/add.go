package cmd

import (
	"bufio"
	"fmt"
	u "net/url"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	"example.com/cli/db"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var custom string
var random bool

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new url and prints the shorted url",
	Long: `Adds a new url and prints the shorted url.
	Usage:
		url-shortener add [url]
		url-shortener add [url] -c/--custom [custom short url]
		url-shortener add [url] -r/--random
	Example:
		url-shortener add https://example.com
		url-shortener add https://example.com -c custom-url
		url-shortener add https://example.com --random
	After adding a new url, you will get a short url that you can use to access the original url.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the url from the arguments
		url := args[0]

		// Check if the url already exists
		alreadyExists, _ := db.CheckIfUrlExists(url)
		if alreadyExists != ""{
			fmt.Printf("This URL already exists:\nhttp://localhost:9000/%s\n", u.PathEscape(alreadyExists))
			return
		}

		if startWithDash(url) {
			fmt.Println("URL cannot start with '-', '!'. Try again.")
			return
		}
		// Get the max length of the short URL
		length, _ := strconv.ParseInt(os.Getenv("SHORT_URL_MAX_LEN"), 10, 64)
		var shortUrl string 

		// no flags
		if !random  && custom == "" {
		var isCustom string
		reader := bufio.NewReader(os.Stdin)

		// Ask user if they want to use a custom short URL
		fmt.Println("Do you want to use a custom short URL? (y/Y):")
    isCustom, _ = reader.ReadString('\n')
		isCustom = isCustom[:1]

		if isCustom == "y" || isCustom == "Y" {
			// Get custom short URL from user
			fmt.Println("Enter a custom short URL:")
			shortUrl, _ = reader.ReadString('\n')

			err := godotenv.Load(".env")
			if err != nil {
					fmt.Println("Error loading .env file")
					return
			}
			// Get corrent length of short URL
			for int64(utf8.RuneCountInString(shortUrl)) > length + 1 || startWithDash(shortUrl){
				fmt.Println("Short URL should be less than", length, "characters and can't start with '-', '!'. Try again.")
				shortUrl, _ = reader.ReadString('\n')
			}
			// Remove the newline character
			shortUrl = shortUrl[:len(shortUrl)-1]
			// fmt.Println(shortUrl, utf8.RuneCountInString(shortUrl))
		} else {
			// or generate a random short URL
			shortUrl = db.GenerateShort()
			fmt.Println("Random short URL:", shortUrl)
		}
		// Insert the url into the database
		fmt.Println("Adding URL...")
		err := db.InsertOne(url, shortUrl)

		if err != nil {
			fmt.Println("Error inserting the URL:", err)
			return
		}

		// Print the short URL
		fmt.Printf("\nShort URL for %s:\n\nhttp://localhost:%s/%s\n", url, PORT,u.PathEscape(shortUrl))
		} else { // with using flags
			// Check flags

			// -c or --custom flag
			if custom != "" && !random {
				// Check if the custom short URL already exists
				alreadyExists, _ := db.GetOriginalUrl(custom)
				if alreadyExists != "" {
					fmt.Printf("This custom short URL is already used for %s\n", alreadyExists)
					return
				}

				// Check if the custom short URL has more characters than the max length
				if utf8.RuneCountInString(custom) > int(length) || startWithDash(custom) {
					fmt.Println("Short URL should be less than", length, "characters and cannot start with '-', '!'. Try again.")
					return
				}
				shortUrl = custom
			} else if custom == "" && random { // -r or --random flag
				shortUrl = db.GenerateShort()
				fmt.Println("Random short URL:", shortUrl)
			} else { // both flags, error
				fmt.Println("Please use 0 or 1 flag: -c/--custom or -r/--random")
				return
			}
			// Insert the url into the database
			fmt.Println("Adding URL...")
			err := db.InsertOne(url, shortUrl)

			if err != nil {
				fmt.Println("Error inserting the URL:", err)
				return
			}

			// Print the short URL
			fmt.Printf("\nShort URL for %s:\n\nhttp://localhost:%s/%s\n", url, PORT,u.PathEscape(shortUrl))
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&custom, "custom", "c", "", "custom short URL")
	addCmd.Flags().BoolVarP(&random, "random", "r", false, "create random short URL")

}

// parameters starting with a dash throw an error
func startWithDash(s string) bool {
	return strings.HasPrefix(s, "-") ||  strings.HasPrefix(s, "!")
}