package filterService

import (
	"bytes"
	"optimize/internal/util"
	"regexp"
)

func Images_present_in(text *[]byte) []byte {
	var result bytes.Buffer
	result.Write(md_image_links(text))
	result.Write(images_in_front_matter(text))
	return result.Bytes()
}

func md_image_links(text *[]byte) []byte {
	md_image_links := regexp.
		MustCompile("!\\[.*\\]\\((.*)\\)").
		FindAllSubmatch(*text, -1)
	only_image_paths := []byte{}
	for _, full_match := range md_image_links {
		only_image_paths = append(only_image_paths, full_match[1]...)
	}
	return only_image_paths
}

func images_in_front_matter(text *[]byte) []byte {
	md_image_links := regexp.
		MustCompile("(?m)^image: (.*\\.(jpg|png|jpeg))$").
		FindAllSubmatch(*text, -1)
	only_image_paths := []byte{}
	for _, full_match := range md_image_links {
		only_image_paths = append(only_image_paths, full_match[1]...)
	}
	return only_image_paths
}

func Image_paths_being_referenced(image_files_list []string, md_image_list *[]byte) []string {
	used_images := []string{}
	for _, image_file := range image_files_list {
		image_is_referenced := util.ByteArrayContainsString(image_file, md_image_list)
		if image_is_referenced {
			used_images = append(used_images, image_file)
		}
	}
	return used_images
}

func Unused_image_paths(image_files_list []string, md_image_list *[]byte) []string {
	unused_images := []string{}
	for _, image_file := range image_files_list {
		image_is_referenced := util.ByteArrayContainsString(image_file, md_image_list)
		if !image_is_referenced {
			unused_images = append(unused_images, image_file)
		}
	}
	return unused_images
}
