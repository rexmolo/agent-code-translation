package main

import (
	"strconv"
)

func EvenOddCount(num int) [2]int {
	evenCount := 0
	oddCount := 0

	// Handle negative numbers
	if num < 0 {
		num = -num
	}

	// Special case: 0 has one digit that is even
	if num == 0 {
		return [2]int{1, 0}
	}

	// Convert to string and iterate through each digit
	str := strconv.Itoa(num)

	for _, c := range str {
		digit := c - '0' // Convert rune to integer digit
		if digit%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	return [2]int{evenCount, oddCount}
}

func main() {
	// Test cases
	// Should print [1 1]
	print(EvenOddCount(-12))
	// Should print [1 2]
	print(EvenOddCount(123))
}