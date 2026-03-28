package main

import (
	"strconv"
)

func CountNums(arr []int) int {
	digitsSum := func(n int) int {
		neg := 1
		if n < 0 {
			n = -n
			neg = -1
		}

		// Handle the case when n is 0
		if n == 0 {
			return 0
		}

		// Convert to string to get digits
		str := strconv.Itoa(n)
		sum := 0
		for i, c := range str {
			digit := int(c - '0')
			if i == 0 {
				digit = digit * neg
			}
			sum += digit
		}
		return sum
	}

	count := 0
	for _, n := range arr {
		if digitsSum(n) > 0 {
			count++
		}
	}
	return count
}

func main() {
	// Test cases to verify
	// CountNums([]int{}) == 0
	// CountNums([]int{-1, 11, -11}) == 1
	// CountNums([]int{1, 1, 2}) == 3
}
