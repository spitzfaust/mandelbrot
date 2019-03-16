package main

import (
	"github.com/lucasb-eyer/go-colorful"
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
	for n := 0; n < pc.maxIterations; n += 1 {
		x := (zX*zX - zY*zY) + cX
		y := (zY*zX + zX*zY) + cY
		if (x*x + y*y) > 4 {
			h := float64(360) * float64(n) / float64(pc.maxIterations)
			s := float64(1)
			v := float64(1)
			return &ColorPoint{
				x:     pX,
				y:     pY,
				color: colorful.Hsv(h, s, v),
			}
		}
		zX = x
		zY = y
	}
	return &ColorPoint{
		x:     pX,
		y:     pY,
		color: colorful.Hsv(1, 1, 0),
	}
}
