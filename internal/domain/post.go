package domain

import (
	"errors"
	"fmt"
	"path"
	"regexp"
	"strings"
)

type PostRepository interface {
	Load() (map[string][]byte, error)
}

type Post struct {
	filename string
	path     string
	content  []byte
	images   []byte
}

func NewPost(filepath string, content []byte) (Post, error) {
	name := path.Base(filepath)
	if strings.EqualFold(name, ".") {
		return Post{}, errors.New(fmt.Sprintf("Invalid filepath: %s", filepath))
	}

	newPost := Post{
		filename: name,
		path:     filepath,
		content:  content,
	}
	newPost.images = newPost.extractImageReferences()

	return newPost, nil
}

func (p Post) extractImageReferences() []byte {
	md_images := mdReferences(p.content)
	front_matter_images := frontMatterReferences(p.content)
	all_images := append(md_images, front_matter_images...)
	return onlyImagePaths(all_images)
}

func (p Post) GetFilename() string {
	return p.filename
}

func (p Post) GetPath() string {
	return p.path
}

func (p Post) GetReferencedImages() string {
	return string(p.images)
}

func (p Post) GetFullContent() string {
	return string(p.content)
}

func onlyImagePaths(matches [][][]byte) []byte {
	only_image_paths := []byte{}
	for _, full_match := range matches {
		only_image_paths = append(only_image_paths, full_match[1]...)
	}
	return only_image_paths
}

func mdReferences(text []byte) [][][]byte {
	md_regex := regexp.MustCompile("!\\[.*\\]\\((.*\\.(jpg|png|jpeg|gif|webp))\\)")
	return md_regex.FindAllSubmatch(text, -1)
}

func frontMatterReferences(text []byte) [][][]byte {
	front_matter_regex := regexp.MustCompile("(?m)^image: (.*\\.(jpg|png|jpeg|webp))$")
	return front_matter_regex.FindAllSubmatch(text, -1)
}
