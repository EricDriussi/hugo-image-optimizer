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

func (p Post) GetPath() string {
	return p.path.Value()
}

func (p Post) GetCleanImageReferences() []string {
	var cleaned_image_references []string
	for _, image := range p.content.Images() {
		cleaned_image_references = append(cleaned_image_references, pathlib.Clean("/"+image))
	}
	return cleaned_image_references
}

func (p Post) GetFullContent() string {
	return string(p.content.Value())
}
