package domain_test

import (
	"fmt"
	"testing"

	"github.com/EricDriussi/hugo-image-optimizer/internal/domain"
	"github.com/stretchr/testify/assert"
)

func Test_DomainPost(t *testing.T) {
	t.Run("errors out if given an empty file path", func(t *testing.T) {
		post, err := domain.NewPost("", []byte("Some content"))

		assert.Error(t, err)
		assert.Equal(t, domain.Post{}, post)
	})

	t.Run("builds as expected with no errors", func(t *testing.T) {
		path := "a/random/path/filename.md"
		content := []byte("some content")

		post, err := domain.NewPost(path, content)

		assert.NoError(t, err)
		assert.Equal(t, path, post.Path())
		assert.Equal(t, string(content), post.Content())
	})

	t.Run("cleans image references", func(t *testing.T) {
		path := "a/random/path/filename.md"

		imagePath := "/path/src.png"
		imageReference := fmt.Sprintf("![image](../.././%s)", imagePath)
		content := contentWithImage(imageReference)

		post, err := domain.NewPost(path, content)

		assert.NoError(t, err)
		assert.Equal(t, path, post.Path())
		assert.Equal(t, string(content), post.Content())
		assert.Contains(t, post.ReferencedImagePaths(), imagePath)
	})
}

func contentWithImage(reference string) []byte {
	content := fmt.Sprintf(`line 1
					line %s 2
					line 4`,
		reference)
	return []byte(content)
}
