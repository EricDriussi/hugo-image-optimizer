package filesystemrepo_test

import (
	"crypto/rand"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"testing"
)

func runWithFixtures(t *testing.T, tests func()) {
	setupImageFixtures()
	tests()
	teardownImageFixtures()
}

func setupImageFixtures() {
	os.MkdirAll("test/data/images/donation/subdir", os.ModePerm)
	os.MkdirAll("test/data/images/whoami", os.ModePerm)
	createDummyPng("test/data/images/an_image.png")
	createDummyGif("test/data/images/a_gif.gif")
	createDummyJpg("test/data/images/another_image.jpeg")
	createDummyJpg("test/data/images/donation/subdir/ignore_me.jpg")
	createDummyJpg("test/data/images/whoami/avatar.jpg")
}

func createDummyPng(filePath string) {
	img := createRandomImage()
	imgFile, err := os.Create(filePath)
	defer imgFile.Close()
	if err != nil {
		log.Fatal("Cannot create test image:", err)
	}
	png.Encode(imgFile, img.SubImage(img.Rect))
}

func createDummyJpg(filePath string) {
	img := createRandomImage()
	imgFile, err := os.Create(filePath)
	defer imgFile.Close()
	if err != nil {
		log.Fatal("Cannot create test image:", err)
	}
	jpeg.Encode(imgFile, img.SubImage(img.Rect), &jpeg.Options{})
}

func createDummyGif(filePath string) {
	img := createRandomImage()
	imgFile, err := os.Create(filePath)
	defer imgFile.Close()
	if err != nil {
		log.Fatal("Cannot create test image:", err)
	}
	gif.Encode(imgFile, img.SubImage(img.Rect), &gif.Options{})
}

func createRandomImage() (created *image.NRGBA) {
	rect := image.Rect(0, 0, 100, 100)
	pix := make([]uint8, rect.Dx()*rect.Dy()*4)
	rand.Read(pix)
	return &image.NRGBA{
		Pix:    pix,
		Stride: rect.Dx() * 4,
		Rect:   rect,
	}
}

func teardownImageFixtures() {
	os.RemoveAll("test/data/images")
}
