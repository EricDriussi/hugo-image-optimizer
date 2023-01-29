package imageService

import (
	"optimize/internal/image_service/converter"
	"optimize/internal/util"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func Convert_images(list []string) error {
	images_path := viper.GetString("dirs.images")

	return filepath.Walk(images_path, func(filepath string, file os.FileInfo, error error) error {
		if file.IsDir() || filepath_is_excluded(filepath) {
			return nil
		}
		is_gif := strings.HasSuffix(file.Name(), ".gif")
		is_png := strings.HasSuffix(file.Name(), ".png")
		is_jpg := strings.HasSuffix(file.Name(), ".jpg")
		is_jpeg := strings.HasSuffix(file.Name(), ".jgep")

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
		if is_jpg {
			err := converter.Jpg(filepath)
			os.Remove(filepath)
			return err
		}
		if is_jpeg {
			err := converter.Jpeg(filepath)
			os.Remove(filepath)
			return err
		}
		return nil
	})
}

func RM_images(list []string) error {
	images_path := viper.GetString("dirs.images")

	return filepath.Walk(images_path, func(filepath string, file os.FileInfo, error error) error {
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
	images_path := viper.GetString("dirs.images")
	image_list := []string{}

	read_images(images_path, &image_list)
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
