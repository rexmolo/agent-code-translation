package main

// SameChars checks if two words have the same characters.
func SameChars(s0 string, s1 string) bool {
	set0 := make(map[rune]struct{})
	set1 := make(map[rune]struct{})

	for _, r := range s0 {
		set0[r] = struct{}{}
	}

	for _, r := range s1 {
		set1[r] = struct{}{}
	}

	if len(set0) != len(set1) {
		return false
	}

	for k := range set0 {
		if _, exists := set1[k]; !exists {
			return false
		}
	}

	return true
}

func main() {
	// Example usage
	println(SameChars("eabcdzzzz", "dddzzzzzzzddeddabc")) // true
	println(SameChars("abcd", "dddddddabc"))             // true
	println(SameChars("dddddddabc", "abcd"))             // true
	println(SameChars("eabcd", "dddddddabc"))            // false
	println(SameChars("abcd", "dddddddabce"))           // false
	println(SameChars("eabcdzzzz", "dddzzzzzzzddddabc")) // false
}
