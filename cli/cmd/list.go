/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	u "net/url"

	"example.com/cli/db"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show all the urls you have created",
	Long: `List command shows all the urls you have shortened.
	Usage:
		url-shortener list
	Example output:
		- short-url1: original-url1: shortened-url1
		- short-url2: original-url2: shortened-url2`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		urls, err := db.DisplayAllUrls()
		if err != nil {
			fmt.Println("Error displaying all URLs: ", err)
			return
		}
		// Check if there are no URLs created
		if len(urls) == 0 {
			fmt.Println("No URLs found. Add a new URL using 'url-shortener add [url]' command.")
			return
		}
		// Print all the URLs
		fmt.Println("All URLs (short URL : original URL : shortened URL):")
		for _, url := range urls {
			fmt.Println("\t-", url["shortUrl"] ,":", url["url"], ":", fmt.Sprintf("http://localhost:%v/%s", PORT, u.PathEscape(url["shortUrl"].(string))))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
