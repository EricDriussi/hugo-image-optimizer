package image_test

import (
	"strings"
	"testing"

	"github.com/EricDriussi/hugo-image-optimizer/internal/domain/image"
	"github.com/stretchr/testify/assert"
)

func Test_Extension(t *testing.T) {
	t.Run("handles all valid image extensions", func(t *testing.T) {
		png := ".png"
		jpg := ".jpg"
		jpeg := ".jpeg"
		gif := ".gif"
		valid_extensions := []string{png, jpg, jpeg, gif}

		for _, ext := range valid_extensions {
			filename := strings.Join([]string{"aName", ext}, "")

			extension, err := image.NewExtension(filename)

			assert.NoError(t, err)
			assert.Equal(t, ext, extension.Value())
		}
	})

	t.Run("errors out if", func(t *testing.T) {
		t.Run("if given a file ending in .", func(t *testing.T) {
			extension, err := image.NewExtension("file.")
			assert.Error(t, err)
			assert.Equal(t, image.Extension{}, extension)
		})

		t.Run("if given a file with no extension", func(t *testing.T) {
			extension, err := image.NewExtension("file")
			assert.Error(t, err)
			assert.Equal(t, image.Extension{}, extension)
		})

		t.Run("if given a file with unsupported extension", func(t *testing.T) {
			extension, err := image.NewExtension("file.wtf")
			assert.Error(t, err)
			assert.Equal(t, image.Extension{}, extension)
		})

		t.Run("if given an empty string", func(t *testing.T) {
			extension, err := image.NewExtension("")
			assert.Error(t, err)
			assert.Equal(t, image.Extension{}, extension)
		})
	})
}
