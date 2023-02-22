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
	deletionTask := func(path string) error {
		if strings.Contains(path, image.GetPath()) {
			return os.Remove(image.GetPath())
		}
		return nil
	}

	return filepath.Walk(r.parseImageDirDoing(deletionTask))
}

func (r fsrepo) Load() ([]string, error) {
	images := []string{}
	loadTask := func(path string) error {
		images = append(images, path)
		return nil
	}

	err := filepath.Walk(r.parseImageDirDoing(loadTask))
	return images, err
}

func (r fsrepo) ConvertToWebp(image domain.Image) error {
	conversionTask := func(path string) error {
		if strings.Contains(path, image.GetPath()) {
			return runConversionCommand(image)
		}
		return nil
	}

	return filepath.Walk(r.parseImageDirDoing(conversionTask))
}

func (r fsrepo) parseImageDirDoing(toDoTask func(string) error) (string, filepath.WalkFunc) {
	return r.images_dir, func(path string, file os.FileInfo, error error) error {
		if error != nil {
			return error
		}
		if file.IsDir() {
			return r.skipIfExcluded(file)
		}

		return toDoTask(path)
	}
}

func (r fsrepo) skipIfExcluded(file os.FileInfo) error {
	for _, excluded_dir := range r.excluded_dirs {
		if strings.Contains(excluded_dir, file.Name()) {
			return filepath.SkipDir
		}
	}
	return nil
}
