package services_test

import (
	"errors"
	"testing"

	"github.com/EricDriussi/hugo-image-optimizer/internal/services/post"
	"github.com/EricDriussi/hugo-image-optimizer/test/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_PostService(t *testing.T) {
	t.Run("Loads", func(t *testing.T) {
		t.Run("a post", func(t *testing.T) {
			content := "Some random content"
			path := "a/file/path/post.md"
			posts := map[string][]byte{path: []byte(content)}
			postRepositoryMock := new(mocks.PostRepository)
			postRepositoryMock.On("Load").Return(posts, nil)

			postService := services.NewPost(postRepositoryMock)
			loadedPosts, err := postService.LoadAll()
			postRepositoryMock.AssertExpectations(t)

			assert.Len(t, loadedPosts, 1)
			assert.Equal(t, path, loadedPosts[0].GetPath())
			assert.Equal(t, content, loadedPosts[0].GetFullContent())
			assert.NoError(t, err)
		})

		t.Run("multiple posts", func(t *testing.T) {
			contentOne := "Some random content"
			pathOne := "a/file/path/post.md"
			contentTwo := "Some more random content"
			pathTwo := "a/file/path/to/another/post.md"
			posts := map[string][]byte{pathOne: []byte(contentOne), pathTwo: []byte(contentTwo)}
			postRepositoryMock := new(mocks.PostRepository)
			postRepositoryMock.On("Load").Return(posts, nil)

			postService := services.NewPost(postRepositoryMock)
			loadedPosts, err := postService.LoadAll()
			postRepositoryMock.AssertExpectations(t)

			assert.Len(t, loadedPosts, 2)
			assert.Equal(t, pathOne, loadedPosts[0].GetPath())
			assert.Equal(t, contentOne, loadedPosts[0].GetFullContent())
			assert.Equal(t, pathTwo, loadedPosts[1].GetPath())
			assert.Equal(t, contentTwo, loadedPosts[1].GetFullContent())
			assert.NoError(t, err)
		})
	})

	t.Run("Doesn't partially load posts", func(t *testing.T) {
		t.Run("if the repository erros out", func(t *testing.T) {
			anError := errors.New("Something went wrong")
			postRepositoryMock := new(mocks.PostRepository)
			postRepositoryMock.On("Load").Return(nil, anError)

			postService := services.NewPost(postRepositoryMock)
			posts, err := postService.LoadAll()
			postRepositoryMock.AssertExpectations(t)

			assert.Nil(t, posts)
			assert.Error(t, err)
		})

		t.Run("if one post erros out", func(t *testing.T) {
			contentOne := "Some random content"
			pathOne := ""
			contentTwo := "Some more random content"
			pathTwo := "a/file/path/to/another/post.md"
			posts := map[string][]byte{pathOne: []byte(contentOne), pathTwo: []byte(contentTwo)}
			postRepositoryMock := new(mocks.PostRepository)
			postRepositoryMock.On("Load").Return(posts, nil)

			postService := services.NewPost(postRepositoryMock)
			loadedPosts, err := postService.LoadAll()
			postRepositoryMock.AssertExpectations(t)

			assert.Nil(t, loadedPosts)
			assert.Error(t, err)
		})
	})
}
