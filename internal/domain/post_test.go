package domain_test

import (
	"fmt"
	"testing"

	"github.com/EricDriussi/hugo-image-optimizer/internal/domain"
	"github.com/stretchr/testify/assert"
)

func Test_PostDomain_Constructor(t *testing.T) {
	t.Run("errors out", func(t *testing.T) {
		t.Run("if given an empty file path", func(t *testing.T) {
			post, err := domain.NewPost("", []byte("Some content"))

			assert.Error(t, err)
			assert.Equal(t, domain.Post{}, post)
		})
	})

	t.Run("extracts the file name from a valid file path", func(t *testing.T) {
		name := "aRandomName.md"
		path := fmt.Sprintf("a/random/path/%s", name)
		post, err := domain.NewPost(path, []byte("Some content"))

		assert.NoError(t, err)
		assert.Equal(t, name, post.GetFilename())
	})

	t.Run("extracts the referenced images", func(t *testing.T) {
		name := "aRandomName.md"
		path := fmt.Sprintf("a/random/path/%s", name)

		t.Run("from the content", func(t *testing.T) {
			image_path := "../path/src.png"
			image_reference := fmt.Sprintf("![image](%s)", image_path)
			content := fmt.Sprintf(`line 1
					line 2
					line %s 3
					line 4`,
				image_reference)
			post, err := domain.NewPost(path, []byte(content))

			assert.NoError(t, err)
			assert.Equal(t, image_path, post.GetReferencedImages())
		})

		t.Run("from the front matter", func(t *testing.T) {
			image_path := "/path/src.png"

			image_reference := fmt.Sprintf("image: %s", image_path)

			content := fmt.Sprintf(`%s
					line 1
					line 2`,
				image_reference)
			post, err := domain.NewPost(path, []byte(content))

			assert.NoError(t, err)
			assert.Equal(t, image_path, post.GetReferencedImages())
		})

		t.Run("from both front matter and content", func(t *testing.T) {
			image_path1 := "path/src.png"
			image_path2 := "../path/src2.webp"
			image_path3 := "../path/src3.gif"

			image_reference1 := fmt.Sprintf("image: %s", image_path1)
			image_reference2 := fmt.Sprintf("![image](%s)", image_path2)
			image_reference3 := fmt.Sprintf("![image](%s)", image_path3)

			content := fmt.Sprintf(`%s
					line 1
					line %s 2
					%s
					line 4`,
				image_reference1, image_reference2, image_reference3)
			post, err := domain.NewPost(path, []byte(content))

			assert.NoError(t, err)
			assert.Contains(t, post.GetReferencedImages(), image_path1)
			assert.Contains(t, post.GetReferencedImages(), image_path2)
			assert.Contains(t, post.GetReferencedImages(), image_path3)
		})
	})
}
