package main

import (
	"fmt"
	"math"
)

func HasCloseElements(numbers []float64, threshold float64) bool {
	for idx, elem := range numbers {
		for idx2, elem2 := range numbers {
			if idx != idx2 {
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
	// Test case 1: should be false
	result1 := HasCloseElements([]float64{1.0, 2.0, 3.0}, 0.5)
	fmt.Println(result1)

	// Test case 2: should be true (2.8 and 3.0 differ by 0.2 < 0.3)
	result2 := HasCloseElements([]float64{1.0, 2.8, 3.0, 4.0, 5.0, 2.0}, 0.3)
	fmt.Println(result2)
}
