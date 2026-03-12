package main

import (
	"fmt"
	"strings"
	"unicode"
)

func FileNameCheck(file_name string) string {
	// Allowed suffixes
	suffixes := []string{"txt", "exe", "dll"}

	// Split by dot
	parts := strings.Split(file_name, ".")

	// Must have exactly 2 parts (one dot)
	if len(parts) != 2 {
		return "No"
	}

	// Check suffix
	validSuffix := false
	for _, suf := range suffixes {
		if parts[1] == suf {
			validSuffix = true
			break
		}
	}
	if !validSuffix {
		return "No"
	}

	// Prefix must not be empty
	if len(parts[0]) == 0 {
		return "No"
	}

	// First character must be a letter
	firstChar := rune(parts[0][0])
	if !unicode.IsLetter(firstChar) {
		return "No"
	}

	// Count digits in prefix
	digitCount := 0
	for _, ch := range parts[0] {
		if unicode.IsDigit(ch) {
			digitCount++
		}
	}

	// No more than 3 digits
	if digitCount > 3 {
		return "No"
	}

	return "Yes"
}

func main() {
	fmt.Println(FileNameCheck("example.txt"))   // Yes
	fmt.Println(FileNameCheck("1example.dll"))  // No
}