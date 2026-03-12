package main

// countUniqueChars is a helper function that calculates the number of unique
// runes (characters) in a given string.
func countUniqueChars(s string) int {
	// A map with an empty struct as value is an efficient way to implement a set in Go.
	unique := make(map[rune]struct{})
	for _, r := range s {
		unique[r] = struct{}{}
	}
	return len(unique)
}

// FindMax accepts a slice of strings and returns the word with the maximum
// number of unique characters. If there's a tie, it returns the one that
// comes first lexicographically.
func FindMax(words []string) string {
	// Handle the edge case of an empty slice, returning an empty string.
	if len(words) == 0 {
		return ""
	}

	// Initialize the best-so-far with the first word in the slice.
	maxWord := words[0]
	maxCount := countUniqueChars(maxWord)

	// Iterate through the rest of the words to find a better candidate.
	for i := 1; i < len(words); i++ {
		word := words[i]
		count := countUniqueChars(word)

		// A new word is better if it has more unique characters...
		if count > maxCount {
			maxCount = count
			maxWord = word
		} else if count == maxCount {
			// ...or if it has the same number of unique characters but is
			// lexicographically smaller (the tie-breaker).
			if word < maxWord {
				maxWord = word
			}
		}
	}

	return maxWord
}
