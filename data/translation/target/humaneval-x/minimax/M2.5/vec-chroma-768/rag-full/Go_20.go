package main

import (
	"math"
)

func FindClosestElements(numbers []float64) [2]float64 {
	minDistance := math.Inf(1)
	var closestPair [2]float64

	for i, elem := range numbers {
		for j, elem2 := range numbers {
			if i != j {
				distance := math.Abs(elem - elem2)
				if distance < minDistance {
					minDistance = distance
					// Sort the pair (smaller first)
					if elem < elem2 {
						closestPair = [2]float64{elem, elem2}
					} else {
						closestPair = [2]float64{elem2, elem}
					}
				}
			}
		}
	}

	return closestPair
}

func main() {
	// Test with examples
	result1 := FindClosestElements([]float64{1.0, 2.0, 3.0, 4.0, 5.0, 2.2})
	println(result1[0], result1[1]) // Expected: 2 2.2

	result2 := FindClosestElements([]float64{1.0, 2.0, 3.0, 4.0, 5.0, 2.0})
	println(result2[0], result2[1]) // Expected: 2 2
}
