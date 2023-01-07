package posts

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func List() []string {
	var (
		working_dir = viper.GetString("dirs.project")
		posts_dir   = viper.GetString("dirs.posts")
	)

	posts_list := []string{}
	path := fmt.Sprintf("%s%s", working_dir, posts_dir)
	error := list_posts(path, &posts_list)

	if error != nil {
		fmt.Println(error)
	}
	return posts_list
}

func list_posts(path string, list *[]string) error {
	return filepath.Walk(path, func(filepath string, info os.FileInfo, error error) error {
		if info.IsDir() {
			return nil
		}

		*list = append(*list, info.Name())
		return nil
	})
}

func Read(filename string) (string, error) {
	post, err := os.ReadFile(fullPathTo(filename))
	if err != nil {
		return "", err
	}
	return string(post), nil
}

func fullPathTo(filename string) string {
	var (
		working_dir = viper.GetString("dirs.project")
		posts_dir   = viper.GetString("dirs.posts")
	)
	return fmt.Sprintf("%s%s%s", working_dir, posts_dir, filename)
}
