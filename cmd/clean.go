package cmd

import (
	"fmt"
	filterService "github.com/EricDriussi/hugo-image-optimizer/internal/filter_service"
	imageService "github.com/EricDriussi/hugo-image-optimizer/internal/image_service"
	postReader "github.com/EricDriussi/hugo-image-optimizer/internal/post_reader_service"

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
		Rm_unused_images()
	},
}

func Rm_unused_images() {
	fmt.Println("Removing unused images")
	all_posts := postReader.All_posts_as_bytes()
	image_references := filterService.Images_present_in(&all_posts)
	image_files := imageService.ImagesInIncludedDirs()
	unused_images := filterService.Unused_image_paths(image_files, &image_references)
	imageService.RM_images(unused_images)
	fmt.Println("Done cleaning!")
}
