package main

import (
	"github.com/lucasb-eyer/go-colorful"
	"math"
)

type PixelCalculator interface {
	Calculate(pX, pY int) *ColorPoint
}

type pixelCalculator struct {
	maxIterations     int
	pictureWidth      int
	pictureHeight     int
	viewRectangleMaxX float64
	viewRectangleMaxY float64
	viewRectangleMinX float64
	viewRectangleMinY float64
}

func NewPixelCalculator(maxIterations, pictureWidth, pictureHeight int, viewRectangleMaxX, viewRectangleMaxY, viewRectangleMinX, viewRectangleMinY float64) PixelCalculator {
	return pixelCalculator{
		maxIterations:     maxIterations,
		pictureWidth:      pictureWidth,
		pictureHeight:     pictureHeight,
		viewRectangleMaxX: viewRectangleMaxX,
		viewRectangleMaxY: viewRectangleMaxY,
		viewRectangleMinX: viewRectangleMinX,
		viewRectangleMinY: viewRectangleMinY,
	}
}

func (pc pixelCalculator) Calculate(pX, pY int) *ColorPoint {
	cX, cY := normalizeToViewRectangle(pX, pY, pc.pictureWidth, pc.pictureHeight, pc.viewRectangleMinX, pc.viewRectangleMinY, pc.viewRectangleMaxX, pc.viewRectangleMaxY)
	zX := cX
	zY := cY
	n := 0
	for ; n < pc.maxIterations; n += 1 {
		x := (zX*zX - zY*zY) + cX
		y := (zY*zX + zX*zY) + cY
		if (x*x + y*y) > 4 {
			break
		}
		zX = x
		zY = y
	}
	smooth := float64(n) + 1 - math.Log(math.Log2(math.Abs(zX*zY)))
	h := float64(360) * smooth / float64(pc.maxIterations)
	s := float64(1)
	var v float64
	if n < pc.maxIterations {
		v = 1
	} else {
		v = 0
	}
	return &ColorPoint{
		x:     pX,
		y:     pY,
		color: colorful.Hsv(h, s, v),
	}
}
