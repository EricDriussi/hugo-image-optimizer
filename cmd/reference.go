package cmd

import (
	postReader "hugo-images/internal/post_reader_service"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(referenceCmd)
}

var referenceCmd = &cobra.Command{
	Use:   "reference",
	Short: "Updates all image references",
	Long:  "Updates all image references in MD to .webp extension",
	Run: func(cmd *cobra.Command, args []string) {
		Update_References()
	},
}

func Update_References() {
	postReader.Update_image_references()
}
