package imageService_test

import (
	"optimize/internal/config"
	imageService "optimize/internal/image_service"
	"optimize/internal/util"
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestMain(m *testing.M) {
	os.Chdir("../../")

	viper.Set("dirs.images", "test/data/images/")
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
