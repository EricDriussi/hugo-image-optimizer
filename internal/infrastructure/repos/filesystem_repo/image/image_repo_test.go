package filesystemrepo_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/EricDriussi/hugo-image-optimizer/internal/domain"
	filesystemrepo "github.com/EricDriussi/hugo-image-optimizer/internal/infrastructure/repos/filesystem_repo/image"
	"github.com/stretchr/testify/assert"
)

func Test_ImageRepository(t *testing.T) {
	runWithFixtures(t, func() {
		imageRepoTests(t)
	})
}

func imageRepoTests(t *testing.T) {
	images_test_dir := "test/data/images/"
	images_test_excluded_dirs := []string{"whoami", "donation"}

	t.Run("#LOAD", func(t *testing.T) {
		t.Run("Loads all images recursively if directory exists", func(t *testing.T) {
			repo := filesystemrepo.NewImage(images_test_dir, []string{})

			loadedImages, err := repo.Load()

			assert.Len(t, loadedImages, 5)
			assert.NoError(t, err)
		})

		t.Run("Doesn't load images from excluded directories", func(t *testing.T) {
			repo := filesystemrepo.NewImage(images_test_dir, images_test_excluded_dirs)

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
			repo := filesystemrepo.NewImage(images_test_dir, images_test_excluded_dirs)

			filename := fmt.Sprintf("%s%s%s", images_test_dir, images_test_excluded_dirs[0], "/testFile.png")
			image, image_err := domain.NewImage(filename)
			assert.NoError(t, image_err)

			f, create_err := os.Create(filename)
			defer f.Close()
			assert.NoError(t, create_err)

			repo_err := repo.Delete(image)
			assert.NoError(t, repo_err)

			rm_err := os.Remove(filename)
			assert.NoError(t, rm_err)
		})

		t.Run("Deletes an image", func(t *testing.T) {
			repo := filesystemrepo.NewImage(images_test_dir, images_test_excluded_dirs)

			filename := fmt.Sprintf("%s%s", images_test_dir, "testFile.png")
			image, image_err := domain.NewImage(filename)
			assert.NoError(t, image_err)

			f, create_err := os.Create(filename)
			defer f.Close()
			assert.NoError(t, create_err)

			repo_err := repo.Delete(image)
			assert.NoError(t, repo_err)

			rm_err := os.Remove(filename)
			assert.Error(t, rm_err)
		})
	})

	t.Run("#CONVERT", func(t *testing.T) {
		t.Run("Doesn't convert images from excluded directories", func(t *testing.T) {
			repo := filesystemrepo.NewImage(images_test_dir, images_test_excluded_dirs)

			filename := fmt.Sprintf("%s%s%s", images_test_dir, images_test_excluded_dirs[0],
				"/avatar")
			jpegFilename := fmt.Sprintf("%s%s", filename, ".jpg")
			webpFilename := fmt.Sprintf("%s%s", filename, ".webp")
			image, image_err := domain.NewImage(jpegFilename)
			assert.NoError(t, image_err)

			f, create_err := os.Create(filename)
			defer f.Close()
			assert.NoError(t, create_err)

			repo_err := repo.ConvertToWebp([]domain.Image{image})
			assert.NoError(t, repo_err)

			rm_err := os.Remove(webpFilename)
			assert.Error(t, rm_err)
		})

		t.Run("Converts a PNG image to webp", func(t *testing.T) {
			repo := filesystemrepo.NewImage(images_test_dir, images_test_excluded_dirs)

			filename := fmt.Sprintf("%s%s", images_test_dir, "an_image")
			pngFilename := fmt.Sprintf("%s%s", filename, ".png")
			webpFilename := fmt.Sprintf("%s%s", filename, ".webp")

			image, image_err := domain.NewImage(pngFilename)
			assert.NoError(t, image_err)

			repo_err := repo.ConvertToWebp([]domain.Image{image})
			assert.NoError(t, repo_err)

			rm_err := os.Remove(webpFilename)
			assert.NoError(t, rm_err)
		})

		t.Run("Converts a JPEG image to webp", func(t *testing.T) {
			repo := filesystemrepo.NewImage(images_test_dir, images_test_excluded_dirs)

			filename := fmt.Sprintf("%s%s", images_test_dir, "another_image")
			jpegFilename := fmt.Sprintf("%s%s", filename, ".jpeg")
			webpFilename := fmt.Sprintf("%s%s", filename, ".webp")

			image, image_err := domain.NewImage(jpegFilename)
			assert.NoError(t, image_err)

			repo_err := repo.ConvertToWebp([]domain.Image{image})
			assert.NoError(t, repo_err)

			rm_err := os.Remove(webpFilename)
			assert.NoError(t, rm_err)
		})

		t.Run("Converts a GIF to webp", func(t *testing.T) {
			repo := filesystemrepo.NewImage(images_test_dir, images_test_excluded_dirs)

			filename := fmt.Sprintf("%s%s", images_test_dir, "a_gif")
			gifFilename := fmt.Sprintf("%s%s", filename, ".gif")
			webpFilename := fmt.Sprintf("%s%s", filename, ".webp")

			image, image_err := domain.NewImage(gifFilename)
			assert.NoError(t, image_err)

			repo_err := repo.ConvertToWebp([]domain.Image{image})
			assert.NoError(t, repo_err)

			rm_err := os.Remove(webpFilename)
			assert.NoError(t, rm_err)
		})
	})
}
