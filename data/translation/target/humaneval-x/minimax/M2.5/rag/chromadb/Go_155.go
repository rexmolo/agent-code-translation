package main

import (
	"math"
	"strconv"
)

func EvenOddCount(num int) [2]int {
	n := int(math.Abs(float64(num)))
	s := strconv.Itoa(n)

evenCount := 0
	oddCount := 0

	for _, c := range s {
		digit := int(c) - '0'
		if digit%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	return [2]int{evenCount, oddCount}
}
