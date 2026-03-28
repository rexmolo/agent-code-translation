package main

import (
	"fmt"
	"strings"
	"unicode"
)

func FileNameCheck(file_name string) string {
	suf := []string{"txt", "exe", "dll"}
	lst := strings.Split(file_name, ".")

	// There should be exactly one dot (i.e., two parts after splitting)
	if len(lst) != 2 {
		return "No"
	}

	// The substring after the dot should be one of ['txt', 'exe', 'dll']
	found := false
	for _, s := range suf {
		if lst[1] == s {
			found = true
			break
		}
	}
	if !found {
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

func main() {
	// Test examples
	fmt.Println(FileNameCheck("example.txt"))  // Yes
	fmt.Println(FileNameCheck("1example.dll")) // No
}