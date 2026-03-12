func Pluck(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	// Find the smallest even value and its index
	minEven := -1
	minIndex := -1

	for i, val := range arr {
		if val%2 == 0 {
			// This is an even number
			if minEven == -1 || val < minEven {
				minEven = val
				minIndex = i
			}
		}
	}

	if minEven == -1 {
		return []int{}
	}

	return []int{minEven, minIndex}
}