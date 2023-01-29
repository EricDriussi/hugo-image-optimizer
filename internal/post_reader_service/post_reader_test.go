package postReader_test

import (
	"optimize/internal/config"
	postReader "optimize/internal/post_reader_service"
	"optimize/internal/util"
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestMain(m *testing.M) {
	os.Chdir("../../")

	viper.Set("dirs.posts", "test/data/posts/")
	config.Load()

	code := m.Run()
	os.Exit(code)
}

func TestReadsAllPosts(t *testing.T) {
	all_posts := postReader.All_posts_as_bytes()

	doesNotContainExpectedStrings := !util.ByteArrayContainsString("One", &all_posts) && !util.ByteArrayContainsString("Two", &all_posts) && !util.ByteArrayContainsString("Three", &all_posts)
	if doesNotContainExpectedStrings {
		t.Fail()
	}
}
