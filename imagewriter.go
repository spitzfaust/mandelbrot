package main

import (
	"image"
	"image/png"
	"os"
)

type ImageWriter interface {
	Write(path string, image *image.RGBA)
}

type imageWriter struct{}

func NewImageWriter() ImageWriter {
	return imageWriter{}
}

func (imageWriter) Write(path string, image *image.RGBA) {
	// outputFile is where the image will be written to
	outputFile, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, image)
	if err != nil {
		panic(err)
	}
}
