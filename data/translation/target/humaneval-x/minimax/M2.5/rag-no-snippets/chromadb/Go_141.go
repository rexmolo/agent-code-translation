package main

import (
	"fmt"
	"strings"
	"unicode"
)

func FileNameCheck(fileName string) string {
	suffixes := []string{"txt", "exe", "dll"}

	// Split by dot - should result in exactly 2 parts (before and after the dot)
	lst := strings.Split(fileName, ".")
	if len(lst) != 2 {
		return "No"
	}

	// Check if suffix is valid
	suffixValid := false
	for _, s := range suffixes {
		if lst[1] == s {
			suffixValid = true
			break
		}
	}
	if !suffixValid {
		return "No"
	}

	// Check if there's something before the dot
	if len(lst[0]) == 0 {
		return "No"
	}

	// Check if first character is a letter (a-z or A-Z)
	firstChar := rune(lst[0][0])
	if !unicode.IsLetter(firstChar) {
		return "No"
	}

	// Count digits in the part before the dot
	digitCount := 0
	for _, c := range lst[0] {
		if unicode.IsDigit(c) {
			digitCount++
		}
	}

	// Check not more than 3 digits
	if digitCount > 3 {
		return "No"
	}

	return "Yes"
}

func main() {
	fmt.Println(FileNameCheck("example.txt"))   // Yes
	fmt.Println(FileNameCheck("1example.dll"))   // No
}