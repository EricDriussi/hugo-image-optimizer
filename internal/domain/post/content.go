package post

import (
	"regexp"
)

type PostContent struct {
	full_content     []byte
	image_references [][]byte
}

func NewContent(content []byte) PostContent {
	return PostContent{
		full_content:     content,
		image_references: extractImageReferences(content),
	}
}

func extractImageReferences(content []byte) [][]byte {
	md_images := getMdReferencesMatchesIn(content)
	front_matter_images := getFrontMatterReferencesMatchesIn(content)
	all_images := append(md_images, front_matter_images...)
	return onlyImagePaths(all_images)
}

func onlyImagePaths(image_references [][][]byte) [][]byte {
	only_paths := [][]byte{}
	for _, ref := range image_references {
		image_path := ref[1]
		only_paths = append(only_paths, image_path)
	}
	return only_paths
}

func getMdReferencesMatchesIn(text []byte) [][][]byte {
	md_regex := regexp.MustCompile("!\\[.*\\]\\((.*\\.(jpg|png|jpeg|gif))\\)")
	return md_regex.FindAllSubmatch(text, -1)
}

func getFrontMatterReferencesMatchesIn(text []byte) [][][]byte {
	front_matter_regex := regexp.MustCompile("(?m)^image: (.*\\.(jpg|png|jpeg))$")
	return front_matter_regex.FindAllSubmatch(text, -1)
}

func (c PostContent) Images() []string {
	var srting_paths []string
	bytes_paths := c.image_references
	for _, image_path := range bytes_paths {
		srting_paths = append(srting_paths, string(image_path))
	}
	return srting_paths
}

func (c PostContent) Value() []byte {
	return c.full_content
}
