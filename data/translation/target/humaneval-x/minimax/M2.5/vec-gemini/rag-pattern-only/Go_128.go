package main

func ProdSigns(arr []int) interface{} {
	if len(arr) == 0 {
		return nil
	}

	// Check if 0 is in arr
	hasZero := false
	for _, v := range arr {
		if v == 0 {
			hasZero = true
			break
		}
	}

	// Count negative numbers
	negativeCount := 0
	for _, v := range arr {
		if v < 0 {
			negativeCount++
		}
	}

	// Calculate product of signs
	var prod int
	if hasZero {
		prod = 0
	} else {
		// (-1) ^ negativeCount
		if negativeCount%2 == 0 {
			prod = 1
		} else {
			prod = -1
		}
	}

	// Calculate sum of absolute values
	sumAbs := 0
	for _, v := range arr {
		if v < 0 {
			sumAbs += -v
		} else {
			sumAbs += v
		}
	}

	return prod * sumAbs
}