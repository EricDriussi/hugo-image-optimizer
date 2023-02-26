package services

import (
	"errors"
	"fmt"

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
	if convErr := s.imageRepository.ConvertToWebp(loadedImages); convErr != nil {
		fmt.Println("[WARNING]: Some images were not converted :(")
	}
	return nil
}

func (s ImageService) RemoveAllExcept(imageReferences []string) error {
	loadedImages, err := s.Load()
	if err != nil {
		return err
	}
	for _, image := range loadedImages {
		if image.IsNotPresentIn(imageReferences) {
			rmErr := s.imageRepository.Delete(image)
			if rmErr != nil {
				return errors.New("Couldn't delete unused images :(")
			}
		}
	}

	return nil
}

func (s ImageService) Load() ([]domain.Image, error) {
	rawImages, err := s.imageRepository.Load()
	if err != nil {
		return nil, errors.New("Couldn't load images :(")
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
