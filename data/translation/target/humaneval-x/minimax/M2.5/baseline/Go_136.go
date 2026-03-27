package main

import (
	"fmt"
)

func LargestSmallestIntegers(lst []int) [2]interface{} {
	var negatives []int
	var positives []int

	// Separate negatives and positives
	for _, num := range lst {
		if num < 0 {
			negatives = append(negatives, num)
		} else if num > 0 {
			positives = append(positives, num)
		}
	}

	// Find largest negative (max of negatives)
	var largestNegative interface{}
	if len(negatives) > 0 {
		maxNeg := negatives[0]
		for _, n := range negatives[1:] {
			if n > maxNeg {
				maxNeg = n
			}
		}
		largestNegative = maxNeg
	}

	// Find smallest positive (min of positives)
	var smallestPositive interface{}
	if len(positives) > 0 {
		minPos := positives[0]
		for _, p := range positives[1:] {
			if p < minPos {
				minPos = p
			}
		}
		smallestPositive = minPos
	}

	return [2]interface{}{largestNegative, smallestPositive}
}

func main() {
	// Test cases
	fmt.Println(LargestSmallestIntegers([]int{2, 4, 1, 3, 5, 7})) // [None 1]
	fmt.Println(LargestSmallestIntegers([]int{}))               // [None None]
	fmt.Println(LargestSmallestIntegers([]int{0}))                // [None None]
}