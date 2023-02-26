package post

import (
	"fmt"
	"regexp"
)

type imageMatch = [][][]byte

type PostContent struct {
	full_content []byte
}

var (
	validExt             = "(jpg|png|jpeg|gif)"
	imagePath            = fmt.Sprintf("(.*\\.)%s", validExt)
	mdReference          = fmt.Sprintf("(!\\[.*\\])\\((%s)\\)", imagePath)
	frontMatterReference = fmt.Sprintf("(?m)^(image: )(%s)$", imagePath)
)

func NewContent(content []byte) PostContent {
	return PostContent{
		full_content: content,
	}
}

func (c PostContent) Images() []string {
	md_images := c.getImagesInMdBody()
	front_matter_images := c.getImagesInFrontMatter()
	all_images := append(md_images, front_matter_images...)
	return filterImagePaths(all_images)
}

func (c *PostContent) UpdateImageReferences() {
	c.updateMdReferences()
	c.updateFrontMatterReferences()
}

func (c PostContent) Value() []byte {
	return c.full_content
}

func (c PostContent) getImagesInMdBody() imageMatch {
	md_regex := regexp.MustCompile(mdReference)
	return md_regex.FindAllSubmatch(c.full_content, -1)
}

func (c PostContent) getImagesInFrontMatter() imageMatch {
	front_matter_regex := regexp.MustCompile(frontMatterReference)
	return front_matter_regex.FindAllSubmatch(c.full_content, -1)
}

func filterImagePaths(image_matches imageMatch) []string {
	paths := []string{}
	for _, match := range image_matches {
		image_path_submatch := match[2]
		paths = append(paths, string(image_path_submatch))
	}
	return paths
}

func (c *PostContent) updateMdReferences() {
	md_regex := regexp.MustCompile(mdReference)
	replacement := "${1}(${3}webp)"
	c.full_content = md_regex.ReplaceAll(c.full_content, []byte(replacement))
}

func (c *PostContent) updateFrontMatterReferences() {
	front_matter_regex := regexp.MustCompile(frontMatterReference)
	replacement := "${1}${3}webp"
	c.full_content = front_matter_regex.ReplaceAll(c.full_content, []byte(replacement))
}
