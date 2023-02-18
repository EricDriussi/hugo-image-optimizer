package cmd

import (
	"fmt"
	imageService "github.com/EricDriussi/hugo-image-optimizer/internal/image_service"
	"os"

	"github.com/spf13/cobra"
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
			image_files := imageService.ImagesInIncludedDirs()
			Rm_unused_images()
			Convert_to_webp(image_files)
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
