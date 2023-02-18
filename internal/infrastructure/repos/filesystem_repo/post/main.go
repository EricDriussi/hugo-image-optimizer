package filesystemrepo

import (
	"os"
	"path/filepath"
)

type fsrepo struct {
	posts_dir string
}

func NewPost(posts_dir string) fsrepo {
	return fsrepo{
		posts_dir: posts_dir,
	}
}

func (r fsrepo) Load() (map[string][]byte, error) {
	posts := map[string][]byte{}
	err := filepath.Walk(r.posts_dir, loadPostsInto(posts))
	return posts, err
}

func loadPostsInto(posts map[string][]byte) filepath.WalkFunc {
	return func(filepath string, file os.FileInfo, error error) error {
		if error != nil {
			return error
		}
		if file.IsDir() {
			return nil
		}

		single_post, err := os.ReadFile(filepath)
		posts[filepath] = single_post
		return err
	}
}
