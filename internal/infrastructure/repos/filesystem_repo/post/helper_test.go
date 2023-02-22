package filesystemrepo_test

import (
	"log"
	"os"
	"testing"
)

func runWithFixtures(t *testing.T, tests func()) {
	setupPostFixtures()
	tests()
	teardownPostFixtures()
}

func setupPostFixtures() {
	os.MkdirAll("test/data/posts/subdir", os.ModePerm)
	createDummyFile("test/data/posts/a_post.md")
	createDummyFile("test/data/posts/another_post.md")
	createDummyFile("test/data/posts/subdir/a_different_post.md")
}

func createDummyFile(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Cannot create test post:", err)
	}
	defer file.Close()
}

func teardownPostFixtures() {
	os.RemoveAll("test/data/posts")
}
