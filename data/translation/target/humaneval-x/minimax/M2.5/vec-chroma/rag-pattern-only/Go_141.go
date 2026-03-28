package main

import (
	"strings"
	"unicode"
)

func FileNameCheck(fileName string) string {
	validSuffixes := []string{"txt", "exe", "dll"}

	parts := strings.Split(fileName, ".")
	if len(parts) != 2 {
		return "No"
	}

	// Check if suffix is valid
	validSuffix := false
	for _, suf := range validSuffixes {
		if parts[1] == suf {
			validSuffix = true
			break
		}
	}
	if !validSuffix {
		return "No"
	}

	// Check if prefix is not empty
	if len(parts[0]) == 0 {
		return "No"
	}

	// Check if first character is a letter
	firstChar := rune(parts[0][0])
	if !unicode.IsLetter(firstChar) {
		return "No"
	}

	// Count digits in the prefix
	digitCount := 0
	for _, ch := range parts[0] {
		if unicode.IsDigit(ch) {
			digitCount++
		}
	}
	if digitCount > 3 {
		return "No"
	}

	return "Yes"
}
