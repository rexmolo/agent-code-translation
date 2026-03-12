package main

import (
	"fmt"
	"strings"
	"unicode"
)

func FileNameCheck(fileName string) string {
	suf := []string{"txt", "exe", "dll"}
	lst := strings.Split(fileName, ".")
	
	// There should be exactly one dot (resulting in 2 parts)
	if len(lst) != 2 {
		return "No"
	}
	
	// The substring after the dot should be one of ['txt', 'exe', 'dll']
	if !contains(suf, lst[1]) {
		return "No"
	}
	
	// The substring before the dot should not be empty
	if len(lst[0]) == 0 {
		return "No"
	}
	
	// The substring before the dot should start with a letter
	firstChar := rune(lst[0][0])
	if !unicode.IsLetter(firstChar) {
		return "No"
	}
	
	// Count digits in the substring before the dot
	digitCount := 0
	for _, c := range lst[0] {
		if unicode.IsDigit(c) {
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

func main() {
	// Test cases
	fmt.Println(FileNameCheck("example.txt"))   // Yes
	fmt.Println(FileNameCheck("1example.dll")) // No
}
