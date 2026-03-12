package main

import (
	"math"
)

func FindClosestElements(numbers []float64) [2]float64 {
	if len(numbers) < 2 {
		// Handle edge case - though spec says at least 2 elements
		return [2]float64{}
	}

	var closestPair [2]float64
	distance := math.Inf(1) // Start with positive infinity

	// Compare each pair only once (i < j instead of i != j)
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			newDistance := math.Abs(numbers[i] - numbers[j])
			if newDistance < distance {
				distance = newDistance
				closestPair = [2]float64{numbers[i], numbers[j]}
				// Ensure smaller number is first
				if closestPair[0] > closestPair[1] {
					closestPair[0], closestPair[1] = closestPair[1], closestPair[0]
				}
			}
		}
	}

	return closestPair
}