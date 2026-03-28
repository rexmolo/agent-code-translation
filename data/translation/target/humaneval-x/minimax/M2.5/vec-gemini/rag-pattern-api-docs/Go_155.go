package main

import (
	"math"
	"strconv"
)

func EvenOddCount(num int) [2]int {
	evenCount := 0
	oddCount := 0

	absNum := int(math.Abs(float64(num)))
	str := strconv.Itoa(absNum)

	for _, c := range str {
		digit := int(c - '0')
		if digit%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	return [2]int{evenCount, oddCount}
}