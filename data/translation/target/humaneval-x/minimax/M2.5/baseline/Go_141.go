package main

import (
	"fmt"
	"strings"
)

func FileNameCheck(fileName string) string {
	suffixes := []string{"txt", "exe", "dll"}

	// Split by dot - should result in exactly 2 parts
	lst := strings.Split(fileName, ".")
	if len(lst) != 2 {
		return "No"
	}

	// The substring after the dot should be one of ['txt', 'exe', 'dll']
	if !contains(suffixes, lst[1]) {
		return "No"
	}

	// The substring before the dot should not be empty
	if len(lst[0]) == 0 {
		return "No"
	}

	// The substring before the dot should start with a latin alphabet letter
	firstChar := lst[0][0]
	if !isLatinLetter(firstChar) {
		return "No"
	}

	// Count digits in the substring before the dot
	digitCount := 0
	for _, c := range lst[0] {
		if c >= '0' && c <= '9' {
			digitCount++
		}
	}

	// There should not be more than three digits
	if digitCount > 3 {
		return "No"
	}

	return "Yes"
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func isLatinLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func main() {
	fmt.Println(FileNameCheck("example.txt"))   // => Yes
	fmt.Println(FileNameCheck("1example.dll"))  // => No
}