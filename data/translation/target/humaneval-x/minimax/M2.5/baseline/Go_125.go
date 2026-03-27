package main

import (
	"strings"
	"unicode"
)

func SplitWords(txt string) interface{} {
	if strings.Contains(txt, " ") {
		return strings.Fields(txt)
	} else if strings.Contains(txt, ",") {
		return strings.Split(strings.ReplaceAll(txt, ",", " "), " ")
	} else {
		count := 0
		for _, ch := range txt {
			if unicode.IsLower(ch) && int(ch)%2 == 0 {
				count++
			}
		}
		return count
	}
}

func main() {
	// Example usage - can be tested interactively
}