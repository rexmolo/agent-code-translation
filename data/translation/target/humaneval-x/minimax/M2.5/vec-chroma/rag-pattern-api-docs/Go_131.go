package main

import (
	"strconv"
)

func Digits(n int) int {
	product := 1
	oddCount := 0

	str := strconv.Itoa(n)
	for _, c := range str {
		digit := int(c - '0')
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