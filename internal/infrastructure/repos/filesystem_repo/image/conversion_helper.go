package filesystemrepo

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/EricDriussi/hugo-image-optimizer/internal/domain"
)

func runConversionCommand(image domain.Image) error {
	if image.IsGif() {
		return gif2webp(image)
	} else {
		return img2webp(image)
	}
}

func gif2webp(image domain.Image) error {
	cmd := gif2webpCommandBuilder(image)
	if err := cmd.Run(); err != nil {
		return errors.New(fmt.Sprintf("Couldn't convert gif: %s\n", image.GetPath()))
	}
	return os.Remove(image.GetPath())
}

func gif2webpCommandBuilder(image domain.Image) *exec.Cmd {
	filepath_without_ext := strings.TrimSuffix(image.GetPath(), image.GetExtension())
	webp_filepath := fmt.Sprintf("%s.webp", filepath_without_ext)

	cmd_params := []string{"-q", "50", "-mixed", image.GetPath(), "-o", webp_filepath}
	return exec.Command("gif2webp", cmd_params...)
}

func img2webp(image domain.Image) error {
	cmd := cwebpCommandBuilder(image)
	if err := cmd.Run(); err != nil {
		return errors.New(fmt.Sprintf("Couldn't convert image: %s\n", image.GetPath()))
	}
	return os.Remove(image.GetPath())
}

func cwebpCommandBuilder(image domain.Image) *exec.Cmd {
	filepath_without_ext := strings.TrimSuffix(image.GetPath(), image.GetExtension())
	webp_filepath := fmt.Sprintf("%s.webp", filepath_without_ext)

	cmd_params := []string{"-q", "50", image.GetPath(), "-o", webp_filepath}
	return exec.Command("cwebp", cmd_params...)
}
