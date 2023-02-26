package services_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/EricDriussi/hugo-image-optimizer/internal/domain"
	services "github.com/EricDriussi/hugo-image-optimizer/internal/services/post"
	"github.com/EricDriussi/hugo-image-optimizer/test/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_PostService(t *testing.T) {
	pathOne := "a/file/path/post.md"
	imageReferenceOne := "![image](../path/src.png)"
	contentOne := fmt.Sprintf(`line 1
					line %s 2
					line 4`,
		imageReferenceOne)

	pathTwo := "a/file/path/to/another/post.md"
	imageReferenceTwo := "![image](../path/src2.jpg)"
	contentTwo := fmt.Sprintf(`line 1
					%s
					line 4`,
		imageReferenceTwo)

	posts := map[string][]byte{pathOne: []byte(contentOne), pathTwo: []byte(contentTwo)}
	anError := errors.New("Something went wrong")

	t.Run("Extracts all images", func(t *testing.T) {
		postRepositoryMock := new(mocks.PostRepository)
		postRepositoryMock.On("Load").Return(posts, nil)

		postService := services.NewPost(postRepositoryMock)
		imageReferences, err := postService.AllReferencedImagePaths()
		postRepositoryMock.AssertExpectations(t)

		assert.Len(t, imageReferences, 2)
		assert.NoError(t, err)
	})

	t.Run("Loads posts", func(t *testing.T) {
		postRepositoryMock := new(mocks.PostRepository)
		postRepositoryMock.On("Load").Return(posts, nil)

		postService := services.NewPost(postRepositoryMock)
		loadedPosts, err := postService.Load()
		postRepositoryMock.AssertExpectations(t)

		assert.Len(t, loadedPosts, 2)
		assert.NoError(t, err)

		t.Run("No partial loading", func(t *testing.T) {
			t.Run("if the repository erros out", func(t *testing.T) {
				postRepositoryMock := new(mocks.PostRepository)
				postRepositoryMock.On("Load").Return(nil, anError)

				postService := services.NewPost(postRepositoryMock)
				posts, err := postService.Load()
				postRepositoryMock.AssertExpectations(t)

				assert.Nil(t, posts)
				assert.Error(t, err)
				assert.ErrorContains(t, err, "Couldn't load posts :(")
			})

			t.Run("if one post erros out", func(t *testing.T) {
				posts := map[string][]byte{"": []byte(contentOne), pathTwo: []byte(contentTwo)}
				postRepositoryMock := new(mocks.PostRepository)
				postRepositoryMock.On("Load").Return(posts, nil)

				postService := services.NewPost(postRepositoryMock)
				loadedPosts, err := postService.Load()
				postRepositoryMock.AssertExpectations(t)

				assert.Nil(t, loadedPosts)
				assert.Error(t, err)
				assert.ErrorContains(t, err, "[ABORTING] Couldn't build post")
			})
		})
	})

	t.Run("Updates", func(t *testing.T) {
		imageReferenceOne := "![image](../path/src.webp)"
		contentOne := fmt.Sprintf(`line 1
					line %s 2
					line 4`,
			imageReferenceOne)

		imageReferenceTwo := "![image](../path/src2.webp)"
		contentTwo := fmt.Sprintf(`line 1
					%s
					line 4`,
			imageReferenceTwo)

		postOne, errOne := domain.NewPost(pathOne, []byte(contentOne))
		assert.NoError(t, errOne)
		postTwo, errTwo := domain.NewPost(pathTwo, []byte(contentTwo))
		assert.NoError(t, errTwo)

		t.Run("all image references", func(t *testing.T) {
			postRepositoryMock := new(mocks.PostRepository)
			postRepositoryMock.On("Load").Return(posts, nil)
			postRepositoryMock.On("Write", postOne).Return(nil)
			postRepositoryMock.On("Write", postTwo).Return(nil)

			postService := services.NewPost(postRepositoryMock)
			err := postService.UpdateAllImageReferences()
			postRepositoryMock.AssertExpectations(t)

			assert.NoError(t, err)
		})

		t.Run("stopping if one update fails", func(t *testing.T) {
			postRepositoryMock := new(mocks.PostRepository)
			postRepositoryMock.On("Load").Return(posts, nil)
			postRepositoryMock.On("Write", postOne).Return(anError)

			postService := services.NewPost(postRepositoryMock)
			err := postService.UpdateAllImageReferences()
			postRepositoryMock.AssertExpectations(t)

			assert.Error(t, err)
			assert.ErrorContains(t, err, "Couldn't update posts :(")
		})
	})
}
