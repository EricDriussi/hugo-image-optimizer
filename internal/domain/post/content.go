package post

import (
	"fmt"
	"regexp"
)

type imageMatch = [][][]byte

type PostContent struct {
	content []byte
}

var (
	validExt             = "(jpg|png|jpeg|gif)"
	imagePath            = fmt.Sprintf("(.*\\.)%s", validExt)
	mdReference          = fmt.Sprintf("(!\\[.*\\])\\((%s)\\)", imagePath)
	frontMatterReference = fmt.Sprintf("(?m)^(image: )(%s)$", imagePath)
)

func NewContent(content []byte) PostContent {
	return PostContent{
		content: content,
	}
}

func (c PostContent) Images() []string {
	mdImages := c.imagesInBody()
	frontMatterImages := c.imagesInFrontMatter()
	allImages := append(mdImages, frontMatterImages...)
	return filterImagePaths(allImages)
}

func (c *PostContent) ChangeImgExtToWebp() {
	c.updateMdReferences()
	c.updateFrontMatterReferences()
}

func (c PostContent) Value() []byte {
	return c.content
}

func (c PostContent) imagesInBody() imageMatch {
	mdRegex := regexp.MustCompile(mdReference)
	return mdRegex.FindAllSubmatch(c.content, -1)
}

func (c PostContent) imagesInFrontMatter() imageMatch {
	frontMatterRegex := regexp.MustCompile(frontMatterReference)
	return frontMatterRegex.FindAllSubmatch(c.content, -1)
}

func filterImagePaths(imageMatches imageMatch) []string {
	paths := []string{}
	for _, match := range imageMatches {
		imagePathSubmatch := match[2]
		paths = append(paths, string(imagePathSubmatch))
	}
	return paths
}

func (c *PostContent) updateMdReferences() {
	mdRegex := regexp.MustCompile(mdReference)
	replacement := "${1}(${3}webp)"
	c.content = mdRegex.ReplaceAll(c.content, []byte(replacement))
}

func (c *PostContent) updateFrontMatterReferences() {
	frontMatterRegex := regexp.MustCompile(frontMatterReference)
	replacement := "${1}${3}webp"
	c.content = frontMatterRegex.ReplaceAll(c.content, []byte(replacement))
}
