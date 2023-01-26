package imageService_test

import (
	"hugo-images/internal/config"
	imageService "hugo-images/internal/image_service"
	"hugo-images/internal/util"
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestMain(m *testing.M) {
	os.Chdir("../../")

	viper.Set("dirs.project", "test/data/")
	viper.Set("dirs.images", "images/")
	config.Load()

	code := m.Run()
	os.Exit(code)
}

var peter = `![image](../images/a_gif.gif)![image](../images/another_image.jpeg)`

func TestExcludesDirsFromConfig(t *testing.T) {
	viper.Set("dirs.images_exclude", "whoami donation")

	images := imageService.ImagesInIncludedDirs()

	listsIgnoredImages := util.StringIsInArray(images, "avatar.jpg")
	doesNotListIncludedImage := !util.StringIsInArray(images, "an_image.png")
	if listsIgnoredImages || doesNotListIncludedImage {
		t.Fail()
	}
}
