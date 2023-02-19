package post_test

import (
	"fmt"
	"testing"

	"github.com/EricDriussi/hugo-image-optimizer/internal/domain/post"
	"github.com/stretchr/testify/assert"
)

func Test_FileName(t *testing.T) {
	t.Run("errors out if given an empty file path", func(t *testing.T) {
		name, err := post.NewName("")

		assert.Error(t, err)
		assert.Equal(t, post.FileName{}, name)
	})

	t.Run("extracts the file name from a valid file path", func(t *testing.T) {
		expectedName := "aRandomName.md"
		path := fmt.Sprintf("a/random/path/%s", expectedName)
		name, err := post.NewName(path)

		assert.NoError(t, err)
		assert.Equal(t, expectedName, name.Value())
	})
}
