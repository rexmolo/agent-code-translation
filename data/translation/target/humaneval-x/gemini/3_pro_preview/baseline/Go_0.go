package main

import (
	"fmt"
	"math"
)

// HasCloseElements checks if in the given slice of numbers, any two numbers are closer to each other than the given threshold.
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
	// Test cases from the original Python docstring
	fmt.Println(HasCloseElements([]float64{1.0, 2.0, 3.0}, 0.5))                 // Expected: false
	fmt.Println(HasCloseElements([]float64{1.0, 2.8, 3.0, 4.0, 5.0, 2.0}, 0.3)) // Expected: true
}