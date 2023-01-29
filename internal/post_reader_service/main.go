package postReader

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func All_posts_as_bytes() []byte {
	posts_path := viper.GetString("dirs.posts")

	posts_stack := []byte{}
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

func Update_image_references() {
	posts_path := viper.GetString("dirs.posts")

	err := filepath.Walk(posts_path, change_extensions)
	if err != nil {
		panic(err)
	}
}

func change_extensions(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, replace(content), 0)
	if err != nil {
		return err
	}
	return nil
}

func replace(content []byte) []byte {
	gif := strings.ReplaceAll(string(content), ".gif)", ".webp)")
	png := strings.ReplaceAll(string(gif), ".png)", ".webp)")
	jpg := strings.ReplaceAll(string(png), ".jpg)", ".webp)")
	jpeg := strings.ReplaceAll(string(jpg), ".jpeg)", ".webp)")
	return []byte(jpeg)
}
