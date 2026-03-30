package main

import (
	"strings"
	"unicode"
)

func FileNameCheck(file_name string) string {
	suffixes := []string{"txt", "exe", "dll"}

	// Split by dot and check we get exactly 2 parts
	parts := strings.Split(file_name, ".")
	if len(parts) != 2 {
		return "No"
	}

	// Check if suffix is in the allowed list
	validSuffix := false
	for _, s := range suffixes {
		if parts[1] == s {
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
	for _, c := range parts[0] {
		if unicode.IsDigit(c) {
			digitCount++
		}
	}
	if digitCount > 3 {
		return "No"
	}

	return "Yes"
}
