package main

import "fmt"

func LargestSmallestIntegers(lst []int) [2]interface{} {
	var largestNegative *int
	var smallestPositive *int

	// Find largest negative and smallest positive integers
	for _, v := range lst {
		if v < 0 {
			if largestNegative == nil || v > *largestNegative {
				largestNegative = &v
			}
		}
		if v > 0 {
			if smallestPositive == nil || v < *smallestPositive {
				smallestPositive = &v
			}
		}
	}

	// Convert to interface{} (nil if not found)
	var largest, smallest interface{}
	if largestNegative != nil {
		largest = *largestNegative
	}
	if smallestPositive != nil {
		smallest = *smallestPositive
	}

	return [2]interface{}{largest, smallest}
}

func main() {
	// Test cases
	tests := [][]int{
		{2, 4, 1, 3, 5, 7},
		{},
		{0},
		{-5, -3, -1},
		{1, 2, 3},
		{-2, 0, 5, -8, 3},
	}

	for _, lst := range tests {
		result := LargestSmallestIntegers(lst)
		fmt.Printf("Input: %v -> Output: %v\n", lst, result)
	}
}
