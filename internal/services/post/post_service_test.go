package services_test

import (
	"errors"
	"fmt"
	"testing"

	services "github.com/EricDriussi/hugo-image-optimizer/internal/services/post"
	"github.com/EricDriussi/hugo-image-optimizer/test/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_PostService(t *testing.T) {
	pathOne := "a/file/path/post.md"
	image_referenceOne := "![image](../path/src.png)"
	contentOne := fmt.Sprintf(`line 1
					line %s 2
					line 4`,
		image_referenceOne)

	pathTwo := "a/file/path/to/another/post.md"
	image_referenceTwo := "![image](../path/src2.jpg)"
	contentTwo := fmt.Sprintf(`line 1
					%s
					line 4`,
		image_referenceTwo)

	posts := map[string][]byte{pathOne: []byte(contentOne), pathTwo: []byte(contentTwo)}

	t.Run("Extracts all images", func(t *testing.T) {
		postRepositoryMock := new(mocks.PostRepository)
		postRepositoryMock.On("Load").Return(posts, nil)

		postService := services.NewPost(postRepositoryMock)
		imageReferences, err := postService.GetAllReferencedImagePaths()
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
				anError := errors.New("Something went wrong")
				postRepositoryMock := new(mocks.PostRepository)
				postRepositoryMock.On("Load").Return(nil, anError)

				postService := services.NewPost(postRepositoryMock)
				posts, err := postService.Load()
				postRepositoryMock.AssertExpectations(t)

				assert.Nil(t, posts)
				assert.Error(t, err)
				assert.ErrorContains(t, err, "Repository failed to load posts")
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
}
