package postReader

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func All_posts_as_bytes() []byte {
	var (
		working_dir = viper.GetString("dirs.project")
		posts_dir   = viper.GetString("dirs.posts")
	)

	posts_stack := []byte{}
	posts_path := fmt.Sprintf("%s%s", working_dir, posts_dir)
	filepath.Walk(posts_path, func(filepath string, file os.FileInfo, error error) error {
		if file.IsDir() {
			return nil
		}
		single_post, _ := os.ReadFile(filepath)
		posts_stack = append(posts_stack, single_post...)
		return nil
	})

	return append([]byte{}, posts_stack...)
}
