package filesystemrepo

import (
	"os"
	"path/filepath"
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
	err := filepath.Walk(r.postsDir, r.loadPostsInto(posts))
	return posts, err
}

func (r fsrepo) loadPostsInto(posts map[string][]byte) filepath.WalkFunc {
	return func(filepath string, file os.FileInfo, error error) error {
		if error != nil {
			return error
		}
		if file.IsDir() {
			return nil
		}

		singlePost, err := os.ReadFile(filepath)
		posts[filepath] = singlePost
		return err
	}
}
