package main

import (
	"math"
)

func FindClosestElements(numbers []float64) [2]float64 {
	closestPair := [2]float64{}
	distance := math.Inf(1)
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