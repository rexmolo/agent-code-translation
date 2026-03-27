package main

import "math"

func ProdSigns(arr []int) interface{} {
	if len(arr) == 0 {
		return nil
	}

	hasZero := false
	negCount := 0
	sumAbs := 0

	for _, v := range arr {
		if v == 0 {
			hasZero = true
		}
		if v < 0 {
			negCount++
		}
		sumAbs += int(math.Abs(float64(v)))
	}

	var prod int
	if hasZero {
		prod = 0
	} else {
		// (-1) ^ negCount
		if negCount%2 == 0 {
			prod = 1
		} else {
			prod = -1
		}
	}

	return prod * sumAbs
}