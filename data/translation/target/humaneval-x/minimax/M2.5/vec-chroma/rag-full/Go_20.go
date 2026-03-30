package main

import (
	"fmt"
	"math"
)

func FindClosestElements(numbers []float64) [2]float64 {
	var closestPair [2]float64
	var distance float64 = math.MaxFloat64

	n := len(numbers)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			newDistance := math.Abs(numbers[i] - numbers[j])
			if newDistance < distance {
				distance = newDistance
				// Ensure smaller number is first
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
	// Test examples
	result := FindClosestElements([]float64{1.0, 2.0, 3.0, 4.0, 5.0, 2.2})
	fmt.Printf("(%v, %v)\n", result[0], result[1])

	result = FindClosestElements([]float64{1.0, 2.0, 3.0, 4.0, 5.0, 2.0})
	fmt.Printf("(%v, %v)\n", result[0], result[1])
}
