func IsSorted(lst []int) bool {
	// Count occurrences of each number
	countDigit := make(map[int]int)
	for _, n := range lst {
		countDigit[n]++
	}

	// Check if any number appears more than 2 times (more than 1 duplicate)
	for _, n := range lst {
		if countDigit[n] > 2 {
			return false
		}
	}

	// Check if sorted in ascending order
	for i := 1; i < len(lst); i++ {
		if lst[i-1] > lst[i] {
			return false
		}
	}

	return true
}