package filesystemrepo_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/EricDriussi/hugo-image-optimizer/internal/domain"
	filesystemrepo "github.com/EricDriussi/hugo-image-optimizer/internal/infrastructure/repos/filesystem_repo/image"
	"github.com/stretchr/testify/assert"
)

func TestImageRepository(t *testing.T) {
	runWithFixtures(t, func() {
		imageRepoTests(t)
	})
}

func imageRepoTests(t *testing.T) {
	imagesTestDir := "test/data/images/"
	imagesTestExcludedDirs := []string{"whoami", "donation"}

	t.Run("#LOAD", func(t *testing.T) {
		t.Run("Loads all images recursively if directory exists", func(t *testing.T) {
			repo := filesystemrepo.NewImage(imagesTestDir, []string{})

			loadedImages, err := repo.Load()

			assert.Len(t, loadedImages, 5)
			assert.NoError(t, err)
		})

		t.Run("Doesn't load images from excluded directories", func(t *testing.T) {
			repo := filesystemrepo.NewImage(imagesTestDir, imagesTestExcludedDirs)

			loadedImages, err := repo.Load()

			assert.Len(t, loadedImages, 3)
			assert.NoError(t, err)
		})

		t.Run("Errors out if given a non existent directory", func(t *testing.T) {
			repo := filesystemrepo.NewImage("test/data/images/NON_EXISTENT/", []string{})

			loadedImages, err := repo.Load()

			assert.Len(t, loadedImages, 0)
			assert.Error(t, err)
		})
	})

	t.Run("#DELETE", func(t *testing.T) {
		t.Run("Doesn't delete images from excluded directories", func(t *testing.T) {
			repo := filesystemrepo.NewImage(imagesTestDir, imagesTestExcludedDirs)

			filename := fmt.Sprintf("%s%s%s", imagesTestDir, imagesTestExcludedDirs[0], "/testFile.png")
			image, imageErr := domain.NewImage(filename)
			assert.NoError(t, imageErr)

			f, createErr := os.Create(filename)
			defer f.Close()
			assert.NoError(t, createErr)

			repoErr := repo.Delete(image)
			assert.NoError(t, repoErr)

			rmErr := os.Remove(filename)
			assert.NoError(t, rmErr)
		})

		t.Run("Deletes an image", func(t *testing.T) {
			repo := filesystemrepo.NewImage(imagesTestDir, imagesTestExcludedDirs)

			filename := fmt.Sprintf("%s%s", imagesTestDir, "testFile.png")
			image, imageErr := domain.NewImage(filename)
			assert.NoError(t, imageErr)

			f, createErr := os.Create(filename)
			defer f.Close()
			assert.NoError(t, createErr)

			repoErr := repo.Delete(image)
			assert.NoError(t, repoErr)

			rmErr := os.Remove(filename)
			assert.Error(t, rmErr)
		})
	})

	t.Run("#CONVERT", func(t *testing.T) {
		t.Run("Doesn't convert images from excluded directories", func(t *testing.T) {
			repo := filesystemrepo.NewImage(imagesTestDir, imagesTestExcludedDirs)

			filename := fmt.Sprintf("%s%s%s", imagesTestDir, imagesTestExcludedDirs[0],
				"/avatar")
			jpegFilename := fmt.Sprintf("%s%s", filename, ".jpg")
			webpFilename := fmt.Sprintf("%s%s", filename, ".webp")
			image, imageErr := domain.NewImage(jpegFilename)
			assert.NoError(t, imageErr)

			f, createErr := os.Create(filename)
			defer f.Close()
			assert.NoError(t, createErr)

			repoErr := repo.ConvertToWebp([]domain.Image{image})
			assert.NoError(t, repoErr)

			rmErr := os.Remove(webpFilename)
			assert.Error(t, rmErr)
		})

		t.Run("Converts a PNG image to webp", func(t *testing.T) {
			repo := filesystemrepo.NewImage(imagesTestDir, imagesTestExcludedDirs)

			filename := fmt.Sprintf("%s%s", imagesTestDir, "an_image")
			pngFilename := fmt.Sprintf("%s%s", filename, ".png")
			webpFilename := fmt.Sprintf("%s%s", filename, ".webp")

			image, imageErr := domain.NewImage(pngFilename)
			assert.NoError(t, imageErr)

			repoErr := repo.ConvertToWebp([]domain.Image{image})
			assert.NoError(t, repoErr)

			rmErr := os.Remove(webpFilename)
			assert.NoError(t, rmErr)
		})

		t.Run("Converts a JPEG image to webp", func(t *testing.T) {
			repo := filesystemrepo.NewImage(imagesTestDir, imagesTestExcludedDirs)

			filename := fmt.Sprintf("%s%s", imagesTestDir, "another_image")
			jpegFilename := fmt.Sprintf("%s%s", filename, ".jpeg")
			webpFilename := fmt.Sprintf("%s%s", filename, ".webp")

			image, imageErr := domain.NewImage(jpegFilename)
			assert.NoError(t, imageErr)

			repoErr := repo.ConvertToWebp([]domain.Image{image})
			assert.NoError(t, repoErr)

			rmErr := os.Remove(webpFilename)
			assert.NoError(t, rmErr)
		})

		t.Run("Converts a GIF to webp", func(t *testing.T) {
			repo := filesystemrepo.NewImage(imagesTestDir, imagesTestExcludedDirs)

			filename := fmt.Sprintf("%s%s", imagesTestDir, "a_gif")
			gifFilename := fmt.Sprintf("%s%s", filename, ".gif")
			webpFilename := fmt.Sprintf("%s%s", filename, ".webp")

			image, imageErr := domain.NewImage(gifFilename)
			assert.NoError(t, imageErr)

			repoErr := repo.ConvertToWebp([]domain.Image{image})
			assert.NoError(t, repoErr)

			rmErr := os.Remove(webpFilename)
			assert.NoError(t, rmErr)
		})
	})
}
