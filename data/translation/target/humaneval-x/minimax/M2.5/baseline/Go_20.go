package main

import (
	"math"
)

func FindClosestElements(numbers []float64) [2]float64 {
	if len(numbers) < 2 {
		return [2]float64{0, 0}
	}

	// Initialize with first pair
	closestPair := [2]float64{numbers[0], numbers[1]}
	minDistance := math.Abs(numbers[0] - numbers[1])

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			distance := math.Abs(numbers[i] - numbers[j])
			if distance < minDistance {
				minDistance = distance
				// Ensure smaller element is first
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

func main() {
	// Example usage
	result := FindClosestElements([]float64{1.0, 2.0, 3.0, 4.0, 5.0, 2.2})
	println(result[0], result[1]) // Output: 2 2.2

	result2 := FindClosestElements([]float64{1.0, 2.0, 3.0, 4.0, 5.0, 2.0})
	println(result2[0], result2[1]) // Output: 2 2
}