package main

import (
	"fmt"
	"strconv"
)

func CountNums(arr []int) int {
	// Nested function to compute digit sum with signed first digit
	digitsSum := func(n int) int {
		neg := 1
		if n < 0 {
			n = -n
			neg = -1
		}
		// Convert to string and compute digit sum
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

	// Count elements with digit sum > 0
	count := 0
	for _, n := range arr {
		if digitsSum(n) > 0 {
			count++
		}
	}
	return count
}

func main() {
	// Test cases
	fmt.Println(CountNums([]int{}) == 0)          // true
	fmt.Println(CountNums([]int{-1, 11, -11}) == 1) // true
	fmt.Println(CountNums([]int{1, 1, 2}) == 3)    // true
}