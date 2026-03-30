package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func FileNameCheck(fileName string) string {
	// Valid suffixes
	suf := []string{"txt", "exe", "dll"}

	// Split by '.'
	lst := strings.Split(fileName, ".")

	// Check if there are exactly 2 parts
	if len(lst) != 2 {
		return "No"
	}

	// Check if suffix (after dot) is valid
	validSuffix := false
	for _, s := range suf {
		if lst[1] == s {
			validSuffix = true
			break
		}
	}
	if !validSuffix {
		return "No"
	}

	// Check if name (before dot) is not empty
	if len(lst[0]) == 0 {
		return "No"
	}

	// Check if first character of name is a letter
	firstChar := rune(lst[0][0])
	if !unicode.IsLetter(firstChar) {
		return "No"
	}

	// Count digits in name part
	digitCount := 0
	for _, c := range lst[0] {
		if unicode.IsDigit(c) {
			digitCount++
		}
	}

	// Check if there are more than 3 digits
	if digitCount > 3 {
		return "No"
	}

	return "Yes"
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fileName, _ := reader.ReadString('\n')
	fileName = strings.TrimSpace(fileName)
	fmt.Println(FileNameCheck(fileName))
}
