package cmd

import (
	filterService "hugo-images/internal/filter_service"
	imageService "hugo-images/internal/image_service"
	postReader "hugo-images/internal/post_reader_service"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cleanCmd)
}

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "rm unused images",
	Long:  "Removes images not being referenced in posts",
	Run: func(cmd *cobra.Command, args []string) {
		all_posts := postReader.All_posts_as_bytes()
		Rm_unused_images(all_posts)
	},
}

func Rm_unused_images(all_posts []byte) {
	image_files := imageService.ImagesInIncludedDirs()
	image_references := filterService.MD_images_present_in(&all_posts)
	unused_images := filterService.Unused_images(image_files, &image_references)
	imageService.RM_images(unused_images)
}
