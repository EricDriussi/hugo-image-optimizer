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
	md_images := getMdReferencesFrom(content)
	front_matter_images := getFrontMatterReferencesFrom(content)
	all_images := append(md_images, front_matter_images...)
	return filterImagePaths(all_images)
}

func filterImagePaths(image_references [][][]byte) [][]byte {
	only_paths := [][]byte{}
	for _, ref := range image_references {
		image_path := ref[1]
		only_paths = append(only_paths, image_path)
	}
	return only_paths
}

func getMdReferencesFrom(text []byte) [][][]byte {
	md_regex := regexp.MustCompile("!\\[.*\\]\\((.*\\.(jpg|png|jpeg|gif|webp))\\)")
	return md_regex.FindAllSubmatch(text, -1)
}

func getFrontMatterReferencesFrom(text []byte) [][][]byte {
	front_matter_regex := regexp.MustCompile("(?m)^image: (.*\\.(jpg|png|jpeg|webp))$")
	return front_matter_regex.FindAllSubmatch(text, -1)
}

func (c PostContent) Value() []byte {
	return c.full_content
}

func (c PostContent) Images() []string {
	var image_paths_as_strings []string
	image_paths_as_bytes := c.image_references
	for _, image_path := range image_paths_as_bytes {
		image_paths_as_strings = append(image_paths_as_strings, string(image_path))
	}
	return image_paths_as_strings
}
