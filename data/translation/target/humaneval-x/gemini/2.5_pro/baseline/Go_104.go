package main

import (
	"fmt"
	"sort"
	"strconv"
)

// UniqueDigits takes a slice of positive integers and returns a sorted slice
// containing only the elements that have no even digits.
func UniqueDigits(x []int) []int {
	var oddDigitElements []int
	for _, i := range x {
		// Convert the integer to a string to iterate over its digits.
		s := strconv.Itoa(i)
		hasEvenDigit := false
		for _, c := range s {
			// Convert the character digit back to an integer.
			// c is a rune (int32). Subtracting '0' gives the integer value.
			digit := int(c - '0')
			if digit%2 == 0 {
				hasEvenDigit = true
				break
			}
		}
		// If no even digit was found, add the number to our result slice.
		if !hasEvenDigit {
			oddDigitElements = append(oddDigitElements, i)
		}
	}
	// Sort the resulting slice in increasing order.
	sort.Ints(oddDigitElements)
	return oddDigitElements
}

func main() {
	// Example 1 from the docstring
	input1 := []int{15, 33, 1422, 1}
	result1 := UniqueDigits(input1)
	fmt.Println(result1)

	// Example 2 from the docstring
	input2 := []int{152, 323, 1422, 10}
	result2 := UniqueDigits(input2)
	fmt.Println(result2)
}
