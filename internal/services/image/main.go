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

func (s ImageService) Convert() error {
	loadedImages, err := s.Load()
	if err != nil {
		return err
	}

	for _, image := range loadedImages {
		rm_err := s.imageRepository.ConvertToWebp(image)
		if rm_err != nil {
			return errors.New("Failed to convert images")
		}
	}
	return nil
}

func (s ImageService) RemoveAllExcept(image_references []string) error {
	loadedImages, err := s.Load()
	if err != nil {
		return err
	}
	for _, image := range loadedImages {
		if image.IsNotPresentIn(image_references) {
			rm_err := s.imageRepository.Delete(image)
			if rm_err != nil {
				return errors.New("Failed to delete images")
			}
		}
	}

	return nil
}

func (s ImageService) Load() ([]domain.Image, error) {
	rawImages, err := s.imageRepository.Load()
	if err != nil {
		return nil, errors.New("Failed to load images")
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
