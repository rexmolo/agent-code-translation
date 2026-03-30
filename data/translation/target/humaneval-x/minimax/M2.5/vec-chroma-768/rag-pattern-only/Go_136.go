package main

import (
	"fmt"
)

func LargestSmallestIntegers(lst []int) [2]interface{} {
	var negatives []int
	var positives []int

	// Filter negative and positive integers
	for _, x := range lst {
		if x < 0 {
			negatives = append(negatives, x)
		} else if x > 0 {
			positives = append(positives, x)
		}
	}

	var largestNeg, smallestPos interface{}

	// Find largest negative (closest to zero)
	if len(negatives) > 0 {
		largestNeg = negatives[0]
		for _, x := range negatives[1:] {
			if x > largestNeg.(int) {
				largestNeg = x
			}
		}
	}

	// Find smallest positive
	if len(positives) > 0 {
		smallestPos = positives[0]
		for _, x := range positives[1:] {
			if x < smallestPos.(int) {
				smallestPos = x
			}
		}
	}

	return [2]interface{}{largestNeg, smallestPos}
}

func main() {
	// Test examples
	testCases := [][]int{
		{2, 4, 1, 3, 5, 7},
		{},
		{0},
		{-5, -3, -1, 2, 4},
		{-10, -5, 1, 3},
	}

	for _, tc := range testCases {
		result := LargestSmallestIntegers(tc)
		fmt.Printf("LargestSmallestIntegers(%v) = %v\n", tc, result)
	}
}
