package main

import (
	"math"
	"slices"
)

func FindClosestElements(numbers []float64) [2]float64 {
	// Make a copy and sort the numbers
	sortedNumbers := make([]float64, len(numbers))
	copy(sortedNumbers, numbers)
	slices.Sort(sortedNumbers)

	// Find the closest pair by checking adjacent elements in sorted list
	var closestPair [2]float64
	minDistance := math.MaxFloat64

	for i := 0; i < len(sortedNumbers)-1; i++ {
		distance := math.Abs(sortedNumbers[i] - sortedNumbers[i+1])
		if distance < minDistance {
			minDistance = distance
			closestPair = [2]float64{sortedNumbers[i], sortedNumbers[i+1]}
		}
	}

	return closestPair
}