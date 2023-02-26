package filesystemrepo

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/EricDriussi/hugo-image-optimizer/internal/domain"
	"golang.org/x/sync/errgroup"
)

type fsrepo struct {
	imagesDir    string
	excludedDirs []string
}

var g errgroup.Group

func NewImage(imagesDir string, excluded []string) fsrepo {
	return fsrepo{
		imagesDir:    imagesDir,
		excludedDirs: excluded,
	}
}

func (r fsrepo) Delete(image domain.Image) error {
	deletionTask := func(path string) error {
		if strings.Contains(path, image.Path()) {
			return os.Remove(image.Path())
		}
		return nil
	}

	return filepath.Walk(r.parseImagesDirDoing(deletionTask))
}

func (r fsrepo) Load() ([]string, error) {
	images := []string{}
	loadTask := func(path string) error {
		images = append(images, path)
		return nil
	}

	err := filepath.Walk(r.parseImagesDirDoing(loadTask))
	return images, err
}

func (r fsrepo) ConvertToWebp(images []domain.Image) error {
	conversionTask := func(path string) error {
		for _, image := range images {
			if strings.Contains(path, image.Path()) {
				imageCopy := image
				g.Go(func() error {
					return runConversionCommand(imageCopy)
				})
			}
		}
		return nil
	}

	filepath.Walk(r.parseImagesDirDoing(conversionTask))

	if rmErr := g.Wait(); rmErr != nil {
		return errors.New("Some images were not converted :(")
	}
	return nil
}

func (r fsrepo) parseImagesDirDoing(toDoTask func(string) error) (string, filepath.WalkFunc) {
	return r.imagesDir, func(path string, file os.FileInfo, error error) error {
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
	for _, excludedDir := range r.excludedDirs {
		if strings.Contains(excludedDir, file.Name()) {
			return filepath.SkipDir
		}
	}
	return nil
}
