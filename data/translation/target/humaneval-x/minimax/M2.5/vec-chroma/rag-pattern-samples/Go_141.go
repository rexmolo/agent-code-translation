package main

import (
	"fmt"
	"strings"
	"unicode"
)

func FileNameCheck(fileName string) string {
	suf := []string{"txt", "exe", "dll"}
	lst := strings.Split(fileName, ".")

	if len(lst) != 2 {
		return "No"
	}

	if !contains(suf, lst[1]) {
		return "No"
	}

	if len(lst[0]) == 0 {
		return "No"
	}

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

	if digitCount > 3 {
		return "No"
	}

	return "Yes"
}

// Helper function to check if a slice contains an element
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func main() {
	// Test cases
	fmt.Println(FileNameCheck("example.txt"))   // Expected: Yes
	fmt.Println(FileNameCheck("1example.dll")) // Expected: No
	fmt.Println(FileNameCheck("test.exe"))     // Expected: Yes
	fmt.Println(FileNameCheck("test.cpp"))     // Expected: No
	fmt.Println(FileNameCheck(".txt"))         // Expected: No (empty prefix)
	fmt.Println(FileNameCheck("test.txte"))    // Expected: No (no dot)
}
