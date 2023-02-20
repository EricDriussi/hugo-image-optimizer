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
	image_pathOne := "../path/src.png"
	image_referenceOne := fmt.Sprintf("![image](%s)", image_pathOne)
	contentOne := fmt.Sprintf(`line 1
					line %s 2
					line 4`,
		image_referenceOne)

	pathTwo := "a/file/path/to/another/post.md"
	image_pathTwo := "../path/src2.jpg"
	image_referenceTwo := fmt.Sprintf("![image](%s)", image_pathTwo)
	contentTwo := fmt.Sprintf(`line 1
					%s
					line 4`,
		image_referenceTwo)

	t.Run("Extracts all images", func(t *testing.T) {
		posts := map[string][]byte{pathOne: []byte(contentOne), pathTwo: []byte(contentTwo)}
		postRepositoryMock := new(mocks.PostRepository)
		postRepositoryMock.On("Load").Return(posts, nil)

		postService := services.NewPost(postRepositoryMock)
		imageReferences, err := postService.GetImagesInPosts()
		postRepositoryMock.AssertExpectations(t)

		assert.Len(t, imageReferences, 2)
		assert.Contains(t, imageReferences, image_pathOne)
		assert.Contains(t, imageReferences, image_pathTwo)
		assert.NoError(t, err)

		t.Run("No partial loading", func(t *testing.T) {
			t.Run("if the repository erros out", func(t *testing.T) {
				anError := errors.New("Something went wrong")
				postRepositoryMock := new(mocks.PostRepository)
				postRepositoryMock.On("Load").Return(nil, anError)

				postService := services.NewPost(postRepositoryMock)
				posts, err := postService.GetImagesInPosts()
				postRepositoryMock.AssertExpectations(t)

				assert.Nil(t, posts)
				assert.Error(t, err)
			})

			t.Run("if one post erros out", func(t *testing.T) {
				pathOne := ""
				posts := map[string][]byte{pathOne: []byte(contentOne), pathTwo: []byte(contentTwo)}
				postRepositoryMock := new(mocks.PostRepository)
				postRepositoryMock.On("Load").Return(posts, nil)

				postService := services.NewPost(postRepositoryMock)
				loadedPosts, err := postService.GetImagesInPosts()
				postRepositoryMock.AssertExpectations(t)

				assert.Nil(t, loadedPosts)
				assert.Error(t, err)
			})
		})
	})

	t.Run("Loads posts", func(t *testing.T) {
		posts := map[string][]byte{pathOne: []byte(contentOne), pathTwo: []byte(contentTwo)}
		postRepositoryMock := new(mocks.PostRepository)
		postRepositoryMock.On("Load").Return(posts, nil)

		postService := services.NewPost(postRepositoryMock)
		loadedPosts, err := postService.Load()
		postRepositoryMock.AssertExpectations(t)

		assert.Len(t, loadedPosts, 2)
		assert.Equal(t, pathOne, loadedPosts[0].GetPath())
		assert.Equal(t, contentOne, loadedPosts[0].GetFullContent())
		assert.Equal(t, pathTwo, loadedPosts[1].GetPath())
		assert.Equal(t, contentTwo, loadedPosts[1].GetFullContent())
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
			})

			t.Run("if one post erros out", func(t *testing.T) {
				pathOne := ""
				posts := map[string][]byte{pathOne: []byte(contentOne), pathTwo: []byte(contentTwo)}
				postRepositoryMock := new(mocks.PostRepository)
				postRepositoryMock.On("Load").Return(posts, nil)

				postService := services.NewPost(postRepositoryMock)
				loadedPosts, err := postService.Load()
				postRepositoryMock.AssertExpectations(t)

				assert.Nil(t, loadedPosts)
				assert.Error(t, err)
			})
		})
	})
}
