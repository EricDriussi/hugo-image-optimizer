package image

import (
	"errors"
	"fmt"
	"path"
	"regexp"
)

type Extension struct {
	valid_extension string
}

func NewExtension(filepath string) (Extension, error) {
	extension, ok := validateExtension(filepath)
	if !ok {
		return Extension{}, errors.New(fmt.Sprintf("Invalid extension, ignoring: %s", filepath))
	}
	return Extension{valid_extension: extension}, nil
}

func validateExtension(filepath string) (string, bool) {
	extension := path.Ext(filepath)
	if len(extension) == 0 {
		return "", false
	}

	valid_extensions := regexp.MustCompile(".(jpg|png|jpeg|gif|webp)")
	if !valid_extensions.Match([]byte(extension)) {
		return "", false
	}
	return extension, true
}

func (n Extension) Value() string {
	return n.valid_extension
}
