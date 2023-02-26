package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(convertCmd)
}

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert images to webp",
	Long:  "Converts all images (jpg, png, gif) to webp",
	Run: func(cmd *cobra.Command, args []string) {
		ConvertToWebp()
	},
}

func ConvertToWebp() {
	imageService := buildImageService()

	if err := imageService.Convert(); err != nil {
		log.Fatal(err)
	}
}
