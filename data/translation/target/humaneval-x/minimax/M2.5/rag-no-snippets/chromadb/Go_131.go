package main

import (
	"strconv"
)

func Digits(n int) int {
	product := 1
	oddCount := 0

	// Convert n to string to iterate over digits
	strN := strconv.Itoa(n)

	for _, char := range strN {
		// Convert rune to digit (int) by subtracting the ASCII value of '0'
		digit := int(char - '0')
		if digit%2 == 1 {
			product *= digit
			oddCount++
		}
	}

	if oddCount == 0 {
		return 0
	}
	return product
}

func main() {
	// Test cases
	println(Digits(1))   // 1
	println(Digits(4))   // 0
	println(Digits(235)) // 15
}