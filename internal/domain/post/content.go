package post

import (
	"fmt"
	"regexp"
)

type PostContent struct {
	full_content []byte
	// TODO. remove and leave Images() func?
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
	md_regex := regexp.MustCompile(mdRefRegex)
	return md_regex.FindAllSubmatch(text, -1)
}

func getFrontMatterReferencesMatchesIn(text []byte) [][][]byte {
	front_matter_regex := regexp.MustCompile(frontMatterRefRegex)
	return front_matter_regex.FindAllSubmatch(text, -1)
}

var (
	validExt            = "(jpg|png|jpeg|gif)"
	imagePath           = fmt.Sprintf("(.*\\.)%s", validExt)
	mdRefRegex          = fmt.Sprintf("!\\[.*\\]\\((%s)\\)", imagePath)
	frontMatterRefRegex = fmt.Sprintf("(?m)^image: (%s)$", imagePath)
)

func (c PostContent) Images() []string {
	var srting_paths []string
	bytes_paths := c.image_references
	for _, image_path := range bytes_paths {
		srting_paths = append(srting_paths, string(image_path))
	}
	return srting_paths
}

func (c *PostContent) UpdateImageReferences() {
	any_img_regex := regexp.MustCompile("(.*)(jpg|png|jpeg|gif)(.*)")
	webp_repl := []byte("${1}webp${3}")
	for k, image_ref := range c.image_references {
		c.image_references[k] = any_img_regex.ReplaceAll(image_ref, webp_repl)
	}
	c.full_content = any_img_regex.ReplaceAll(c.full_content, webp_repl)
}

func (c PostContent) Value() []byte {
	return c.full_content
}
