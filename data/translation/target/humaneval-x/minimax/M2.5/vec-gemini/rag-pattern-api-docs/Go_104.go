package main

import (
	"fmt"
	"sort"
	"strconv"
)

func UniqueDigits(x []int) []int {
	var oddDigitElements []int

	for _, i := range x {
		// Convert int to string to check each digit
		str := strconv.Itoa(i)
		allOdd := true
		for _, c := range str {
			// Convert rune to digit value
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

	// Sort the result in increasing order
	sort.Slice(oddDigitElements, func(i, j int) bool {
		return oddDigitElements[i] < oddDigitElements[j]
	})

	return oddDigitElements
}

func main() {
	// Test examples
	fmt.Println(UniqueDigits([]int{15, 33, 1422, 1}))   // Output: [1 15 33]
	fmt.Println(UniqueDigits([]int{152, 323, 1422, 10})) // Output: []
}
