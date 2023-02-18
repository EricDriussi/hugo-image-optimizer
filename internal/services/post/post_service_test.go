package services_test

import (
	"errors"
	"testing"

	"github.com/EricDriussi/hugo-image-optimizer/internal/services/post"
	"github.com/EricDriussi/hugo-image-optimizer/test/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_PostService(t *testing.T) {
	t.Run("Loads a post", func(t *testing.T) {
		content := "Some random content"
		path := "a/file/path/post.md"
		posts := map[string][]byte{path: []byte(content)}
		postRepositoryMock := new(mocks.PostRepository)
		postRepositoryMock.On("Load").Return(posts, nil)

		postService := services.NewPost(postRepositoryMock)
		loadedPosts, err := postService.LoadAll()
		postRepositoryMock.AssertExpectations(t)

		assert.Len(t, loadedPosts, 1)
		assert.Contains(t, loadedPosts[0].GetPath(), path)
		assert.Contains(t, loadedPosts[0].GetFullContent(), content)
		assert.NoError(t, err)
	})

	t.Run("Loads multiple posts", func(t *testing.T) {
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
		assert.Contains(t, loadedPosts[0].GetPath(), pathOne)
		assert.Contains(t, loadedPosts[0].GetFullContent(), contentOne)
		assert.Contains(t, loadedPosts[1].GetPath(), pathTwo)
		assert.Contains(t, loadedPosts[1].GetFullContent(), contentTwo)
		assert.NoError(t, err)
	})

	t.Run("Doesn't partially load posts if the repository erros out", func(t *testing.T) {
		anError := errors.New("Something went wrong")
		postRepositoryMock := new(mocks.PostRepository)
		postRepositoryMock.On("Load").Return(nil, anError)

		postService := services.NewPost(postRepositoryMock)
		posts, err := postService.LoadAll()
		postRepositoryMock.AssertExpectations(t)

		assert.Nil(t, posts)
		assert.Error(t, err)
	})
}
