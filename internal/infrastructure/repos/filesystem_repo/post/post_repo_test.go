package filesystemrepo_test

import (
	"os"
	"path"
	"runtime"
	"testing"

	filesystemrepo "github.com/EricDriussi/hugo-image-optimizer/internal/infrastructure/repos/filesystem_repo/post"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func Test_PostRepository_Load(t *testing.T) {
	t.Run("Loads all posts if directory exists", func(t *testing.T) {
		repo := filesystemrepo.NewPost("test/data/posts/")

		loadedPosts, err := repo.Load()

		assert.Len(t, loadedPosts, 3)
		assert.NoError(t, err)
	})

	t.Run("Errors out if given a non existent directory", func(t *testing.T) {
		repo := filesystemrepo.NewPost("test/data/posts/NON_EXISTENT/")

		loadedPosts, err := repo.Load()

		assert.Len(t, loadedPosts, 0)
		assert.Error(t, err)
	})
}

func setCWDToProjectRoot() {
	_, filename, _, _ := runtime.Caller(0)
	project_root := path.Join(path.Dir(filename), "../../../../..")
	if err := os.Chdir(project_root); err != nil {
		panic(err)
	}
}

func setup() {
	setCWDToProjectRoot()
	os.MkdirAll("test/data/posts/subdir", os.ModePerm)
	f1, _ := os.Create("test/data/posts/a_post.md")
	defer f1.Close()
	f2, _ := os.Create("test/data/posts/another_post.md")
	defer f2.Close()
	f3, _ := os.Create("test/data/posts/subdir/a_different_post.md")
	defer f3.Close()
}

func shutdown() {
	os.RemoveAll("test/data/posts")
}
