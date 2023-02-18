package domain

type PostRepository interface {
	Load() (map[string][]byte, error)
}

type Post struct {
	filename string
	path     string
	content  []byte
	images   []byte
}

func NewPost(path string, content []byte) Post {
	return Post{path: path, content: content}
}

func (p Post) GetFilename() string {
	return p.filename
}

func (p Post) GetPath() string {
	return p.path
}

func (p Post) GetReferencedImages() []byte {
	return p.images
}

func (p Post) GetFullContent() string {
	return string(p.content)
}
