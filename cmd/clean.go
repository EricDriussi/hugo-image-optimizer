package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cleanCmd)
}

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove unused images",
	Long:  "Removes images not being referenced in posts",
	Run: func(cmd *cobra.Command, args []string) {
		RmUnusedImages()
	},
}

func RmUnusedImages() {
	postService := buildPostService()
	imageService := buildImageService()

	imageReferences, err := postService.AllReferencedImagePaths()
	if err != nil {
		log.Fatal(err)
	}
	if err := imageService.RemoveAllExcept(imageReferences); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Unused images have been removed")
}
