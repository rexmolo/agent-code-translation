package main

import (
	"sort"
	"strconv"
)

func UniqueDigits(x []int) []int {
	var result []int

	for _, num := range x {
		// Convert number to string to iterate over digits
		s := strconv.Itoa(num)

		// Check if all digits are odd (no even digits)
		allOdd := true
		for _, r := range s {
			digit := int(r - '0')
			if digit%2 == 0 {
				allOdd = false
				break
			}
		}

		if allOdd {
			result = append(result, num)
		}
	}

	// Sort the result in increasing order
	sort.Ints(result)
	return result
}
