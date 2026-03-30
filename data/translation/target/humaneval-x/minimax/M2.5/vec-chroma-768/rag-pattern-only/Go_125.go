package main

import (
	"strings"
)

func SplitWords(txt string) interface{} {
	// Check if there's a whitespace in the text
	if strings.Contains(txt, " ") {
		return strings.Fields(txt)
	}

	// Check if there's a comma in the text
	if strings.Contains(txt, ",") {
		return strings.Fields(strings.ReplaceAll(txt, ",", " "))
	}

	// Count lowercase letters with even order in the alphabet (a=0, c=2, e=4, ...)
	count := 0
	for _, ch := range txt {
		if ch >= 'a' && ch <= 'z' && (int(ch-'a'))%2 == 0 {
			count++
		}
	}
	return count
}

func main() {
	// Example usage - can be tested with stdin/stdout if needed
}