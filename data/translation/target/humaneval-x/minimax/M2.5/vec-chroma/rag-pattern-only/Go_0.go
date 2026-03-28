package main

import (
	"math"
)

func HasCloseElements(numbers []float64, threshold float64) bool {
	// If there are fewer than 2 elements, no two elements can be close
	if len(numbers) < 2 {
		return false
	}

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

// For testing purposes
func main() {
	// Test case 1: [1.0, 2.0, 3.0] with threshold 0.5 -> false
	// Test case 2: [1.0, 2.8, 3.0, 4.0, 5.0, 2.0] with threshold 0.3 -> true
}