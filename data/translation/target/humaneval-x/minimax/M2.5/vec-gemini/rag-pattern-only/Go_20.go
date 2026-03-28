package main

import (
	"fmt"
)

// FindClosestElements finds the two closest elements in a slice of floats
// and returns them sorted (smaller first, larger second).
func FindClosestElements(numbers []float64) [2]float64 {
	if len(numbers) < 2 {
		panic("List must contain at least two elements")
	}

	// Initialize with the first two elements
	closestPair := [2]float64{numbers[0], numbers[1]}
	if numbers[0] > numbers[1] {
		closestPair[0], closestPair[1] = numbers[1], numbers[0]
	}
	minDistance := abs(numbers[0] - numbers[1])

	// Compare all pairs
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			distance := abs(numbers[i] - numbers[j])
			if distance < minDistance {
				minDistance = distance
				// Sort the pair (smaller first)
				if numbers[i] < numbers[j] {
					closestPair = [2]float64{numbers[i], numbers[j]}
				} else {
					closestPair = [2]float64{numbers[j], numbers[i]}
				}
			}
		}
	}

	return closestPair
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(FindClosestElements([]float64{1.0, 2.0, 3.0, 4.0, 5.0, 2.2}))
	fmt.Println(FindClosestElements([]float64{1.0, 2.0, 3.0, 4.0, 5.0, 2.0}))
}
