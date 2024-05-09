/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"example.com/cli/db"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete one or more shortened URL",
	Long: `Delete one or more a shortened URL.
	You can use "list" command before deleting to see 
	all the URLs and their shortened forms.
	Usage:
		url-shortener delete [short-url]
		url-shortener delete [short-url1] [short-url2] ... [short-urlN]
	Example:
		url-shortener delete XYZ123`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		// Delete the URL
		for _, shortUrl := range args{
			fmt.Printf("Deleting %s...\n", shortUrl)
			// Check if db.DeleteOne returns an error
			err := db.DeleteOne(shortUrl)
			if err != nil {
				fmt.Println("Error while deleting the URL:", err)
				return
			}
			fmt.Printf("Url %s was deleted from the database.\n", shortUrl)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
