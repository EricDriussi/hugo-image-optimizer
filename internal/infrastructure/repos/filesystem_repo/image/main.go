package filesystemrepo

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
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

func (r fsrepo) ConvertToWebp(image domain.Image) error {
	if image.IsGif() {
		cmd := gif2webpCommandBuilder(image)
		if err := cmd.Run(); err != nil {
			return errors.New(fmt.Sprintf("Couldn't convert gif: %s\n", image.GetPath()))
		}
		return nil
	} else {
		cmd := cwebpCommandBuilder(image)
		if err := cmd.Run(); err != nil {
			return errors.New(fmt.Sprintf("Couldn't convert image: %s\n", image.GetPath()))
		}
		return nil
	}
}

func gif2webpCommandBuilder(image domain.Image) *exec.Cmd {
	filepath_without_ext := strings.TrimSuffix(image.GetPath(), image.GetExtension())
	webp_filepath := fmt.Sprintf("%s.webp", filepath_without_ext)

	cmd_params := []string{"-q", "50", "-mixed", image.GetPath(), "-o", webp_filepath}
	return exec.Command("gif2webp", cmd_params...)
}

func cwebpCommandBuilder(image domain.Image) *exec.Cmd {
	filepath_without_ext := strings.TrimSuffix(image.GetPath(), image.GetExtension())
	webp_filepath := fmt.Sprintf("%s.webp", filepath_without_ext)

	cmd_params := []string{"-q", "50", image.GetPath(), "-o", webp_filepath}
	return exec.Command("cwebp", cmd_params...)
}
