package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"os"
	"runtime"
)

func main() {
	pictureWidth := flag.Int("width", 1024, "width of the target picture")
	pictureHeight := flag.Int("height", 1024, "height of the target picture")
	viewRectangleMinX := flag.Float64("minX", -2, "view rectangle min x")
	viewRectangleMinY := flag.Float64("minY", -1, "view rectangle min y")
	viewRectangleMaxX := flag.Float64("maxX", 1, "view rectangle max x")
	viewRectangleMaxY := flag.Float64("maxY", 1, "view rectangle max y")
	iterations := flag.Int("i", 100, "number of iterations")
	imagePath := flag.String("p", "./mandelbrot.png", "image path")

	flag.Parse()

	fmt.Printf("width: %d\nheight: %d\nminX: %f\nminY: %f\nmaxX: %f\nmaxY: %f\niterations: %d\nimagePath: %s",
		*pictureWidth,
		*pictureHeight,
		*viewRectangleMaxX,
		*viewRectangleMaxY,
		*viewRectangleMinX,
		*viewRectangleMinY,
		*iterations,
		*imagePath)
	numOfProcs := runtime.GOMAXPROCS(0)
	fmt.Printf("GOMAXPROCS %d", numOfProcs)

	pc := NewPixelCalculator(*iterations, * pictureWidth, *pictureHeight, *viewRectangleMaxX, *viewRectangleMaxY, *viewRectangleMinX, *viewRectangleMinY)

	img := image.NewRGBA(image.Rect(0, 0, *pictureWidth, *pictureHeight))

	numberOfPoints := *pictureWidth * *pictureHeight

	points := make(chan point, 10)
	calculatedPoints := make(chan *ColorPoint, 10)

	for w := 0; w < numOfProcs; w += 1 {
		go worker(pc, points, calculatedPoints)
	}

	go addPointsToCalculate(pictureWidth, pictureHeight, points)

	for i := 0; i < numberOfPoints; i += 1 {
		colorPoint := <-calculatedPoints
		img.Set(colorPoint.x, colorPoint.y, colorPoint.color)
	}

	// outputFile is a File type which satisfies Writer interface
	outputFile, err := os.Create(*imagePath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// Encode takes a writer interface and an image interface
	// We pass it the File and the RGBA
	err = png.Encode(outputFile, img)
	if err != nil {
		panic(err)
	}

}

func addPointsToCalculate(pictureWidth *int, pictureHeight *int, points chan<- point) {
	for pX := 0; pX < *pictureWidth; pX += 1 {
		for pY := 0; pY < *pictureHeight; pY += 1 {
			points <- point{x: pX, y: pY}
		}
	}
	// No new points will be added to the channel
	close(points)
}

type point struct {
	x int
	y int
}

func worker(pixelCalculator PixelCalculator, points <-chan point, calculatedPoints chan<- *ColorPoint) {
	for point := range points {
		calculatedPoints <- pixelCalculator.Calculate(point.x, point.y)
	}
}
