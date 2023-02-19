package domain

import (
	"errors"
	"fmt"

	"github.com/EricDriussi/hugo-image-optimizer/internal/domain/post"
)

type PostRepository interface {
	Load() (map[string][]byte, error)
}

type Post struct {
	filename post.FileName
	path     string
	content  post.PostContent
	images   []byte
}

func NewPost(filepath string, rawContent []byte) (Post, error) {
	name, err := post.NewName(filepath)
	if err != nil {
		return Post{}, errors.New(fmt.Sprintf("Invalid filepath: %s", filepath))
	}

	content := post.NewContent(rawContent)

	newPost := Post{
		filename: name,
		path:     filepath,
		content:  content,
		images:   content.Images(),
	}

	return newPost, nil
}

func (p Post) GetFilename() string {
	return p.filename.Value()
}

func (p Post) GetPath() string {
	return p.path
}

func (p Post) GetReferencedImages() string {
	return string(p.images)
}

func (p Post) GetFullContent() string {
	return string(p.content.Value())
}
