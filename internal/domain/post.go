package domain

import (
	"github.com/EricDriussi/hugo-image-optimizer/internal/domain/post"
)

type PostRepository interface {
	Load() (map[string][]byte, error)
}

type Post struct {
	path    post.Path
	content post.PostContent
}

func NewPost(filepath string, rawContent []byte) (Post, error) {
	path, err := post.NewPath(filepath)
	if err != nil {
		return Post{}, err
	}
	content := post.NewContent(rawContent)

	newPost := Post{
		path:    path,
		content: content,
	}

	return newPost, nil
}

func (p Post) GetPath() string {
	return p.path.Value()
}

func (p Post) GetReferencedImages() []string {
	return p.content.Images()
}

func (p Post) GetFullContent() string {
	return string(p.content.Value())
}
