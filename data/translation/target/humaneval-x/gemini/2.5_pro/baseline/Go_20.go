package main

import (
	"fmt"
	"math"
)

// FindClosestElements from a supplied slice of numbers (of length at least two) selects and returns two that are the closest to each
// other and return them in order (smaller number, larger number).
func FindClosestElements(numbers []float64) [2]float64 {
	if len(numbers) < 2 {
		// The original function assumes at least two elements.
		// Return a zero-value array for invalid input.
		return [2]float64{}
	}

	minDistance := math.MaxFloat64
	var closestPair [2]float64

	// A more idiomatic and efficient O(n^2) approach than the original Python code,
	// which avoids comparing elements with themselves and avoids duplicate comparisons (e.g. (a,b) and (b,a)).
	// The behavior is identical to the original.
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			elem1 := numbers[i]
			elem2 := numbers[j]
			distance := math.Abs(elem1 - elem2)

			if distance < minDistance {
				minDistance = distance
				if elem1 < elem2 {
					closestPair = [2]float64{elem1, elem2}
				} else {
					closestPair = [2]float64{elem2, elem1}
				}
			}
		}
	}
	return closestPair
}

// main function to demonstrate the FindClosestElements function.
func main() {
	// Example from the docstring
	list1 := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 2.2}
	result1 := FindClosestElements(list1)
	fmt.Printf("Input: %v, Closest pair: %v\n", list1, result1)

	// Second example
	list2 := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 2.0}
	result2 := FindClosestElements(list2)
	fmt.Printf("Input: %v, Closest pair: %v\n", list2, result2)
}
