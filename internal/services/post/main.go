package services

import (
	"errors"

	"github.com/EricDriussi/hugo-image-optimizer/internal/domain/post"
)

type PostService struct {
	postRepository domain.PostRepository
}

func NewPost(postRepository domain.PostRepository) PostService {
	return PostService{
		postRepository: postRepository,
	}
}

func (s PostService) LoadAll() ([]domain.Post, error) {
	posts, err := s.postRepository.Load()
	if err != nil {
		return nil, errors.New("Repository failed to load posts")
	}

	var domainPosts []domain.Post
	for path, content := range posts {
		domainPosts = append(domainPosts, domain.NewPost(path, content))
	}

	return domainPosts, nil
}
