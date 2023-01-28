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
	Short: "updates all image references",
	Long:  "Updates all image references in MD for the .webp extension",
	Run: func(cmd *cobra.Command, args []string) {
		all_posts := postReader.All_posts_as_bytes()
		Update_References(all_posts)
	},
}

func Update_References(all_posts []byte) {
}
