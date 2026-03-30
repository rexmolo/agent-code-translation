package main

func SameChars(s0 string, s1 string) bool {
	// Convert s0 to a set of unique characters
	set0 := make(map[rune]bool)
	for _, r := range s0 {
		set0[r] = true
	}

	// Convert s1 to a set of unique characters
	set1 := make(map[rune]bool)
	for _, r := range s1 {
		set1[r] = true
	}

	// First check if they have the same number of unique characters
	if len(set0) != len(set1) {
		return false
	}

	// Check if all characters in set0 exist in set1
	for r := range set0 {
		if !set1[r] {
			return false
		}
	}

	return true
}

func main() {
	// Test cases (optional)
	// SameChars("eabcdzzzz", "dddzzzzzzzddeddabc") // true
	// SameChars("abcd", "dddddddabc")              // true
	// SameChars("eabcd", "dddddddabc")             // false
}
