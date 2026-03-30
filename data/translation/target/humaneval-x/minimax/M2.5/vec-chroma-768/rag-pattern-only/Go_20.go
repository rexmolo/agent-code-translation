package main

import (
	"fmt"
	"math"
)

func FindClosestElements(numbers []float64) [2]float64 {
	// Initialize with a sentinel value for "not set"
	var closestPair [2]float64
	var minDistance float64 = math.Inf(1)
	initialized := false

	for idx, elem := range numbers {
		for idx2, elem2 := range numbers {
			if idx != idx2 {
				newDistance := math.Abs(elem - elem2)
				if !initialized {
					minDistance = newDistance
					if elem < elem2 {
						closestPair[0] = elem
						closestPair[1] = elem2
					} else {
						closestPair[0] = elem2
						closestPair[1] = elem
					}
					initialized = true
				} else if newDistance < minDistance {
					minDistance = newDistance
					if elem < elem2 {
						closestPair[0] = elem
						closestPair[1] = elem2
					} else {
						closestPair[0] = elem2
						closestPair[1] = elem
					}
				}
			}
		}
	}

	return closestPair
}

func main() {
	// Test the function
	result := FindClosestElements([]float64{1.0, 2.0, 3.0, 4.0, 5.0, 2.2})
	fmt.Printf("Result: [%.1f, %.1f]\n", result[0], result[1])

	result2 := FindClosestElements([]float64{1.0, 2.0, 3.0, 4.0, 5.0, 2.0})
	fmt.Printf("Result: [%.1f, %.1f]\n", result2[0], result2[1])
}
