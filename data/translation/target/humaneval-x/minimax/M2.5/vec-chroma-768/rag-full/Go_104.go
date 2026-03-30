package main

import (
	"sort"
	"strconv"
)

func UniqueDigits(x []int) []int {
	var result []int

	for _, num := range x {
		// Convert number to string to check each digit
		strNum := strconv.Itoa(num)

		// Check if all digits are odd
		allOdd := true
		for _, ch := range strNum {
			digit := int(ch - '0') // Convert rune to digit value
			if digit%2 == 0 { // Even digit found
				allOdd = false
				break
			}
		}

		if allOdd {
			result = append(result, num)
		}
	}

	sort.Ints(result)
	return result
}
