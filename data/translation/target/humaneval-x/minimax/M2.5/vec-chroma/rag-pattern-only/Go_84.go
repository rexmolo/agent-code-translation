package main

import (
	"strconv"
)

func Solve(N int) string {
	// Convert N to string to iterate over each digit
	s := strconv.Itoa(N)
	
	// Sum all digits
	sum := 0
	for _, c := range s {
		sum += int(c - '0')
	}
	
	// Convert sum to binary string (base 2)
	return strconv.FormatInt(int64(sum), 2)
}
