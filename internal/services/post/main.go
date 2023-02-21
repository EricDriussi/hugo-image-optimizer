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

func (s PostService) GetAllReferencedImagePaths() ([]string, error) {
	all_posts, err := s.Load()
	var images []string
	for _, post := range all_posts {
		images = append(images, post.GetCleanImageReferences()...)
	}
	return images, err
}

func (s PostService) Load() ([]domain.Post, error) {
	rawPosts, err := s.postRepository.Load()
	if err != nil {
		return nil, errors.New("Repository failed to load posts")
	}

	return s.buildPostsAllOrNothing(rawPosts)
}

func (s PostService) buildPostsAllOrNothing(rawPosts map[string][]byte) ([]domain.Post, error) {
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
