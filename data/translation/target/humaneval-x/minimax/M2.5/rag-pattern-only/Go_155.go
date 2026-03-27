package main

import (
	"strconv"
)

func EvenOddCount(num int) [2]int {
	evenCount := 0
	oddCount := 0

	// Get absolute value
	if num < 0 {
		num = -num
	}

	// Convert number to string
	str := strconv.Itoa(num)

	// Iterate through each character (digit)
	for _, c := range str {
		digit := int(c - '0') // Convert rune to integer digit
		if digit%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	return [2]int{evenCount, oddCount}
}

func main() {
	// Example usage
	println(EvenOddCount(-12)[0], EvenOddCount(-12)[1]) // Output: 1 1
	println(EvenOddCount(123)[0], EvenOddCount(123)[1])  // Output: 1 2
}
