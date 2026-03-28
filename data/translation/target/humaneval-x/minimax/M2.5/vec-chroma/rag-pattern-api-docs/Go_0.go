package main

import (
	"math"
)

// HasCloseElements checks if in given slice of numbers, there are any two numbers
// closer to each other than the given threshold.
func HasCloseElements(numbers []float64, threshold float64) bool {
	for i, elem := range numbers {
		for j, elem2 := range numbers {
			if i != j {
				distance := math.Abs(elem - elem2)
				if distance < threshold {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	// Example usage (for demonstration)
	_ = HasCloseElements([]float64{1.0, 2.0, 3.0}, 0.5)   // false
	_ = HasCloseElements([]float64{1.0, 2.8, 3.0, 4.0, 5.0, 2.0}, 0.3) // true
}
