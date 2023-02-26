package cmd

import (
	"fmt"
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
	fmt.Println("Converting all images to .webp")
	fmt.Println(".gif and large .png might take a while...")
	imageService := buildImageService()

	if err := imageService.Convert(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done converting!")
}
