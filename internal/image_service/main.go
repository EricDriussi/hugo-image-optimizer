package imageService

import (
	"fmt"
	"hugo-images/internal/image_service/converter"
	"hugo-images/internal/util"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func Convert_images(list []string) error {
	var (
		working_dir = viper.GetString("dirs.project")
		images_dir  = viper.GetString("dirs.images")
	)

	path := fmt.Sprintf("%s%s", working_dir, images_dir)

	return filepath.Walk(path, func(filepath string, file os.FileInfo, error error) error {
		if file.IsDir() || filepath_is_excluded(filepath) {
			return nil
		}
		is_gif := strings.HasSuffix(file.Name(), ".gif")
		is_png := strings.HasSuffix(file.Name(), ".png")
		is_jgp := strings.HasSuffix(file.Name(), ".jgp")

		if is_gif {
			err := converter.Gif(filepath)
			os.Remove(filepath)
			return err
		}
		if is_png {
			err := converter.Png(filepath)
			os.Remove(filepath)
			return err
		}
		if is_jgp {
			err := converter.Jpg(filepath)
			os.Remove(filepath)
			return err
		}
		return nil
	})
}

func RM_images(list []string) error {
	var (
		working_dir = viper.GetString("dirs.project")
		images_dir  = viper.GetString("dirs.images")
	)

	path := fmt.Sprintf("%s%s", working_dir, images_dir)

	return filepath.Walk(path, func(filepath string, file os.FileInfo, error error) error {
		if file.IsDir() || filepath_is_excluded(filepath) {
			return nil
		}
		file_is_not_in_deletion_list := !util.StringIsInArray(list, file.Name())
		if file_is_not_in_deletion_list {
			return nil
		}
		return os.Remove(filepath)
	})
}

func ImagesInIncludedDirs() []string {
	var (
		working_dir = viper.GetString("dirs.project")
		images_dir  = viper.GetString("dirs.images")
	)

	path := fmt.Sprintf("%s%s", working_dir, images_dir)
	image_list := []string{}

	read_images(path, &image_list)
	return image_list
}

func read_images(path string, list *[]string) error {
	return filepath.Walk(path, func(filepath string, file os.FileInfo, error error) error {
		if file.IsDir() || filepath_is_excluded(filepath) {
			return nil
		}
		*list = append(*list, file.Name())
		return nil
	})
}

func filepath_is_excluded(path string) bool {
	splittedPath := strings.Split(path, "/")
	for i, dir_name := range splittedPath {
		for _, excluded := range viper.GetStringSlice("dirs.images_exclude") {
			if dir_name == excluded {
				return true
			}
		}
		current_dir := strings.Join(splittedPath[:i+1], "/")
		// check if the current directory contains a ".hugo_build.lock" file or a ".git/" dir
		lock_file := current_dir + ".hugo_build.lock"
		if _, err := os.Stat(lock_file); !os.IsNotExist(err) {
			break
		}
		git_dir := current_dir + ".git/"
		if _, err := os.Stat(git_dir); !os.IsNotExist(err) {
			break
		}
	}
	return false
}
