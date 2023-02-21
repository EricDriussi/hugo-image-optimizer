package cmd

import (
	"fmt"
	"log"

	imRepo "github.com/EricDriussi/hugo-image-optimizer/internal/infrastructure/repos/filesystem_repo/image"
	postRepo "github.com/EricDriussi/hugo-image-optimizer/internal/infrastructure/repos/filesystem_repo/post"
	imSrv "github.com/EricDriussi/hugo-image-optimizer/internal/services/image"
	postSrv "github.com/EricDriussi/hugo-image-optimizer/internal/services/post"

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
	imageReferences := getReferences()
	imageService := getImageService()

	if err := imageService.RemoveAllExcept(imageReferences); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Unused images have been removed")
}

func getReferences() []string {
	posts_path := viper.GetString("dirs.posts")
	postRepo := postRepo.NewPost(posts_path)
	postService := postSrv.NewPost(postRepo)
	imageReferences, err := postService.GetAllReferencedImagePaths()
	if err != nil {
		log.Fatal(err)
	}
	return imageReferences
}

func getImageService() imSrv.ImageService {
	images_path := viper.GetString("dirs.images")
	excluded_images_path := viper.GetStringSlice("dirs.images_exclude")
	imageRepo := imRepo.NewImage(images_path, excluded_images_path)
	return imSrv.NewImage(imageRepo)
}
