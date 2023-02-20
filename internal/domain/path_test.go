package domain_test

import (
	"testing"

	"github.com/EricDriussi/hugo-image-optimizer/internal/domain"
	"github.com/stretchr/testify/assert"
)

func Test_Path(t *testing.T) {
	t.Run("errors out if given an empty file path", func(t *testing.T) {
		name, err := domain.NewPath("")
		assert.Error(t, err)
		assert.Equal(t, domain.Path{}, name)
	})
}
