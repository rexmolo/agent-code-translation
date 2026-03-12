package main

import (
	"strconv"
)

// EvenOddCount given an integer. return an array that has the number of even and odd digits respectively.
//
// Example:
//
//	evenOddCount(-12) ==> [1, 1]
//	evenOddCount(123) ==> [1, 2]
func EvenOddCount(num int) [2]int {
	evenCount := 0
	oddCount := 0

	if num < 0 {
		num = -num
	}

	numStr := strconv.Itoa(num)

	for _, digitChar := range numStr {
		digit := int(digitChar - '0')
		if digit%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}
	return [2]int{evenCount, oddCount}
}
