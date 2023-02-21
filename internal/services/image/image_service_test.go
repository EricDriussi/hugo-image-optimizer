package services_test

import (
	"errors"
	"testing"

	"github.com/EricDriussi/hugo-image-optimizer/internal/domain"
	services "github.com/EricDriussi/hugo-image-optimizer/internal/services/image"
	"github.com/EricDriussi/hugo-image-optimizer/test/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_ImageService(t *testing.T) {
	pathOne := "a/file/path/image.png"
	pathTwo := "a/file/path/to/another/image.jpg"
	images := []string{pathOne, pathTwo}

	t.Run("Deletes only unreferenced images", func(t *testing.T) {
		image, image_err := domain.NewImage(pathTwo)
		assert.NoError(t, image_err)

		imageRepositoryMock := new(mocks.ImageRepository)
		imageRepositoryMock.On("Load").Return(images, nil)
		imageRepositoryMock.On("Delete", image).Return(nil)

		imageService := services.NewImage(imageRepositoryMock)
		image_reference := "/path/image.png"
		service_err := imageService.RemoveAllExcept([]string{image_reference})
		imageRepositoryMock.AssertExpectations(t)

		assert.NoError(t, service_err)
	})

	t.Run("Loads all images", func(t *testing.T) {
		imageRepositoryMock := new(mocks.ImageRepository)
		imageRepositoryMock.On("Load").Return(images, nil)

		imageService := services.NewImage(imageRepositoryMock)
		loadedImages, err := imageService.Load()
		imageRepositoryMock.AssertExpectations(t)

		assert.Len(t, loadedImages, 2)
		assert.NoError(t, err)
	})

	t.Run("Discards broken paths when loading", func(t *testing.T) {
		images := []string{pathOne, "borked"}
		imageRepositoryMock := new(mocks.ImageRepository)
		imageRepositoryMock.On("Load").Return(images, nil)

		imageService := services.NewImage(imageRepositoryMock)
		loadedImages, err := imageService.Load()
		imageRepositoryMock.AssertExpectations(t)

		assert.Len(t, loadedImages, 1)
		assert.Equal(t, pathOne, loadedImages[0].GetPath())
		assert.NoError(t, err)
	})

	t.Run("No partial loading if the repository erros out", func(t *testing.T) {
		anError := errors.New("Something went wrong")
		imageRepositoryMock := new(mocks.ImageRepository)
		imageRepositoryMock.On("Load").Return(nil, anError)

		imageService := services.NewImage(imageRepositoryMock)
		loadedImages, err := imageService.Load()
		imageRepositoryMock.AssertExpectations(t)

		assert.Nil(t, loadedImages)
		assert.Error(t, err)
		assert.ErrorContains(t, err, "Repository failed to load images")
	})
}
