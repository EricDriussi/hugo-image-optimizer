package services

import (
	"errors"
	"fmt"

	"github.com/EricDriussi/hugo-image-optimizer/internal/domain"
)

type PostService struct {
	postRepository domain.PostRepository
}

func NewPost(postRepository domain.PostRepository) PostService {
	return PostService{
		postRepository: postRepository,
	}
}

func (s PostService) AllReferencedImagePaths() ([]string, error) {
	allPosts, err := s.Load()
	var images []string
	for _, post := range allPosts {
		images = append(images, post.ReferencedImagePaths()...)
	}
	return images, err
}

func (s PostService) UpdateAllImageReferences() error {
	allPosts, loadErr := s.Load()
	if loadErr != nil {
		return loadErr
	}
	for _, post := range allPosts {
		post.UpdateImageReferences()
		if writeErr := s.postRepository.Write(post); writeErr != nil {
			return errors.New("Couldn't update posts :(")
		}
	}
	return nil
}

func (s PostService) Load() ([]domain.Post, error) {
	rawPosts, err := s.postRepository.Load()
	if err != nil {
		return nil, errors.New("Couldn't load posts :(")
	}

	return buildPostsAllOrNothing(rawPosts)
}

func buildPostsAllOrNothing(rawPosts map[string][]byte) ([]domain.Post, error) {
	var posts []domain.Post
	for path, content := range rawPosts {
		post, err := domain.NewPost(path, content)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("[ABORTING] Couldn't build post: %s", path))
		}
		posts = append(posts, post)
	}
	return posts, nil
}
