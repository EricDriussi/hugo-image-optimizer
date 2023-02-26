package cmd

import (
	"log"
	"os"

	"github.com/EricDriussi/hugo-image-optimizer/internal/config"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates all image references",
	Long:  "Updates all image references in MD to .webp",
	Run: func(cmd *cobra.Command, args []string) {
		os.Chdir(websitePath)
		config.Load(cfgFile)
		UpdateReferences()
	},
}

func UpdateReferences() {
	postService := buildPostService()
	if err := postService.UpdateAllImageReferences(); err != nil {
		log.Fatal(err)
	}
}
