package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(referenceCmd)
}

var referenceCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates all image references",
	Long:  "Updates all image references in MD to .webp extension",
	Run: func(cmd *cobra.Command, args []string) {
		UpdateReferences()
	},
}

func UpdateReferences() {
	fmt.Println("Updating image references in posts")
	postService := buildPostService()
	if err := postService.UpdateAllImageReferences(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done updating!")
}
