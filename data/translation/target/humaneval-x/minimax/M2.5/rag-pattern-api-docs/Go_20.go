package main

import (
	"math"
	"fmt"
)

func FindClosestElements(numbers []float64) [2]float64 {
	var closestPair [2]float64
	distance := math.Inf(1) // Start with positive infinity

	for i, elem := range numbers {
		for j, elem2 := range numbers {
			if i != j {
				newDistance := math.Abs(elem - elem2)
				if newDistance < distance {
					distance = newDistance
					// Sort the pair (smaller, larger)
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
	// Test examples
	result1 := FindClosestElements([]float64{1.0, 2.0, 3.0, 4.0, 5.0, 2.2})
	fmt.Println(result1) // [2 2.2]

	result2 := FindClosestElements([]float64{1.0, 2.0, 3.0, 4.0, 5.0, 2.0})
	fmt.Println(result2) // [2 2]
}
