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
	aPath := "a/file/path/image.png"
	anotherPath := "a/file/path/to/another/image.jpg"
	imagePaths := []string{aPath, anotherPath}
	anError := errors.New("Something went wrong")

	t.Run("Deletes", func(t *testing.T) {
		unreferencedImage, imageErr := domain.NewImage(anotherPath)
		assert.NoError(t, imageErr)
		imageReference := "/path/image.png"

		t.Run("only unreferenced images", func(t *testing.T) {
			imageRepositoryMock := new(mocks.ImageRepository)
			imageRepositoryMock.On("Load").Return(imagePaths, nil)
			imageRepositoryMock.On("Delete", unreferencedImage).Return(nil)

			imageService := services.NewImage(imageRepositoryMock)
			serviceErr := imageService.RemoveAllExcept([]string{imageReference})
			imageRepositoryMock.AssertExpectations(t)

			assert.NoError(t, serviceErr)
		})

		t.Run("stopping if the repository erros out", func(t *testing.T) {
			imageRepositoryMock := new(mocks.ImageRepository)
			imageRepositoryMock.On("Load").Return(imagePaths, nil)
			imageRepositoryMock.On("Delete", unreferencedImage).Return(anError)

			imageService := services.NewImage(imageRepositoryMock)
			err := imageService.RemoveAllExcept([]string{imageReference})
			imageRepositoryMock.AssertExpectations(t)

			assert.Error(t, err)
			assert.ErrorContains(t, err, "Couldn't delete unused images :(")
		})
	})

	t.Run("Loads", func(t *testing.T) {
		t.Run("all images", func(t *testing.T) {
			imageRepositoryMock := new(mocks.ImageRepository)
			imageRepositoryMock.On("Load").Return(imagePaths, nil)

			imageService := services.NewImage(imageRepositoryMock)
			loadedImages, err := imageService.Load()
			imageRepositoryMock.AssertExpectations(t)

			assert.Len(t, loadedImages, 2)
			assert.NoError(t, err)
		})

		t.Run("discarding broken paths", func(t *testing.T) {
			brokenImagePaths := []string{aPath, "borked"}
			imageRepositoryMock := new(mocks.ImageRepository)
			imageRepositoryMock.On("Load").Return(brokenImagePaths, nil)

			imageService := services.NewImage(imageRepositoryMock)
			loadedImages, err := imageService.Load()
			imageRepositoryMock.AssertExpectations(t)

			assert.Len(t, loadedImages, 1)
			assert.Equal(t, aPath, loadedImages[0].Path())
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
			assert.ErrorContains(t, err, "Couldn't load images :(")
		})
	})

	t.Run("Converts", func(t *testing.T) {
		anImage, image1Err := domain.NewImage(aPath)
		assert.NoError(t, image1Err)
		anotherImage, image2Err := domain.NewImage(anotherPath)
		assert.NoError(t, image2Err)
		images := []domain.Image{anImage, anotherImage}

		t.Run("all images", func(t *testing.T) {
			imageRepositoryMock := new(mocks.ImageRepository)
			imageRepositoryMock.On("Load").Return(imagePaths, nil)
			imageRepositoryMock.On("ConvertToWebp", images).Return(nil)

			imageService := services.NewImage(imageRepositoryMock)
			err := imageService.Convert()

			imageRepositoryMock.AssertExpectations(t)
			assert.NoError(t, err)
		})

		t.Run("all images even if some conversions fail", func(t *testing.T) {
			imageRepositoryMock := new(mocks.ImageRepository)
			imageRepositoryMock.On("Load").Return(imagePaths, nil)
			imageRepositoryMock.On("ConvertToWebp", images).Return(anError)

			imageService := services.NewImage(imageRepositoryMock)
			err := imageService.Convert()
			imageRepositoryMock.AssertExpectations(t)

			assert.NoError(t, err)
		})
	})
}
