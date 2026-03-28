package main

import (
	"fmt"
	"math"
)

func FindClosestElements(numbers []float64) [2]float64 {
	var closestPair [2]float64
	var distance float64
	initialized := false

	for i, elem := range numbers {
		for j, elem2 := range numbers {
			if i != j {
				newDistance := math.Abs(elem - elem2)
				if !initialized || newDistance < distance {
					distance = newDistance
					initialized = true
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
	// Test cases
	result := FindClosestElements([]float64{1.0, 2.0, 3.0, 4.0, 5.0, 2.2})
	fmt.Println(result)
	result = FindClosestElements([]float64{1.0, 2.0, 3.0, 4.0, 5.0, 2.0})
	fmt.Println(result)
}
