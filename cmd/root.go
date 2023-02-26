package cmd

import (
	"fmt"
	"os"

	"github.com/EricDriussi/hugo-image-optimizer/internal/config"
	imRepo "github.com/EricDriussi/hugo-image-optimizer/internal/infrastructure/repos/filesystem_repo/image"
	postRepo "github.com/EricDriussi/hugo-image-optimizer/internal/infrastructure/repos/filesystem_repo/post"
	imSrv "github.com/EricDriussi/hugo-image-optimizer/internal/services/image"
	postSrv "github.com/EricDriussi/hugo-image-optimizer/internal/services/post"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	version     bool
	cfgFile     string
	websitePath string
)

func init() {
	rootCmd.Flags().BoolVarP(&version, "version", "v", false, "Version number")

	rootCmd.Flags().StringVar(&cfgFile, "config", "", "Config file")
	cleanCmd.Flags().StringVar(&cfgFile, "config", "", "Config file")
	convertCmd.Flags().StringVar(&cfgFile, "config", "", "Config file")
	updateCmd.Flags().StringVar(&cfgFile, "config", "", "Config file")

	rootCmd.Flags().StringVar(&websitePath, "website-path", ".", "Website path")
	cleanCmd.Flags().StringVar(&websitePath, "website-path", ".", "Website path")
	convertCmd.Flags().StringVar(&websitePath, "website-path", ".", "Website path")
	updateCmd.Flags().StringVar(&websitePath, "website-path", ".", "Website path")
}

var rootCmd = &cobra.Command{
	Use:   "optimize",
	Short: "Simple Hugo image optimizer",
	Long: `
  Removes unused images.
  Converts all images to webp
  Updates all uses
`,
	Run: func(cmd *cobra.Command, args []string) {
		if version {
			fmt.Println("v1.0.0")
		} else {
			os.Chdir(websitePath)
			config.Load(cfgFile)
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
