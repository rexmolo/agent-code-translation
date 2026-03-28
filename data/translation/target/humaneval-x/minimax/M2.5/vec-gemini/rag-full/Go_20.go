package main

import (
	"math"
)

func FindClosestElements(numbers []float64) [2]float64 {
	var closestPair [2]float64
	distance := math.MaxFloat64
	initialized := false

	for i, elem := range numbers {
		for j, elem2 := range numbers {
			if i != j {
				newDistance := math.Abs(elem - elem2)
				if !initialized || newDistance < distance {
					distance = newDistance
					initialized = true
					// Sort the pair (smaller first, larger second)
					if elem < elem2 {
						closestPair[0] = elem
						closestPair[1] = elem2
					} else {
						closestPair[0] = elem2
						closestPair[1] = elem
					}
				}
			}
		}
	}

	return closestPair
}
