package post

import (
	"errors"
	"fmt"
	"path"
	"strings"
)

type Path struct {
	post_path string
}

func NewPath(filepath string) (Path, error) {
	emptyFilePath := strings.EqualFold(path.Base(filepath), ".")
	if emptyFilePath {
		return Path{}, errors.New(fmt.Sprintf("Invalid file path: %s", filepath))
	}
	return Path{post_path: filepath}, nil
}

func (n Path) Value() string {
	return n.post_path
}
