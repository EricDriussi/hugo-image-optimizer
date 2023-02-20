package cmd

import (
	"fmt"
	"log"
	"os"

	filesystemrepo "github.com/EricDriussi/hugo-image-optimizer/internal/infrastructure/repos/filesystem_repo/post"
	services "github.com/EricDriussi/hugo-image-optimizer/internal/services/post"

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
			posts_path := viper.GetString("dirs.posts")
			postRepo := filesystemrepo.NewPost(posts_path)
			postService := services.NewPost(postRepo)
			all_image_references, err := postService.GetImagesInPosts()
			if err != nil {
				log.Fatal("Something went wrong: ", err)
			}

			Rm_unused_images()
			Convert_to_webp(all_image_references)
			Update_References()
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
