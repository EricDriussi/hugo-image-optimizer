package filesystemrepo_test

import (
	"testing"

	filesystemrepo "github.com/EricDriussi/hugo-image-optimizer/internal/infrastructure/repos/filesystem_repo/post"
	"github.com/stretchr/testify/assert"
)

func Test_PostRepository_Load(t *testing.T) {
	runWithFixtures(t, func() {
		postRepoTests(t)
	})
}

func postRepoTests(t *testing.T) {
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
