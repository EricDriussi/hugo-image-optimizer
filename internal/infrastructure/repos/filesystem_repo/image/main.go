package filesystemrepo

import (
	"os"
	"path/filepath"
	"strings"
)

type fsrepo struct {
	images_dir    string
	excluded_dirs []string
}

func NewImage(images_dir string, excluded []string) fsrepo {
	return fsrepo{
		images_dir:    images_dir,
		excluded_dirs: excluded,
	}
}

func (r fsrepo) Load() ([]string, error) {
	images := []string{}
	err := filepath.Walk(r.images_dir, loadImagesExcluding(&images, r.excluded_dirs))
	return images, err
}

func loadImagesExcluding(images_list *[]string, excluded []string) filepath.WalkFunc {
	return func(filepath string, file os.FileInfo, error error) error {
		if error != nil {
			return error
		}
		if file.IsDir() || filepath_is_excluded(filepath, excluded) {
			return nil
		}
		*images_list = append(*images_list, file.Name())
		return nil
	}
}

func filepath_is_excluded(path string, excluded []string) bool {
	splittedPath := strings.Split(path, "/")
	for i, dir_name := range splittedPath {
		for _, excluded := range excluded {
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
