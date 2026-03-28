package main

import (
	"fmt"
	"math"
	"sort"
)

// FindClosestElements finds the two closest numbers in a slice and returns them
// in ascending order. The input slice must have at least 2 elements.
func FindClosestElements(numbers []float64) [2]float64 {
	if len(numbers) < 2 {
		panic("FindClosestElements requires at least 2 numbers")
	}

	// Sort a copy to preserve the original slice and find adjacent closest elements
	sorted := make([]float64, len(numbers))
	copy(sorted, numbers)
	sort.Float64s(sorted)

	// The closest pair must be adjacent in sorted order
	minDist := math.Abs(sorted[0] - sorted[1])
	closestPair := [2]float64{sorted[0], sorted[1]}

	for i := 1; i < len(sorted)-1; i++ {
		dist := math.Abs(sorted[i] - sorted[i+1])
		if dist < minDist {
			minDist = dist
			closestPair = [2]float64{sorted[i], sorted[i+1]}
		}
	}

	return closestPair
}

func main() {
	fmt.Println(FindClosestElements([]float64{1.0, 2.0, 3.0, 4.0, 5.0, 2.2}))
	fmt.Println(FindClosestElements([]float64{1.0, 2.0, 3.0, 4.0, 5.0, 2.0}))
}
