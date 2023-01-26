package main

import (
	"fmt"
	"hugo-images/internal/config"
	"hugo-images/internal/image_service"
)

func main() {
	config.Load()
	images := imageService.ImagesInIncludedDirs()
	fmt.Println(images)
}
