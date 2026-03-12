func HowManyTimes(str string, substring string) int {
	count := 0

	// Handle edge cases: empty string or empty substring
	if len(substring) == 0 || len(str) == 0 {
		return 0
	}

	// Iterate through each possible starting position
	for i := 0; i <= len(str)-len(substring); i++ {
		// Check if substring matches at position i (including overlapping cases)
		if str[i:i+len(substring)] == substring {
			count++
		}
	}

	return count
}
