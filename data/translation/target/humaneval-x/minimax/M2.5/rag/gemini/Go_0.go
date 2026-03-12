package main

import "math"

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
	// Example usage:
	numbers1 := []float64{1.0, 2.0, 3.0}
	print(HasCloseElements(numbers1, 0.5)) // false
	
	numbers2 := []float64{1.0, 2.8, 3.0, 4.0, 5.0, 2.0}
	print(HasCloseElements(numbers2, 0.3)) // true
}
