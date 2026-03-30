package main

import (
    "fmt"
    "strings"
)

func FileNameCheck(file_name string) string {
	suf := []string{"txt", "exe", "dll"}
	lst := strings.Split(file_name, ".")

	// Check if there is exactly one dot (resulting in 2 parts)
	if len(lst) != 2 {
		return "No"
	}

	// Check if extension is one of the allowed suffixes
	extValid := false
	for _, s := range suf {
		if lst[1] == s {
			extValid = true
			break
		}
	}
	if !extValid {
		return "No"
	}

	// Check if the substring before the dot is not empty
	if len(lst[0]) == 0 {
		return "No"
	}

	// Check if it starts with a letter from the latin alphabet
	firstChar := lst[0][0]
	if !((firstChar >= 'a' && firstChar <= 'z') || (firstChar >= 'A' && firstChar <= 'Z')) {
		return "No"
	}

	// Count digits in the prefix (replacing list comprehension)
	digitCount := 0
	for _, c := range lst[0] {
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
	fmt.Println(FileNameCheck("example.txt"))   // Yes
	fmt.Println(FileNameCheck("1example.dll"))   // No
}