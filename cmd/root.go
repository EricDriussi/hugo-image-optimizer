package cmd

import (
	"fmt"
	"os"

	imRepo "github.com/EricDriussi/hugo-image-optimizer/internal/infrastructure/repos/filesystem_repo/image"
	postRepo "github.com/EricDriussi/hugo-image-optimizer/internal/infrastructure/repos/filesystem_repo/post"
	imSrv "github.com/EricDriussi/hugo-image-optimizer/internal/services/image"
	postSrv "github.com/EricDriussi/hugo-image-optimizer/internal/services/post"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var version bool

func init() {
	rootCmd.Flags().BoolVarP(&version, "version", "v", false, "Version number")
}

var rootCmd = &cobra.Command{
	Use:   "optimize",
	Short: "Simple Hugo image optimizer",
	Long: `
  Removes unused images.
  Converts all images to webp
  Updates all uses
  WIP!
`,
	Run: func(cmd *cobra.Command, args []string) {
		if version {
			fmt.Println("v1.0.0")
		} else {
			RmUnusedImages()
			ConvertToWebp()
			UpdateReferences()
		}
	},
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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
