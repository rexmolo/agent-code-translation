package main

import "math"

// FindClosestElements finds the two closest elements in a slice of floats
// and returns them in order (smaller, larger).
func FindClosestElements(numbers []float64) [2]float64 {
	// Initialize with positive infinity to ensure any real distance is smaller
	minDistance := math.Inf(1)
	var closestPair [2]float64

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			distance := math.Abs(numbers[i] - numbers[j])
			if distance < minDistance {
				minDistance = distance
				// Ensure smaller number is first
				if numbers[i] < numbers[j] {
					closestPair[0] = numbers[i]
					closestPair[1] = numbers[j]
				} else {
					closestPair[0] = numbers[j]
					closestPair[1] = numbers[i]
				}
			}
		}
	}

	return closestPair
}