package main

import (
	"fmt"
	"strconv"
)

func EvenOddCount(num int) [2]int {
	evenCount := 0
	oddCount := 0

	// Convert absolute value to string
	numStr := strconv.Itoa(abs(num))

	// Iterate through each character and count even/odd digits
	for _, c := range numStr {
		digit := int(c - '0')
		if digit%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	return [2]int{evenCount, oddCount}
}

// abs returns the absolute value of an integer
func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	// Test the function
	fmt.Println(EvenOddCount(-12)) // [1 1]
	fmt.Println(EvenOddCount(123)) // [1 2]
}