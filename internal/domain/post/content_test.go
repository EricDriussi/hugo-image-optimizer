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
			imagePath := "../path/src.png"
			imageReference := fmt.Sprintf("![image](%s)", imagePath)
			rawContent := fmt.Sprintf(`line 1
					line 2
					line %s 3
					line 4`,
				imageReference)

			content := post.NewContent([]byte(rawContent))
			assert.Contains(t, content.Images(), imagePath)
		})

		t.Run("from the front matter", func(t *testing.T) {
			imagePath := "/path/src.png"
			imageReference := fmt.Sprintf("image: %s", imagePath)
			rawContent := fmt.Sprintf(`%s
					line 1
					line 2`,
				imageReference)

			content := post.NewContent([]byte(rawContent))
			assert.Contains(t, content.Images(), imagePath)
		})

		t.Run("discarding false positives", func(t *testing.T) {
			pngExt := ".png"
			partialPath := "../../path/src"
			imagePath := partialPath + pngExt
			rawContent := fmt.Sprintf(`line 1
					line 2
					line %s 3
					line 4`,
				imagePath)

			content := post.NewContent([]byte(rawContent))

			assert.NotContains(t, content.Images(), imagePath)
		})
	})

	t.Run("updates image references", func(t *testing.T) {
		t.Run("in content", func(t *testing.T) {
			pngExt := ".jpeg"
			partialPath := "../../another/path/pic"
			imagePath := partialPath + pngExt
			imageReference := fmt.Sprintf("![image](%s)", imagePath)
			rawContent := fmt.Sprintf(`line 1
					line 
					line %s 3
					line 4`,
				imageReference)

			content := post.NewContent([]byte(rawContent))
			content.ChangeImgExtToWebp()

			assert.NotContains(t, string(content.Value()), imagePath)
			assert.Contains(t, string(content.Value()), "![image]("+partialPath+".webp)")
		})

		t.Run("in front matter", func(t *testing.T) {
			pngExt := ".png"
			partialPath := "../../path/src"
			imagePath := partialPath + pngExt
			imageReference := fmt.Sprintf("image: %s", imagePath)
			rawContent := fmt.Sprintf(`%s
					line 2
					line 3
					line 4`,
				imageReference)

			content := post.NewContent([]byte(rawContent))
			content.ChangeImgExtToWebp()

			assert.NotContains(t, string(content.Value()), imagePath)
			assert.Contains(t, string(content.Value()), "image: "+partialPath+".webp")
		})

		t.Run("discarding false positives", func(t *testing.T) {
			pngExt := ".png"
			partialPath := "../../path/src"
			imagePath := partialPath + pngExt
			rawContent := fmt.Sprintf(`line 1
					line 2
					line %s 3
					line 4`,
				imagePath)

			content := post.NewContent([]byte(rawContent))
			content.ChangeImgExtToWebp()

			assert.Contains(t, string(content.Value()), imagePath)
			assert.NotContains(t, string(content.Value()), partialPath+".webp")
		})
	})
}
