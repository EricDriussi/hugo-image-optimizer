package converter

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func Gif(filepath string) {
	trimmed_filepath := strings.TrimSuffix(filepath, ".gif")
	output_filepath := fmt.Sprintf("%s.webp", trimmed_filepath)
	cmd := exec.Command("gif2webp", "-q", "50", "-mixed", filepath, "-o", output_filepath)

	if err := cmd.Run(); err != nil {
		log.Fatal(fmt.Sprintf("Couldn't convert GIF: %s\n", filepath), err)
	}
}

func Png(filepath string) {
	cmd := generic_convert_command(filepath, ".png")
	if err := cmd.Run(); err != nil {
		log.Fatal(fmt.Sprintf("Couldn't convert PNG: %s\n", filepath), err)
	}
}

func Jpg(filepath string) {
	cmd := generic_convert_command(filepath, ".jpg")
	if err := cmd.Run(); err != nil {
		log.Fatal(fmt.Sprintf("Couldn't convert JPG: %s\n", filepath), err)
	}
}

func Jpeg(filepath string) {
	cmd := generic_convert_command(filepath, ".jpeg")
	if err := cmd.Run(); err != nil {
		log.Fatal(fmt.Sprintf("Couldn't convert JPEG: %s\n", filepath), err)
	}
}

func generic_convert_command(filepath string, ext string) *exec.Cmd {
	trimmed_filepath := strings.TrimSuffix(filepath, ext)
	output_filepath := fmt.Sprintf("%s.webp", trimmed_filepath)
	return exec.Command("cwebp", "-q", "50", filepath, "-o", output_filepath)
}
