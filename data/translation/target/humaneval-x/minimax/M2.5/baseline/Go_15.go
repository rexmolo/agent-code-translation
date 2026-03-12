package main

import (
	"strconv"
	"strings"
)

func StringSequence(n int) string {
	// Create a slice to hold the string representations
	parts := make([]string, n+1)

	// Fill the slice with numbers from 0 to n (inclusive)
	for i := 0; i <= n; i++ {
		parts[i] = strconv.Itoa(i)
	}

	// Join with spaces
	return strings.Join(parts, " ")
}
