package main

import (
	"fmt"
	"slices"
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

	// Find largest negative (max of negatives) and smallest positive (min of positives)
	var largestNegative, smallestPositive interface{}

	if len(negatives) > 0 {
		largestNegative = slices.Max(negatives)
	}

	if len(positives) > 0 {
		smallestPositive = slices.Min(positives)
	}

	return [2]interface{}{largestNegative, smallestPositive}
}

func main() {
	// Test cases
	fmt.Println(LargestSmallestIntegers([]int{2, 4, 1, 3, 5, 7})) // [<nil> 1]
	fmt.Println(LargestSmallestIntegers([]int{}))               // [<nil> <nil>]
	fmt.Println(LargestSmallestIntegers([]int{0}))              // [<nil> <nil>]
	fmt.Println(LargestSmallestIntegers([]int{-5, -2, -8}))     // [-2 <nil>]
	fmt.Println(LargestSmallestIntegers([]int{1, 2, 3}))        // [<nil> 1]
	fmt.Println(LargestSmallestIntegers([]int{-1, 2, -3, 4}))   // [-1 2]
}
