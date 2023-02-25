package domain

import (
	"strings"

	"github.com/EricDriussi/hugo-image-optimizer/internal/domain/image"
)

type ImageRepository interface {
	Load() ([]string, error)
	Delete(Image) error
	ConvertToWebp([]Image) error
}

type Image struct {
	path      Path
	extension image.Extension
}

func NewImage(filepath string) (Image, error) {
	path, err := NewPath(filepath)
	if err != nil {
		return Image{}, err
	}

	extension, err := image.NewExtension(filepath)
	if err != nil {
		return Image{}, err
	}

	return Image{
		path:      path,
		extension: extension,
	}, nil
}

func (i Image) IsNotPresentIn(list_of_references []string) bool {
	for _, ref := range list_of_references {
		if strings.Contains(i.GetPath(), ref) {
			return false
		}
	}
	return true
}

func (i Image) GetExtension() string {
	return i.extension.Value()
}

func (i Image) IsGif() bool {
	return strings.EqualFold(i.extension.Value(), ".gif")
}

func (i Image) GetPath() string {
	return i.path.Value()
}
