package main

import (
	"image"
	"testing"
)

type testImageWriter struct{}

func (testImageWriter) Write(path string, image *image.RGBA) {
	// do nothing
}

var iterations = 1000
var width = 1024
var height = 1024
var minX = -2.0
var minY = -1.0
var maxX = 1.0
var maxY = 1.0

func Benchmark_sequentialImageGeneration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pc := NewPixelCalculator(iterations, width, height, maxX, maxY, minX, minY)
		img := image.NewRGBA(image.Rect(0, 0, width, height))

		for pX := 0; pX < width; pX++ {
			for pY := 0; pY < height; pY++ {
				cp := pc.Calculate(pX, pY)
				img.Set(cp.x, cp.y, cp.color)
			}
		}

		testImageWriter{}.Write("test.png", img)
	}
}

func Benchmark_concurrentImageGeneration_1worker(b *testing.B) {
	for i := 0; i < b.N; i++ {
		run(iterations, width, height, maxX, maxY, minX, minY, 1, "test.png", testImageWriter{})
	}
}

func Benchmark_concurrentImageGeneration_2workers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		run(iterations, width, height, maxX, maxY, minX, minY, 2, "test.png", testImageWriter{})
	}
}

func Benchmark_concurrentImageGeneration_3workers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		run(iterations, width, height, maxX, maxY, minX, minY, 3, "test.png", testImageWriter{})
	}
}

func Benchmark_concurrentImageGeneration_4workers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		run(iterations, width, height, maxX, maxY, minX, minY, 4, "test.png", testImageWriter{})
	}
}

func Benchmark_concurrentImageGeneration_10workers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		run(iterations, width, height, maxX, maxY, minX, minY, 10, "test.png", testImageWriter{})
	}
}

func Benchmark_concurrentImageGeneration_100workers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		run(iterations, width, height, maxX, maxY, minX, minY, 100, "test.png", testImageWriter{})
	}
}