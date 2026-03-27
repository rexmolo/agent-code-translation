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
		// Convert number to string to extract digits
		str := strconv.Itoa(n)
		// Convert each character digit to integer
		digits := make([]int, len(str))
		for i, c := range str {
			digits[i] = int(c - '0')
		}
		// Apply negative sign to first digit if original was negative
		digits[0] = digits[0] * neg
		// Sum all digits
		sum := 0
		for _, d := range digits {
			sum += d
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