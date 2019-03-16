package main

func normalizeToViewRectangle(pX, pY, w, h int, minX, minY, maxX, maxY float64) (float64, float64) {
	cX := (float64(pX)-0)*((maxX-minX)/(float64(w)-0)) + minX
	cY := (float64(pY)-0)*((maxY-minY)/(float64(h)-0)) + minY

	return cX, cY
}
