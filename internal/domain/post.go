package domain

import (
	pathlib "path"

	"github.com/EricDriussi/hugo-image-optimizer/internal/domain/post"
)

type PostRepository interface {
	Load() (map[string][]byte, error)
}

type Post struct {
	path    Path
	content post.PostContent
}

func NewPost(filepath string, rawContent []byte) (Post, error) {
	path, err := NewPath(filepath)
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

func (p Post) Path() string {
	return p.path.Value()
}

func (p Post) ReferencedImagePaths() []string {
	var cleanedImageReferences []string
	for _, image := range p.content.Images() {
		cleanedImageReferences = append(cleanedImageReferences, pathlib.Clean("/"+image))
	}
	return cleanedImageReferences
}

func (p Post) Content() string {
	return string(p.content.Value())
}
