package cmd

import (
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
	postService := buildPostService()
	if err := postService.UpdateAllImageReferences(); err != nil {
		log.Fatal(err)
	}
}
