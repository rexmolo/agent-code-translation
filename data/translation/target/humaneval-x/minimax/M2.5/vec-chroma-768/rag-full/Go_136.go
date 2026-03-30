package main

import (
	"fmt"
	"slices"
)

func LargestSmallestIntegers(lst []int) [2]interface{} {
	var negatives []int
	var positives []int

	// Filter negatives and positives
	for _, v := range lst {
		if v < 0 {
			negatives = append(negatives, v)
		} else if v > 0 {
			positives = append(positives, v)
		}
	}

	// Find largest negative (the negative integer closest to 0)
	var largestNegative interface{}
	if len(negatives) > 0 {
		largestNegative = slices.Max(negatives)
	}

	// Find smallest positive
	var smallestPositive interface{}
	if len(positives) > 0 {
		smallestPositive = slices.Min(positives)
	}

	return [2]interface{}{largestNegative, smallestPositive}
}

func main() {
	// Test cases
	fmt.Println(LargestSmallestIntegers([]int{2, 4, 1, 3, 5, 7})) // [nil 1]
	fmt.Println(LargestSmallestIntegers([]int{}))                 // [nil nil]
	fmt.Println(LargestSmallestIntegers([]int{0}))                 // [nil nil]
	fmt.Println(LargestSmallestIntegers([]int{-5, -2, -8, -1}))    // [-1 nil]
	fmt.Println(LargestSmallestIntegers([]int{1, 2, 3}))           // [nil 1]
	fmt.Println(LargestSmallestIntegers([]int{-3, 0, 5}))          // [-3 5]
}