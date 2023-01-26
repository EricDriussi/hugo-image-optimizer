package postReader

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

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

func Filter_md_images(text *[]byte) []byte {
	var result bytes.Buffer
	r, _ := regexp.Compile("!\\[.*\\]\\(.*\\)")
	matches := r.FindAll(*text, -1)
	for _, match := range matches {
		result.Write(match)
	}
	return result.Bytes()
}

func bytesAreEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
