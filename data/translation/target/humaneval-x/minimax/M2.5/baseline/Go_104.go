package main

import (
	"sort"
	"strconv"
)

func UniqueDigits(x []int) []int {
	var oddDigitElements []int

	for _, i := range x {
		// Convert number to string to check each digit
		str := strconv.Itoa(i)

		// Check if all digits are odd
		allOdd := true
		for _, c := range str {
			digit := int(c - '0')
			if digit%2 == 0 {
				allOdd = false
				break
			}
		}

		if allOdd {
			oddDigitElements = append(oddDigitElements, i)
		}
	}

	// Sort in increasing order
	sort.Ints(oddDigitElements)

	return oddDigitElements
}
