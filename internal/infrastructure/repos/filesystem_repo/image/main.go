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
	err := filepath.Walk(r.images_dir, r.loadImagesInto(&images))
	return images, err
}

func (r fsrepo) loadImagesInto(images_list *[]string) filepath.WalkFunc {
	return func(path string, file os.FileInfo, error error) error {
		if error != nil {
			return error
		}

		if file.IsDir() {
			if contains(r.excluded_dirs, file.Name()) {
				return filepath.SkipDir
			}
			return nil
		}

		*images_list = append(*images_list, file.Name())
		return nil
	}
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if strings.Contains(s, str) {
			return true
		}
	}
	return false
}