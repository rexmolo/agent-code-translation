package main

func IsSorted(lst []int) bool {
	// Count occurrences of each number
	count := make(map[int]int)
	for _, v := range lst {
		count[v]++
	}

	// Check if any number appears more than twice (more than 1 duplicate)
	for _, v := range lst {
		if count[v] > 2 {
			return false
		}
	}

	// Check if the slice is sorted in ascending order
	for i := 1; i < len(lst); i++ {
		if lst[i-1] > lst[i] {
			return false
		}
	}

	return true
}
