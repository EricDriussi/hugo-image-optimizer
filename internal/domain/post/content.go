package post

import "regexp"

type PostContent struct {
	full_content     []byte
	image_references []byte
}

func NewContent(content []byte) PostContent {
	return PostContent{
		full_content:     content,
		image_references: extractImageReferences(content),
	}
}

func extractImageReferences(content []byte) []byte {
	md_images := mdReferences(content)
	front_matter_images := frontMatterReferences(content)
	all_images := append(md_images, front_matter_images...)
	return onlyImagePaths(all_images)
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

func (c PostContent) Value() []byte {
	return c.full_content
}

func (c PostContent) Images() []byte {
	return c.image_references
}
