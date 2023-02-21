package filesystemrepo_test

import (
	"os"
	"path"
	"runtime"
	"testing"

	filesystemrepo "github.com/EricDriussi/hugo-image-optimizer/internal/infrastructure/repos/filesystem_repo/image"
	"github.com/stretchr/testify/assert"
)

func Test_ImageRepository_Load(t *testing.T) {
	setCWDToProjectRoot()
	validDir := "test/data/images/"

	t.Run("Loads all images recursively if directory exists", func(t *testing.T) {
		repo := filesystemrepo.NewImage(validDir, []string{})

		loadedImages, err := repo.Load()

		assert.Len(t, loadedImages, 5)
		assert.NoError(t, err)
	})

	t.Run("Doesn't load images from excluded directories", func(t *testing.T) {
		excludedDirs := []string{"whoami", "donation"}
		repo := filesystemrepo.NewImage(validDir, excludedDirs)

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
}

func setCWDToProjectRoot() {
	_, filename, _, _ := runtime.Caller(0)
	project_root := path.Join(path.Dir(filename), "../../../../..")
	if err := os.Chdir(project_root); err != nil {
		panic(err)
	}
}
