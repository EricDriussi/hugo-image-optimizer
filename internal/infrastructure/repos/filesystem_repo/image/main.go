package filesystemrepo

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/EricDriussi/hugo-image-optimizer/internal/domain"
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

func (r fsrepo) Delete(image domain.Image) error {
	return filepath.Walk(r.images_dir, func(path string, file os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if file.IsDir() {
			if r.isInExcludedList(file.Name()) {
				return filepath.SkipDir
			}
			return nil
		}
		if strings.Contains(path, image.GetPath()) {
			return os.Remove(image.GetPath())
		}
		return nil
	})
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
			if r.isInExcludedList(file.Name()) {
				return filepath.SkipDir
			}
			return nil
		}

		*images_list = append(*images_list, path)
		return nil
	}
}

func (r fsrepo) isInExcludedList(dir string) bool {
	for _, excluded_dir := range r.excluded_dirs {
		if strings.Contains(excluded_dir, dir) {
			return true
		}
	}
	return false
}
