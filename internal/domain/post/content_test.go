package post_test

import (
	"fmt"
	"testing"

	"github.com/EricDriussi/hugo-image-optimizer/internal/domain/post"
	"github.com/stretchr/testify/assert"
)

func Test_Content(t *testing.T) {
	t.Run("reads the full content", func(t *testing.T) {
		rawContent := []byte("some random content")
		content := post.NewContent(rawContent)
		assert.Equal(t, rawContent, content.Value())
	})

	t.Run("extracts the image references", func(t *testing.T) {
		t.Run("from the content", func(t *testing.T) {
			image_path := "../path/src.png"
			image_reference := fmt.Sprintf("![image](%s)", image_path)
			rawContent := fmt.Sprintf(`line 1
					line 2
					line %s 3
					line 4`,
				image_reference)

			content := post.NewContent([]byte(rawContent))
			assert.Contains(t, content.Images(), image_path)
		})

		t.Run("from the front matter", func(t *testing.T) {
			image_path := "/path/src.png"
			image_reference := fmt.Sprintf("image: %s", image_path)
			rawContent := fmt.Sprintf(`%s
					line 1
					line 2`,
				image_reference)

			content := post.NewContent([]byte(rawContent))
			assert.Contains(t, content.Images(), image_path)
		})

		t.Run("discarding false positives", func(t *testing.T) {
			png_ext := ".png"
			partial_path := "../../path/src"
			image_path := partial_path + png_ext
			rawContent := fmt.Sprintf(`line 1
					line 2
					line %s 3
					line 4`,
				image_path)

			content := post.NewContent([]byte(rawContent))

			assert.NotContains(t, content.Images(), image_path)
		})
	})

	t.Run("updates image references", func(t *testing.T) {
		t.Run("in content", func(t *testing.T) {
			png_ext := ".jpeg"
			partial_path := "../../another/path/pic"
			image_path := partial_path + png_ext
			image_reference := fmt.Sprintf("![image](%s)", image_path)
			rawContent := fmt.Sprintf(`line 1
					line 
					line %s 3
					line 4`,
				image_reference)

			content := post.NewContent([]byte(rawContent))
			content.UpdateImageReferences()

			assert.NotContains(t, string(content.Value()), image_path)
			assert.Contains(t, string(content.Value()), "![image]("+partial_path+".webp)")
		})

		t.Run("in front matter", func(t *testing.T) {
			png_ext := ".png"
			partial_path := "../../path/src"
			image_path := partial_path + png_ext
			image_reference := fmt.Sprintf("image: %s", image_path)
			rawContent := fmt.Sprintf(`%s
					line 2
					line 3
					line 4`,
				image_reference)

			content := post.NewContent([]byte(rawContent))
			content.UpdateImageReferences()

			assert.NotContains(t, string(content.Value()), image_path)
			assert.Contains(t, string(content.Value()), "image: "+partial_path+".webp")
		})

		t.Run("discarding false positives", func(t *testing.T) {
			png_ext := ".png"
			partial_path := "../../path/src"
			image_path := partial_path + png_ext
			rawContent := fmt.Sprintf(`line 1
					line 2
					line %s 3
					line 4`,
				image_path)

			content := post.NewContent([]byte(rawContent))
			content.UpdateImageReferences()

			assert.Contains(t, string(content.Value()), image_path)
			assert.NotContains(t, string(content.Value()), partial_path+".webp")
		})
	})
}
