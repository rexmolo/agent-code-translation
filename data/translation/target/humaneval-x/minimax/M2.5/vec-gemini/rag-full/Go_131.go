package main

import "strconv"

func Digits(n int) int {
	product := 1
	oddCount := 0

	// Convert integer to string to iterate over digits
	s := strconv.Itoa(n)

	for _, c := range s {
		digit := int(c - '0') // Convert character to numeric digit
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