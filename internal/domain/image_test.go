package domain_test

import (
	"strings"
	"testing"

	"github.com/EricDriussi/hugo-image-optimizer/internal/domain"
	"github.com/stretchr/testify/assert"
)

func Test_Image(t *testing.T) {
	t.Run("errors out if given an invalid file", func(t *testing.T) {
		filepath := "path/with/noExtension"
		image, err := domain.NewImage(filepath)

		assert.Error(t, err)
		assert.Equal(t, domain.Image{}, image)
	})

	t.Run("builds as expected with no errors", func(t *testing.T) {
		extension := ".png"
		filepath := strings.Join([]string{"path/with/validExtension", extension}, "")
		image, err := domain.NewImage(filepath)

		assert.NoError(t, err)
		assert.Equal(t, extension, image.GetExtension())
		assert.Equal(t, filepath, image.GetPath())
	})

	t.Run("checks if is present in reference list", func(t *testing.T) {
		filename := "valid.png"
		image_path := strings.Join([]string{"path/with/", filename}, "")

		image_ref1 := strings.Join([]string{"/with/", filename}, "")

		image, err := domain.NewImage(image_path)
		assert.NoError(t, err)

		present := image.IsPresentIn([]string{image_ref1})
		assert.True(t, present)
	})
}
