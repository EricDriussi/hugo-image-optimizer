package filesystemrepo

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/EricDriussi/hugo-image-optimizer/internal/domain"
	"github.com/spf13/viper"
)

var quality = fmt.Sprintf("%d", 100-viper.GetInt("compression.quality"))

func runConversionCommand(image domain.Image) error {
	if image.IsGif() {
		return gif2webp(image)
	} else {
		return img2webp(image)
	}
	// TODO: webp2webp(image) ??
}

func gif2webp(image domain.Image) error {
	cmd := gif2webpCommandBuilder(image)
	if err := cmd.Run(); err != nil {
		return errors.New(fmt.Sprintf("Couldn't convert gif: %s\n", image.Path()))
	}
	return os.Remove(image.Path())
}

func gif2webpCommandBuilder(image domain.Image) *exec.Cmd {
	filepathWithoutExt := strings.TrimSuffix(image.Path(), image.Extension())
	webpFilepath := fmt.Sprintf("%s.webp", filepathWithoutExt)

	cmdParams := []string{"-q", quality, "-mixed", image.Path(), "-o", webpFilepath}
	return exec.Command("gif2webp", cmdParams...)
}

func img2webp(image domain.Image) error {
	cmd := cwebpCommandBuilder(image)
	if err := cmd.Run(); err != nil {
		return errors.New(fmt.Sprintf("Couldn't convert image: %s\n", image.Path()))
	}
	// TODO: if original is already webp cwebp would overwrite it and rm would remove it
	// so no image would remain
	return os.Remove(image.Path())
}

func cwebpCommandBuilder(image domain.Image) *exec.Cmd {
	filepathWithoutExt := strings.TrimSuffix(image.Path(), image.Extension())
	webpFilepath := fmt.Sprintf("%s.webp", filepathWithoutExt)

	cmdParams := []string{"-q", quality, image.Path(), "-o", webpFilepath}
	return exec.Command("cwebp", cmdParams...)
}
