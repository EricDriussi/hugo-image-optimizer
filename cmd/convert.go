package cmd

import (
	imageService "hugo-images/internal/image_service"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(convertCmd)
}

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "convert images to webp",
	Long:  "Converts all images (jpg, png, gif) to webp",
	Run: func(cmd *cobra.Command, args []string) {
		image_files := imageService.ImagesInIncludedDirs()
		Convert_to_webp(image_files)
	},
}

func Convert_to_webp(images []string) {
	imageService.Convert_images(images)
}
