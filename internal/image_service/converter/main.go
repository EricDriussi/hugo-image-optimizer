package converter

// TODO.Native conversion?
// TODO.Specify dependencies [gif2webp, cwebp]

import (
	"fmt"
	"os/exec"
	"strings"
)

func Gif(filepath string) error {
	trimmed_filepath := strings.TrimSuffix(filepath, ".gif")
	command := fmt.Sprintf("gif2webp -q 50 -mixed %s -o %s.webp", filepath, trimmed_filepath)
	cmd := exec.Command(command)

	return cmd.Run()
}

func Png(filepath string) error {
	trimmed_filepath := strings.TrimSuffix(filepath, ".png")
	command := fmt.Sprintf("cwebp -q 50 %s.png -o %s.webp", filepath, trimmed_filepath)
	cmd := exec.Command(command)

	return cmd.Run()
}

func Jpg(filepath string) error {
	trimmed_filepath := strings.TrimSuffix(filepath, ".jpg")
	command := fmt.Sprintf("cwebp -q 50 %s.jpg -o %s.webp", filepath, trimmed_filepath)
	cmd := exec.Command(command)

	return cmd.Run()
}

func Jpeg(filepath string) error {
	trimmed_filepath := strings.TrimSuffix(filepath, ".jpeg")
	command := fmt.Sprintf("cwebp -q 50 %s.jpeg -o %s.webp", filepath, trimmed_filepath)
	cmd := exec.Command(command)

	return cmd.Run()
}
