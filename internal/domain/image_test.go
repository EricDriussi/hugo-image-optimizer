package domain_test

import (
	"fmt"
	"testing"

	"github.com/EricDriussi/hugo-image-optimizer/internal/domain"
	"github.com/stretchr/testify/assert"
)

func Test_DomainImage(t *testing.T) {
	t.Run("errors out if given an invalid file", func(t *testing.T) {
		invalidFilepaths := []string{"path/with/noExtension", ""}
		for _, filepath := range invalidFilepaths {
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
		assert.Equal(t, extension, image.Extension())
		assert.Equal(t, filepath, image.Path())
	})

	t.Run("checks if is present in reference list", func(t *testing.T) {
		filename := "valid.png"
		filepath := fmt.Sprintf("/a/long/path/to/%s", filename)

		image, err := domain.NewImage(filepath)
		assert.NoError(t, err)

		matchingReference := fmt.Sprintf("/path/to/%s", filename)
		nonMatchingReference := fmt.Sprintf("/irrelevant/path/%s", filename)

		assert.False(t, image.IsNotPresentIn([]string{matchingReference, nonMatchingReference}))
		assert.True(t, image.IsNotPresentIn([]string{nonMatchingReference}))
	})
}
