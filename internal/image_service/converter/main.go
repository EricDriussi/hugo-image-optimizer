package converter

// TODO.Native conversion

import (
	"fmt"
	"os/exec"
	"strings"
)

func Gif(filepath string) error {
	trimmed_filepath := strings.TrimSuffix(filepath, ".gif")
	output_filepath := fmt.Sprintf("%s.webp", trimmed_filepath)
	cmd := exec.Command("gif2webp", "-q", "50", "-mixed", filepath, "-o", output_filepath)

	return cmd.Run()
}

func Png(filepath string) error {
	cmd := generic_convert_command(filepath, ".png")
	return cmd.Run()
}

func Jpg(filepath string) error {
	cmd := generic_convert_command(filepath, ".jpg")
	return cmd.Run()
}

func Jpeg(filepath string) error {
	cmd := generic_convert_command(filepath, ".jpeg")
	return cmd.Run()
}

func generic_convert_command(filepath string, ext string) *exec.Cmd {
	trimmed_filepath := strings.TrimSuffix(filepath, ext)
	output_filepath := fmt.Sprintf("%s.webp", trimmed_filepath)
	return exec.Command("cwebp", "-q", "50", filepath, "-o", output_filepath)
}
