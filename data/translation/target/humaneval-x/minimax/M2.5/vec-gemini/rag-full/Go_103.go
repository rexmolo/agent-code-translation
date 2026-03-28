package main

import (
	"fmt"
	"math"
)

func RoundedAvg(n, m int) interface{} {
	if m < n {
		return -1
	}
	
	// Calculate sum of integers from n to m (inclusive)
	sum := 0
	for i := n; i <= m; i++ {
		sum += i
	}
	
	count := m - n + 1
	
	// Calculate average
	avg := float64(sum) / float64(count)
	
	// Round the average
	// Python's round() uses banker's rounding (round half to even)
	// Go's math.Round uses "round half away from zero"
	// We need to implement banker's rounding for .5 cases
	rounded := int(math.Round(avg))
	
	// Check if we hit the .5 case exactly
	if math.Abs(avg-float64(rounded)) == 0.5 {
		// Banker's rounding: round to nearest even
		if rounded%2 == 1 {
			rounded--
		} else {
			rounded++
		}
	}
	
	// Convert to binary string
	return fmt.Sprintf("0b%d", rounded)
}

func main() {
	// Test cases
	testCases := []struct {
		n, m int
	}{
		{1, 5},
		{7, 5},
		{10, 20},
		{20, 33},
	}
	
	for _, tc := range testCases {
		result := RoundedAvg(tc.n, tc.m)
		fmt.Printf("rounded_avg(%d, %d) => %v\n", tc.n, tc.m, result)
	}
}
