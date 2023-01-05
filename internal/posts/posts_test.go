package posts_test

import (
	"hugo-images/internal/config"
	"hugo-images/internal/posts"
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestMain(m *testing.M) {
	viper.AddConfigPath("../../")
	viper.Set("dirs.project", "../../test/data/")
	viper.Set("dirs.posts", "posts/")
	config.Load()

	code := m.Run()
	os.Exit(code)
}

func TestReadsPosts(t *testing.T) {
	result := posts.List()
	if len(result) < 1 {
		t.Fail()
	}
}
