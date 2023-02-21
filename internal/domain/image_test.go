package domain_test

import (
	"fmt"
	"testing"

	"github.com/EricDriussi/hugo-image-optimizer/internal/domain"
	"github.com/stretchr/testify/assert"
)

func Test_Image(t *testing.T) {
	t.Run("errors out if given an invalid file", func(t *testing.T) {
		invalid_filepaths := []string{"path/with/noExtension", ""}
		for _, filepath := range invalid_filepaths {
			image, err := domain.NewImage(filepath)

			assert.Error(t, err)
			assert.Equal(t, domain.Image{}, image)
		}
	})

	t.Run("builds as expected with no errors", func(t *testing.T) {
		extension := ".png"
		filepath := fmt.Sprintf("path/with/validExtension%s", extension)
		image, err := domain.NewImage(filepath)

		assert.NoError(t, err)
		assert.Equal(t, extension, image.GetExtension())
		assert.Equal(t, filepath, image.GetPath())
	})

	t.Run("checks if is present in reference list", func(t *testing.T) {
		filename := "valid.png"
		filepath := fmt.Sprintf("/a/long/path/to/%s", filename)

		image, err := domain.NewImage(filepath)
		assert.NoError(t, err)

		matching_reference := fmt.Sprintf("/path/to/%s", filename)
		non_matching_reference := fmt.Sprintf("/irrelevant/path/%s", filename)
		image_references := []string{matching_reference, non_matching_reference}

		assert.False(t, image.IsNotPresentIn(image_references))
		assert.True(t, image.IsNotPresentIn([]string{non_matching_reference}))
	})
}
