package services

import (
	"errors"

	"github.com/EricDriussi/hugo-image-optimizer/internal/domain"
)

type ImageService struct {
	imageRepository domain.ImageRepository
}

func NewImage(imageRepository domain.ImageRepository) ImageService {
	return ImageService{
		imageRepository: imageRepository,
	}
}

func (s ImageService) RemoveAllExcept(referenced_images []string) error {
	image_files, loadErr := s.Load()
	if loadErr != nil {
		return loadErr
	}

	for _, image := range image_files {
		isUnused := !image.IsPresentIn(referenced_images)
		if isUnused {
			s.imageRepository.Delete(image)
		}
	}

	return nil
}

func (s ImageService) Load() ([]domain.Image, error) {
	rawImages, err := s.imageRepository.Load()
	if err != nil {
		return nil, errors.New("Repository failed to load images")
	}

	images := s.buildImagesIgnoringInvalid(rawImages)
	return images, nil
}

func (s ImageService) buildImagesIgnoringInvalid(rawImages []string) []domain.Image {
	var images []domain.Image
	for _, path := range rawImages {
		image, err := domain.NewImage(path)
		if err == nil {
			images = append(images, image)
		}
	}
	return images
}
