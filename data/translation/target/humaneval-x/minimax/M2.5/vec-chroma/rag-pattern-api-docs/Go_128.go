package main

import "math"

// ProdSigns calculates the sum of absolute values multiplied by the product of all signs.
// Returns nil for empty array, 0 if any element is 0, or the product of -1 raised to
// the power of negative count times the sum of magnitudes.
func ProdSigns(arr []int) interface{} {
	if len(arr) == 0 {
		return nil
	}

	// Check if 0 is in arr and count negative numbers
	hasZero := false
	negativeCount := 0
	for _, v := range arr {
		if v == 0 {
			hasZero = true
		}
		if v < 0 {
			negativeCount++
		}
	}

	// Calculate product of signs
	// If 0 is in arr, product is 0
	// Otherwise, product is (-1) ^ negativeCount
	var prod int
	if hasZero {
		prod = 0
	} else if negativeCount%2 == 1 {
		prod = -1
	} else {
		prod = 1
	}

	// Calculate sum of absolute values
	sumAbs := 0
	for _, v := range arr {
		sumAbs += int(math.Abs(float64(v)))
	}

	return prod * sumAbs
}