package main

import (
	"fmt"
	"strings"
	"unicode"
)

func FileNameCheck(file_name string) string {
	// A file's name contains exactly one dot '.'
	parts := strings.Split(file_name, ".")
	if len(parts) != 2 {
		return "No"
	}

	prefix := parts[0]
	suffix := parts[1]

	// The substring after the dot should be one of these: ['txt', 'exe', 'dll']
	validSuffixes := map[string]bool{"txt": true, "exe": true, "dll": true}
	if !validSuffixes[suffix] {
		return "No"
	}

	// The substring before the dot should not be empty
	if len(prefix) == 0 {
		return "No"
	}

	// It starts with a letter from the latin alphapet ('a'-'z' and 'A'-'Z')
	firstChar := rune(prefix[0])
	if !((firstChar >= 'a' && firstChar <= 'z') || (firstChar >= 'A' && firstChar <= 'Z')) {
		return "No"
	}

	// There should not be more than three digits ('0'-'9') in the file's name.
	// The original Python code only counts digits in the prefix, so we do the same.
	digitCount := 0
	for _, char := range prefix {
		if unicode.IsDigit(char) {
			digitCount++
		}
	}

	if digitCount > 3 {
		return "No"
	}

	return "Yes"
}

func main() {
	fmt.Println(FileNameCheck("example.txt"))
	fmt.Println(FileNameCheck("1example.dll"))
}