package filterService

import (
	"bytes"
	"optimize/internal/util"
	"regexp"
)

func MD_images_present_in(text *[]byte) []byte {
	var result bytes.Buffer
	r, _ := regexp.Compile("!\\[.*\\]\\(.*\\)")
	matches := r.FindAll(*text, -1)
	for _, match := range matches {
		result.Write(match)
	}
	return result.Bytes()
}

func Images_being_referenced(image_files_list []string, md_image_list *[]byte) []string {
	used_images := []string{}
	for _, image_file := range image_files_list {
		image_is_referenced := util.ByteArrayContainsString(image_file, md_image_list)
		if image_is_referenced {
			used_images = append(used_images, image_file)
		}
	}
	return used_images
}

func Unused_images(image_files_list []string, md_image_list *[]byte) []string {
	unused_images := []string{}
	for _, image_file := range image_files_list {
		image_is_referenced := util.ByteArrayContainsString(image_file, md_image_list)
		if !image_is_referenced {
			unused_images = append(unused_images, image_file)
		}
	}
	return unused_images
}
