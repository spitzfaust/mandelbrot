package main

import (
	"flag"
	"fmt"
	"image"
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

	fmt.Printf("width: %d\nheight: %d\nminX: %f\nminY: %f\nmaxX: %f\nmaxY: %f\niterations: %d\nimagePath: %s\n",
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

	run(*iterations, *pictureWidth, *pictureHeight, *viewRectangleMaxX, *viewRectangleMaxY, *viewRectangleMinX, *viewRectangleMinY, numOfProcs, *imagePath, NewImageWriter())
}

func run(iterations int, pictureWidth int, pictureHeight int, viewRectangleMaxX float64, viewRectangleMaxY float64, viewRectangleMinX float64, viewRectangleMinY float64, numOfProcs int, imagePath string, imageWriter ImageWriter) {
	pc := NewPixelCalculator(iterations, pictureWidth, pictureHeight, viewRectangleMaxX, viewRectangleMaxY, viewRectangleMinX, viewRectangleMinY)
	numberOfPoints := pictureWidth * pictureHeight
	// create two buffered channels
	// One for the points to calculate and one for the results
	points := make(chan point, numberOfPoints)
	calculatedPoints := make(chan *ColorPoint, numberOfPoints)
	// Start workers
	for w := 0; w < numOfProcs; w++ {
		go worker(pc, points, calculatedPoints)
	}
	// Fill the channel with points that should be calculated
	go addPointsToCalculate(pictureWidth, pictureHeight, points)
	done := make(chan bool)
	img := image.NewRGBA(image.Rect(0, 0, pictureWidth, pictureHeight))
	go createImage(numberOfPoints, calculatedPoints, img, imagePath, imageWriter, done)
	<-done
}

func createImage(numberOfPoints int, calculatedPoints <-chan *ColorPoint, img *image.RGBA, imagePath string, imageWriter ImageWriter, done chan<- bool) {
	for i := 0; i < numberOfPoints; i++ {
		colorPoint := <-calculatedPoints
		img.Set(colorPoint.x, colorPoint.y, colorPoint.color)
	}
	imageWriter.Write(imagePath, img)
	done <- true
}

func addPointsToCalculate(pictureWidth int, pictureHeight int, points chan<- point) {
	for pX := 0; pX < pictureWidth; pX++ {
		for pY := 0; pY < pictureHeight; pY++ {
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
