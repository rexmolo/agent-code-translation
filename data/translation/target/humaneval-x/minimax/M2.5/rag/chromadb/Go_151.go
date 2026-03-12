package main

import (
	"math"
)

func DoubleTheDifference(lst []float64) int {
	sum := 0
	for _, n := range lst {
		// Skip non-positive numbers
		if n <= 0 {
			continue
		}
		
		// Check if it's an integer (not a float)
		// A number is an integer if its floor equals itself
		if math.Floor(n) != n {
			continue
		}
		
		// Check if it's odd
		// Convert to int64 for modulo operation
		if int64(n)%2 == 0 {
			continue
		}
		
		// Add square to sum
		square := n * n
		sum += int(square)
	}
	return sum
}

// Alternative version without math package:
func DoubleTheDifferenceAlt(lst []float64) int {
	sum := 0
	for _, n := range lst {
		// Skip non-positive numbers
		if n <= 0 {
			continue
		}
		
		// Check if it's an integer by converting to string
		// and looking for a decimal point
		s := strconv.FormatFloat(n, 'f', -1, 64)
		if strings.Contains(s, ".") {
			continue
		}
		
		// Check if it's odd
		if int64(n)%2 == 0 {
			continue
		}
		
		// Add square to sum
		square := n * n
		sum += int(square)
	}
	return sum
}