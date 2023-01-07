package images_test

import (
	"hugo-images/internal/config"
	"hugo-images/internal/images"
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestMain(m *testing.M) {
	viper.AddConfigPath("../../")
	viper.Set("dirs.project", "../../test/data/")
	viper.Set("dirs.images", "images/")
	config.Load()

	code := m.Run()
	os.Exit(code)
}

func TestReadsImages(t *testing.T) {
	result := images.List()
	if len(result) < 1 {
		t.Fail()
	}
}

func TestExcludesImages(t *testing.T) {
	result := images.List()
	if contains(result, "avatar.jpg") {
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
