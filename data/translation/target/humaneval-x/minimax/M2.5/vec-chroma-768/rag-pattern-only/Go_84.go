package main

import (
	"strconv"
)

func Solve(N int) string {
	// Convert N to string to iterate over each digit
	strN := strconv.Itoa(N)
	
	// Sum all the digits
	sum := 0
	for _, c := range strN {
		sum += int(c - '0')
	}
	
	// Convert sum to binary and return as string
	return strconv.FormatInt(int64(sum), 2)
}
