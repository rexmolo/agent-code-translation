package main

import (
	"fmt"
	"strings"
)

func FileNameCheck(fileName string) string {
	// Valid extensions
	validSuffixes := []string{"txt", "exe", "dll"}

	// Split by dot - should result in exactly 2 parts
	parts := strings.Split(fileName, ".")
	if len(parts) != 2 {
		return "No"
	}

	// The substring after the dot should be one of ['txt', 'exe', 'dll']
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

	// The substring before the dot should not be empty
	if len(parts[0]) == 0 {
		return "No"
	}

	// The first character should be a letter from the latin alphabet
	firstChar := rune(parts[0][0])
	if !((firstChar >= 'a' && firstChar <= 'z') || (firstChar >= 'A' && firstChar <= 'Z')) {
		return "No"
	}

	// Count digits in the substring before the dot
	digitCount := 0
	for _, c := range parts[0] {
		if c >= '0' && c <= '9' {
			digitCount++
		}
	}
	if digitCount > 3 {
		return "No"
	}

	return "Yes"
}

func main() {
	var fileName string
	fmt.Scan(&fileName)
	fmt.Println(FileNameCheck(fileName))
}
