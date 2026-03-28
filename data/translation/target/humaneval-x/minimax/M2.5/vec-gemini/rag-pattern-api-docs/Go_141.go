package main

import (
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

	if !unicode.IsLetter(rune(lst[0][0])) {
		return "No"
	}

	digitCount := 0
	for _, c := range lst[0] {
		if unicode.IsDigit(rune(c)) {
			digitCount++
		}
	}
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
