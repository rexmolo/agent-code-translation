package main

import (
	"fmt"
)

// LargestSmallestIntegers creates an array [a, b], where 'a' is
// the largest of negative integers, and 'b' is the smallest
// of positive integers in a slice.
// If there is no negative or positive integer, the corresponding value is nil.
func LargestSmallestIntegers(lst []int) [2]interface{} {
	var largestNegative *int
	var smallestPositive *int

	for _, num := range lst {
		if num < 0 {
			// If this is the first negative number found, or it's larger
			// than the current largest negative (i.e., closer to zero).
			if largestNegative == nil || num > *largestNegative {
				val := num // Create a new variable to hold the value.
				largestNegative = &val // Point to the new variable.
			}
		} else if num > 0 {
			// If this is the first positive number found, or it's smaller
			// than the current smallest positive.
			if smallestPositive == nil || num < *smallestPositive {
				val := num
				smallestPositive = &val
			}
		}
	}

	var result [2]interface{}

	if largestNegative != nil {
		result[0] = *largestNegative
	} else {
		result[0] = nil
	}

	if smallestPositive != nil {
		result[1] = *smallestPositive
	} else {
		result[1] = nil
	}

	return result
}

// main function to demonstrate the usage of LargestSmallestIntegers
// and make the code runnable.
func main() {
	fmt.Println(LargestSmallestIntegers([]int{2, 4, 1, 3, 5, 7}))
	fmt.Println(LargestSmallestIntegers([]int{}))
	fmt.Println(LargestSmallestIntegers([]int{0}))
	fmt.Println(LargestSmallestIntegers([]int{-1, -2, -3, 1, 2, 3}))
	fmt.Println(LargestSmallestIntegers([]int{-10, -5, -2}))
}
