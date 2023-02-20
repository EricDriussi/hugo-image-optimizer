package domain

import (
	"errors"
	"fmt"
)

type Path struct {
	valid_path string
}

func NewPath(filepath string) (Path, error) {
	if len(filepath) == 0 {
		return Path{}, errors.New(fmt.Sprintf("File path can't be empty: %s", filepath))
	}
	return Path{valid_path: filepath}, nil
}

func (n Path) Value() string {
	return n.valid_path
}
