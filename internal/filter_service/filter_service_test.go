package filterService_test

import (
	filterService "github.com/EricDriussi/hugo-image-optimizer/internal/filter_service"
	"github.com/EricDriussi/hugo-image-optimizer/internal/util"
	"testing"
)

func TestFiltersAllMDImagesFromPosts(t *testing.T) {
	all_posts := []byte(posts_text_fixture)
	images_in_posts := filterService.Images_present_in(&all_posts)

	doesNotContainExpectedImages := !util.ByteArrayContainsString("a_gif.gif", &images_in_posts) || !util.ByteArrayContainsString("another_image.jpeg", &images_in_posts) || !util.ByteArrayContainsString("hermit.jpg", &images_in_posts)

	if doesNotContainExpectedImages {
		t.Fail()
	}
}

func TestFiltersImagesUsedInPosts(t *testing.T) {
	all_posts := []byte(posts_text_fixture)
	images_in_use := filterService.Image_paths_being_referenced(image_files_fixture, &all_posts)

	unusedImageIsPresent := util.StringIsInArray(images_in_use, "an_image.png")
	if unusedImageIsPresent {
		t.Fail()
	}
}

func TestFiltersUnusedImages(t *testing.T) {
	all_posts := []byte(posts_text_fixture)
	unused_images := filterService.Unused_image_paths(image_files_fixture, &all_posts)

	unusedImageIsNotPresent := !util.StringIsInArray(unused_images, "an_image.png")
	if len(unused_images) > 1 || unusedImageIsNotPresent {
		t.Fail()
	}
}

var image_files_fixture = []string{"a_gif.gif", "an_image.png", "another_image.jpeg", "hermit.jpg", "lie.png"}

var posts_text_fixture = `
Two
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Duis at tellus at urna condimentum mattis. Sit amet nulla facilisi morbi tempus iaculis urna.
![image](../images/a_gif.gif)

---
image: images/lie.png
---

Three
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Duis at tellus at urna condimentum mattis. Sit amet nulla facilisi morbi tempus iaculis urna.
![image](../images/another_image.jpeg)

---
image: images/hermit.jpg
---

One
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Duis at tellus at urna condimentum mattis. Sit amet nulla facilisi morbi tempus iaculis urna.
`
