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
	anError := errors.New("Something went wrong")

	t.Run("Deletes", func(t *testing.T) {
		image, image_err := domain.NewImage(pathTwo)
		assert.NoError(t, image_err)
		image_reference := "/path/image.png"

		t.Run("only unreferenced images", func(t *testing.T) {
			imageRepositoryMock := new(mocks.ImageRepository)
			imageRepositoryMock.On("Load").Return(images, nil)
			imageRepositoryMock.On("Delete", image).Return(nil)

			imageService := services.NewImage(imageRepositoryMock)
			service_err := imageService.RemoveAllExcept([]string{image_reference})
			imageRepositoryMock.AssertExpectations(t)

			assert.NoError(t, service_err)
		})

		t.Run("stopping if the repository erros out", func(t *testing.T) {
			imageRepositoryMock := new(mocks.ImageRepository)
			imageRepositoryMock.On("Load").Return(images, nil)
			imageRepositoryMock.On("Delete", image).Return(anError)

			imageService := services.NewImage(imageRepositoryMock)
			err := imageService.RemoveAllExcept([]string{image_reference})
			imageRepositoryMock.AssertExpectations(t)

			assert.Error(t, err)
			assert.ErrorContains(t, err, "Failed to delete images")
		})
	})

	t.Run("Loads", func(t *testing.T) {
		t.Run("all images", func(t *testing.T) {
			imageRepositoryMock := new(mocks.ImageRepository)
			imageRepositoryMock.On("Load").Return(images, nil)

			imageService := services.NewImage(imageRepositoryMock)
			loadedImages, err := imageService.Load()
			imageRepositoryMock.AssertExpectations(t)

			assert.Len(t, loadedImages, 2)
			assert.NoError(t, err)
		})

		t.Run("discarding broken paths", func(t *testing.T) {
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

		t.Run("stopping if the repository erros out", func(t *testing.T) {
			imageRepositoryMock := new(mocks.ImageRepository)
			imageRepositoryMock.On("Load").Return(nil, anError)

			imageService := services.NewImage(imageRepositoryMock)
			loadedImages, err := imageService.Load()
			imageRepositoryMock.AssertExpectations(t)

			assert.Nil(t, loadedImages)
			assert.Error(t, err)
			assert.ErrorContains(t, err, "Failed to load images")
		})
	})

	t.Run("Converts", func(t *testing.T) {
		images := []string{pathOne, pathTwo}
		image1, image1_err := domain.NewImage(pathOne)
		assert.NoError(t, image1_err)
		image2, image2_err := domain.NewImage(pathTwo)
		assert.NoError(t, image2_err)

		t.Run("all images", func(t *testing.T) {
			imageRepositoryMock := new(mocks.ImageRepository)
			imageRepositoryMock.On("Load").Return(images, nil)
			imageRepositoryMock.On("ConvertToWebp", []domain.Image{image1, image2}).Return(nil)

			imageService := services.NewImage(imageRepositoryMock)
			err := imageService.Convert()

			imageRepositoryMock.AssertExpectations(t)
			assert.NoError(t, err)
		})

		t.Run("all images even if some conversions fail", func(t *testing.T) {
			imageRepositoryMock := new(mocks.ImageRepository)
			imageRepositoryMock.On("Load").Return(images, nil)
			imageRepositoryMock.On("ConvertToWebp", []domain.Image{image1, image2}).Return(anError)

			imageService := services.NewImage(imageRepositoryMock)
			err := imageService.Convert()
			imageRepositoryMock.AssertExpectations(t)

			assert.NoError(t, err)
		})
	})
}
