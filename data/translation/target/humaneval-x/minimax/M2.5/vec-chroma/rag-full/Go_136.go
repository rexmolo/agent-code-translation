package main

import (
	"fmt"
	"slices"
)

func LargestSmallestIntegers(lst []int) [2]interface{} {
	var negatives []int
	var positives []int

	// Filter integers: negatives (< 0) and positives (> 0)
	for _, v := range lst {
		if v < 0 {
			negatives = append(negatives, v)
		} else if v > 0 {
			positives = append(positives, v)
		}
	}

	// Find largest negative (max of negatives)
	var largestNegative interface{}
	if len(negatives) > 0 {
		largestNegative = slices.Max(negatives)
	}

	// Find smallest positive (min of positives)
	var smallestPositive interface{}
	if len(positives) > 0 {
		smallestPositive = slices.Min(positives)
	}

	return [2]interface{}{largestNegative, smallestPositive}
}

func main() {
	// Test examples
	examples := [][]int{
		{2, 4, 1, 3, 5, 7},
		{},
		{0},
	}

	for _, lst := range examples {
		result := LargestSmallestIntegers(lst)
		fmt.Printf("%v -> %v\n", lst, result)
	}
}
