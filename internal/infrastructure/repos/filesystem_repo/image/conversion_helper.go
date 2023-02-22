package filesystemrepo

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/EricDriussi/hugo-image-optimizer/internal/domain"
)

func runConversionCommand(image domain.Image) error {
	if image.IsGif() {
		cmd := gif2webpCommandBuilder(image)
		if err := cmd.Run(); err != nil {
			return errors.New(fmt.Sprintf("Couldn't convert gif: %s\n", image.GetPath()))
		}
	} else {
		cmd := cwebpCommandBuilder(image)
		if err := cmd.Run(); err != nil {
			return errors.New(fmt.Sprintf("Couldn't convert image: %s\n", image.GetPath()))
		}
	}
	return nil
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
