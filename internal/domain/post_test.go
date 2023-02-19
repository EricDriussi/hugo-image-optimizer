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
		name := "aRandomName.md"
		post_path := fmt.Sprintf("a/random/path/%s", name)

		image_path := "../path/src2.png"
		image_reference := fmt.Sprintf("![image](%s)", image_path)
		content := fmt.Sprintf(`line 1
					line %s 2
					line 4`,
			image_reference)

		post, err := domain.NewPost(post_path, []byte(content))

		assert.NoError(t, err)
		assert.Equal(t, name, post.GetFilename())
		assert.Equal(t, post_path, post.GetPath())
		assert.Equal(t, content, post.GetFullContent())
		assert.Equal(t, image_path, post.GetReferencedImages())
	})
}
