package post_test

import (
	"testing"

	"github.com/EricDriussi/hugo-image-optimizer/internal/domain/post"
	"github.com/stretchr/testify/assert"
)

func Test_PostPath(t *testing.T) {
	t.Run("errors out if given an empty file path", func(t *testing.T) {
		name, err := post.NewPath("")
		assert.Error(t, err)
		assert.Equal(t, post.Path{}, name)
	})
}
