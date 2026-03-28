func ProdSigns(arr []int) interface{} {
	if len(arr) == 0 {
		return nil
	}

	// Check if 0 is in arr (product of signs is 0 if any element is 0)
	hasZero := false
	for _, v := range arr {
		if v == 0 {
			hasZero = true
			break
		}
	}

	// Count negative numbers to determine product sign
	negCount := 0
	for _, v := range arr {
		if v < 0 {
			negCount++
		}
	}

	// Calculate product of signs: 0 if hasZero, otherwise (-1)^negCount
	prod := 0
	if !hasZero {
		if negCount%2 == 0 {
			prod = 1
		} else {
			prod = -1
		}
	}

	// Calculate sum of absolute values
	sum := 0
	for _, v := range arr {
		if v < 0 {
			sum += -v
		} else {
			sum += v
		}
	}

	return prod * sum
}