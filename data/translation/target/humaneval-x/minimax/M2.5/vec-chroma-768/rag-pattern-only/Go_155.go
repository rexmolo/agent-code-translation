package main

import (
	"fmt"
	"strconv"
)

func EvenOddCount(num int) [2]int {
	evenCount := 0
	oddCount := 0

	absNum := num
	if absNum < 0 {
		absNum = -absNum
	}

	for _, c := range []byte(strconv.Itoa(absNum)) {
		digit := int(c - '0')
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
	result1 := EvenOddCount(-12)
	fmt.Printf("EvenOddCount(-12) => [%d, %d]\n", result1[0], result1[1])

	result2 := EvenOddCount(123)
	fmt.Printf("EvenOddCount(123) => [%d, %d]\n", result2[0], result2[1])
}
