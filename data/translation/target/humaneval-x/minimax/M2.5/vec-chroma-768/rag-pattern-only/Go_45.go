package geometry

// TriangleArea calculates the area of a triangle given the base length and height.
// Returns the area as a float64.
func TriangleArea(a float64, h float64) float64 {
	return a * h / 2.0
}