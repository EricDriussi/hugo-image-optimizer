package images_test

import (
	"hugo-images/internal/config"
	"hugo-images/internal/images"
	"hugo-images/internal/posts"
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestMain(m *testing.M) {
	viper.AddConfigPath("../../")
	viper.Set("dirs.project", "../../test/data/")
	viper.Set("dirs.posts", "posts/")
	viper.Set("dirs.images", "images/")
	config.Load()

	code := m.Run()
	os.Exit(code)
}

func TestListsAllImages(t *testing.T) {
	result := images.ListAll()
	if len(result) < 1 {
		t.Fail()
	}
}

func TestListsAllExcludingImages(t *testing.T) {
	result := images.ListAll()
	if contains(result, "avatar.jpg") {
		t.Fail()
	}
}

func TestListsUnusedImages(t *testing.T) {
	posts := posts.List()
	imagesList := images.ListUnusedIn(posts)
	if imagesList["an_image.png"] {
		t.Fail()
	}
}

func TestListsUnusedImagesIgnoresUsedImages(t *testing.T) {
	posts := posts.List()
	imagesList := images.ListUnusedIn(posts)
	if !imagesList["a_gif.gif"] || !imagesList["another_image.jpeg"] {
		t.Fail()
	}
}

func contains(mp map[string]bool, elem string) bool {
	for key := range mp {
		if key == elem {
			return true
		}
	}
	return false
}
