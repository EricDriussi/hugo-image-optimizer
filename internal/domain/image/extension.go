package image

import (
	"errors"
	"fmt"
	"path"
	"regexp"
)

type Extension struct {
	extension string
}

func NewExtension(filepath string) (Extension, error) {
	extension, ok := validateExtension(filepath)
	if !ok {
		return Extension{}, errors.New(fmt.Sprintf("Invalid extension, ignoring: %s", filepath))
	}
	return Extension{extension: extension}, nil
}

func validateExtension(filepath string) (string, bool) {
	extension := path.Ext(filepath)
	isValid := isNotEmpty(extension) && isSupported(extension)
	return extension, isValid
}

func isNotEmpty(ext string) bool {
	return len(ext) > 0
}

func isSupported(ext string) bool {
	validExt := regexp.MustCompile(".(jpg|png|jpeg|gif)")
	return validExt.Match([]byte(ext))
}

func (n Extension) Value() string {
	return n.extension
}
