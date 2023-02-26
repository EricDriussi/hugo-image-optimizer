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

func buildPostService() postSrv.PostService {
	postsPath := viper.GetString("dirs.posts")
	postRepo := postRepo.NewPost(postsPath)
	return postSrv.NewPost(postRepo)
}

func buildImageService() imSrv.ImageService {
	imagesPath := viper.GetString("dirs.images")
	excludedImagesPath := viper.GetStringSlice("dirs.images_exclude")
	imageRepo := imRepo.NewImage(imagesPath, excludedImagesPath)
	return imSrv.NewImage(imageRepo)
}
