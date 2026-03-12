package main

import (
	"fmt"
	"strconv"
)

// digitsSum calculates the sum of digits of a number.
// If the number is negative, the first digit is treated as negative.
// e.g., for -123, the sum is -1 + 2 + 3 = 4.
func digitsSum(n int) int {
	// Handle the sign
	sign := 1
	if n < 0 {
		n = -n
		sign = -1
	}

	// Convert the absolute value of the number to a string
	s := strconv.Itoa(n)
	sum := 0

	// Iterate over the string representation of the number
	for i, r := range s {
		// Convert rune to integer
		digit := int(r - '0')
		if i == 0 {
			// Apply the sign to the first digit
			sum += digit * sign
		} else {
			sum += digit
		}
	}
	return sum
}

// CountNums takes an array of integers and returns
// the number of elements which has a sum of digits > 0.
func CountNums(arr []int) int {
	count := 0
	for _, num := range arr {
		if digitsSum(num) > 0 {
			count++
		}
	}
	return count
}

// main function to run and test the implementation with examples.
func main() {
	// Examples from the original Python docstring
	fmt.Println("count_nums([]) == 0 ->", CountNums([]int{}) == 0)
	fmt.Println("count_nums([-1, 11, -11]) == 1 ->", CountNums([]int{-1, 11, -11}) == 1)
	fmt.Println("count_nums([1, 1, 2]) == 3 ->", CountNums([]int{1, 1, 2}) == 3)
}
