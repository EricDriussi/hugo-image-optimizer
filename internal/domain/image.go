package domain

import "github.com/EricDriussi/hugo-image-optimizer/internal/domain/image"

type ImageRepository interface{}

type Image struct {
	path      Path
	extension image.Extension
}

// TODO.add test!
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
