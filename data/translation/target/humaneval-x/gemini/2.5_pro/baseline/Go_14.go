package main

// AllPrefixes returns a slice of all prefixes of a string, from shortest to longest.
func AllPrefixes(str string) []string {
	// In Go, strings are sequences of bytes. To handle Unicode characters correctly
	// (as Python does by default), we work with runes.
	runes := []rune(str)

	// Pre-allocate the slice since we know the final size.
	result := make([]string, len(runes))

	for i := range runes {
		// Slice the runes and convert the sub-slice back to a string.
		result[i] = string(runes[:i+1])
	}

	return result
}
