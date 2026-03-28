package main

import (
	"sort"
)

func UniqueDigits(x []int) []int {
	var result []int

	for _, n := range x {
		if allDigitsOdd(n) {
			result = append(result, n)
		}
	}

	sort.Ints(result)
	return result
}

// allDigitsOdd checks if all digits of n are odd (no even digits)
func allDigitsOdd(n int) bool {
	if n == 0 {
		return false // 0 is an even digit
	}

	for n > 0 {
		digit := n % 10
		if digit%2 == 0 {
			return false
		}
		n /= 10
	}
	return true
}

func main() {
	// Test examples
	 println(UniqueDigits([]int{15, 33, 1422, 1}))   // Output: [1 15 33]
	 println(UniqueDigits([]int{152, 323, 1422, 10})) // Output: []
}