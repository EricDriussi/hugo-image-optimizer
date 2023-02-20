package cmd

import (
	"fmt"
	"log"

	filterService "github.com/EricDriussi/hugo-image-optimizer/internal/filter_service"
	imageService "github.com/EricDriussi/hugo-image-optimizer/internal/image_service"
	filesystemrepo "github.com/EricDriussi/hugo-image-optimizer/internal/infrastructure/repos/filesystem_repo/post"
	postReader "github.com/EricDriussi/hugo-image-optimizer/internal/post_reader_service"
	services "github.com/EricDriussi/hugo-image-optimizer/internal/services/post"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	posts_path := viper.GetString("dirs.posts")
	image_files := getAllImages(posts_path)
	unused_images := filterService.Unused_image_paths(image_files, &image_references)
	imageService.RM_images(unused_images)
	fmt.Println("Done cleaning!")
}

func getAllImages(posts_path string) []string {
	postRepo := filesystemrepo.NewPost(posts_path)
	postService := services.NewPost(postRepo)
	all_image_references, err := postService.GetImagesInPosts()
	if err != nil {
		log.Fatal("Something went wrong: ", err)
	}
	return all_image_references
}
