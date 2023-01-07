package images

import (
	"fmt"
	"hugo-images/internal/posts"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func ListUnusedIn(postsList []string) map[string]bool {
	images := ListAll()
	for image, present := range images {
		if present {
			break
		}
		for _, single_post := range postsList {
			post_content, _ := posts.Read(single_post)
			appears_in_post := strings.Contains(post_content, image)
			if appears_in_post {
				images[image] = true
			}
		}
	}
	return images
}

func ListAll() map[string]bool {
	var (
		working_dir = viper.GetString("dirs.project")
		images_dir  = viper.GetString("dirs.images")
	)

	image_list := map[string]bool{}
	path := fmt.Sprintf("%s%s", working_dir, images_dir)
	error := read_images(path, image_list)

	if error != nil {
		fmt.Println(error)
	}
	return image_list
}

func read_images(path string, list map[string]bool) error {
	return filepath.Walk(path, func(filepath string, info os.FileInfo, error error) error {
		if info.IsDir() || is_excluded_image_dir(filepath) {
			return nil
		}
		list[info.Name()] = false
		return nil
	})
}

func is_excluded_image_dir(path string) bool {
	splitted_path := strings.Split(path, "/")
	parent_dir := splitted_path[len(splitted_path)-2]
	return contains(viper.GetStringSlice("dirs.images_exclude"), parent_dir)
}

func contains(slice []string, elem string) bool {
	for _, value := range slice {
		if value == elem {
			return true
		}
	}
	return false
}