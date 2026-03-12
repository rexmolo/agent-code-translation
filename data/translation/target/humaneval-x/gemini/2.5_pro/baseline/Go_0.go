package main

import (
	"fmt"
	"math"
)

// HasCloseElements checks if in a given slice of numbers, are any two numbers
// closer to each other than a given threshold.
func HasCloseElements(numbers []float64, threshold float64) bool {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			distance := math.Abs(numbers[i] - numbers[j])
			if distance < threshold {
				return true
			}
		}
	}
	return false
}

// main function to demonstrate the usage of HasCloseElements
func main() {
	// Python: has_close_elements([1.0, 2.0, 3.0], 0.5) -> False
	fmt.Println(HasCloseElements([]float64{1.0, 2.0, 3.0}, 0.5))

	// Python: has_close_elements([1.0, 2.8, 3.0, 4.0, 5.0, 2.0], 0.3) -> True
	fmt.Println(HasCloseElements([]float64{1.0, 2.8, 3.0, 4.0, 5.0, 2.0}, 0.3))
}
