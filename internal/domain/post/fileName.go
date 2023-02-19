package post

import (
	"errors"
	"path"
	"strings"
)

type FileName struct {
	filename string
}

func NewName(filepath string) (FileName, error) {
	name := path.Base(filepath)
	emptyFilePath := strings.EqualFold(name, ".")
	if emptyFilePath {
		return FileName{}, errors.New("Could't build file name")
	}
	return FileName{filename: name}, nil
}

func (n FileName) Value() string {
	return n.filename
}
