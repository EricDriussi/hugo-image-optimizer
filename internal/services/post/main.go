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

func (s PostService) GetImagesInPosts() ([]string, error) {
	all_posts, err := s.loadPosts()
	var images []string
	for _, post := range all_posts {
		images = append(images, post.GetCleanImageReferences()...)
	}
	return images, err
}

func (s PostService) Load() ([]domain.Post, error) {
	return s.loadPosts()
}

func (s PostService) loadPosts() ([]domain.Post, error) {
	rawPosts, err := s.postRepository.Load()
	if err != nil {
		return nil, errors.New("Repository failed to load posts")
	}

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
