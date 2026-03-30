package main

func SameChars(s0 string, s1 string) bool {
	set0 := make(map[rune]struct{})
	for _, c := range s0 {
		set0[c] = struct{}{}
	}

	set1 := make(map[rune]struct{})
	for _, c := range s1 {
		set1[c] = struct{}{}
	}

	if len(set0) != len(set1) {
		return false
	}

	for k := range set0 {
		if _, ok := set1[k]; !ok {
			return false
		}
	}

	return true
}

func main() {
	// Test cases from docstring
	println(SameChars("eabcdzzzz", "dddzzzzzzzddeddabc"), " (expected true)")
	println(SameChars("abcd", "dddddddabc"), " (expected true)")
	println(SameChars("dddddddabc", "abcd"), " (expected true)")
	println(SameChars("eabcd", "dddddddabc"), " (expected false)")
	println(SameChars("abcd", "dddddddabce"), " (expected false)")
	println(SameChars("eabcdzzzz", "dddzzzzzzzddddabc"), " (expected false)")
}