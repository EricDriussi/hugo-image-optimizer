package main

import (
	"fmt"
	"hugo-images/internal/config"
	"hugo-images/internal/images"
	"hugo-images/internal/posts"
)

func main() {
	config.Load()
	images := images.List()
	fmt.Println(images)
	posts := posts.List()
	fmt.Println(posts)
}
