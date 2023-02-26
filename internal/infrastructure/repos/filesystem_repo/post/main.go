package filesystemrepo

import (
	"os"
	"path/filepath"

	"github.com/EricDriussi/hugo-image-optimizer/internal/domain"
)

type fsrepo struct {
	postsDir string
}

func NewPost(postsDir string) fsrepo {
	return fsrepo{
		postsDir: postsDir,
	}
}

func (r fsrepo) Load() (map[string][]byte, error) {
	posts := map[string][]byte{}
	loadTask := func(path string) error {
		singlePost, err := os.ReadFile(path)
		posts[path] = singlePost
		return err
	}

	err := filepath.Walk(r.parsePostsDirDoing(loadTask))
	return posts, err
}

func (r fsrepo) Write(post domain.Post) error {
	return os.WriteFile(post.Path(), []byte(post.Content()), os.ModePerm)
}

func (r fsrepo) parsePostsDirDoing(toDoTask func(string) error) (string, filepath.WalkFunc) {
	return r.postsDir, func(path string, file os.FileInfo, error error) error {
		if error != nil {
			return error
		}
		if file.IsDir() {
			return nil
		}

		return toDoTask(path)
	}
}
