package domain_test

import (
	"fmt"
	"testing"

	"github.com/EricDriussi/hugo-image-optimizer/internal/domain"
	"github.com/stretchr/testify/assert"
)

func Test_PostDomain_Constructor(t *testing.T) {
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
		assert.Equal(t, path, post.GetPath())
		assert.Equal(t, string(content), post.GetFullContent())
	})

	t.Run("cleans image references", func(t *testing.T) {
		path := "a/random/path/filename.md"

		image_path := "/path/src.png"
		image_reference := fmt.Sprintf("![image](../.././%s)", image_path)
		content := contentWithImage(image_reference)

		post, err := domain.NewPost(path, content)

		assert.NoError(t, err)
		assert.Equal(t, path, post.GetPath())
		assert.Equal(t, string(content), post.GetFullContent())
		assert.Contains(t, post.GetCleanImageReferences(), image_path)
	})
}

func contentWithImage(reference string) []byte {
	content := fmt.Sprintf(`line 1
					line %s 2
					line 4`,
		reference)
	return []byte(content)
}
