package post_test

import (
	"fmt"
	"testing"

	"github.com/EricDriussi/hugo-image-optimizer/internal/domain/post"
	"github.com/stretchr/testify/assert"
)

func Test_Content(t *testing.T) {
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
			assert.Equal(t, []byte(image_path), content.Images())
		})

		t.Run("from the front matter", func(t *testing.T) {
			image_path := "/path/src.png"

			image_reference := fmt.Sprintf("image: %s", image_path)

			rawContent := fmt.Sprintf(`%s
					line 1
					line 2`,
				image_reference)
			content := post.NewContent([]byte(rawContent))
			assert.Equal(t, []byte(image_path), content.Images())
		})

		t.Run("from both front matter and content", func(t *testing.T) {
			image_path1 := "path/src.png"
			image_path2 := "../path/src2.webp"
			image_path3 := "../path/src3.gif"

			image_reference1 := fmt.Sprintf("image: %s", image_path1)
			image_reference2 := fmt.Sprintf("![image](%s)", image_path2)
			image_reference3 := fmt.Sprintf("![image](%s)", image_path3)

			rawContent := fmt.Sprintf(`%s
					line 1
					line %s 2
					%s
					line 4`,
				image_reference1, image_reference2, image_reference3)

			content := post.NewContent([]byte(rawContent))

			assert.Contains(t, string(content.Images()), image_path1)
			assert.Contains(t, string(content.Images()), image_path2)
			assert.Contains(t, string(content.Images()), image_path3)
		})
	})
}
