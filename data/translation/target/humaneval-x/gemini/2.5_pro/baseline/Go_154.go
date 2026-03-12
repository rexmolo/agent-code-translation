package main

import (
	"fmt"
	"strings"
)

// CycpatternCheck checks if the second word or any of its rotations is a substring in the first word.
func CycpatternCheck(a, b string) bool {
	lenB := len(b)
	lenA := len(a)

	if lenB == 0 {
		return true
	}
	// A rotation of b cannot be a substring of a if b is longer than a.
	if lenB > lenA {
		return false
	}

	// Concatenate b with itself. This new string `pat` contains all
	// possible cyclic rotations of `b` as substrings.
	// For example, if b="abc", pat="abcabc". The rotations "abc", "bca", "cab"
	// are all substrings of pat.
	pat := b + b

	// Now we check if any substring of 'a' of length lenB is a rotation of 'b'.
	// A string 's' is a rotation of 'b' if and only if len(s) == len(b)
	// and 's' is a substring of 'b' + 'b' (our `pat` string).
	// This mirrors the logic of the original Python code.
	for i := 0; i <= lenA-lenB; i++ {
		substringA := a[i : i+lenB]
		if strings.Contains(pat, substringA) {
			return true
		}
	}

	return false
}

// main function to demonstrate the usage of CycpatternCheck.
func main() {
	fmt.Println(CycpatternCheck("abcd", "abd"))
	fmt.Println(CycpatternCheck("hello", "ell"))
	fmt.Println(CycpatternCheck("whassup", "psus"))
	fmt.Println(CycpatternCheck("abab", "baa"))
	fmt.Println(CycpatternCheck("efef", "eeff"))
	fmt.Println(CycpatternCheck("himenss", "simen"))
}
