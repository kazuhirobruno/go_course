package area

import "math"

// Circle returns the area of a circle with the given radius
func Circle(radius float64) float64 {
	return math.Pi * radius * radius
}

// Rect returns the area of a rectangle with the given width and height
func Rect(width, height float64) float64 {
	return width * height
}