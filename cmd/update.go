package cmd

import (
	"fmt"
	postReader "optimize/internal/post_reader_service"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(referenceCmd)
}

var referenceCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates all image references",
	Long:  "Updates all image references in MD to .webp extension",
	Run: func(cmd *cobra.Command, args []string) {
		Update_References()
	},
}

func Update_References() {
	fmt.Println("Updating image references in posts")
	postReader.Update_image_references()
	fmt.Println("Done updating!")
}
