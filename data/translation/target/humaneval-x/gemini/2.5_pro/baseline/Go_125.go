package main

import (
	"fmt"
	"strings"
)

func SplitWords(txt string) interface{} {
	if strings.Contains(txt, " ") {
		return strings.Fields(txt)
	} else if strings.Contains(txt, ",") {
		// This mimics Python's `txt.replace(',',' ').split()`
		replaced := strings.ReplaceAll(txt, ",", " ")
		return strings.Fields(replaced)
	} else {
		count := 0
		// The original Python code counts lowercase letters where ord(i)%2 == 0.
		// This corresponds to letters like 'b', 'd', 'f', etc.
		// Go's rune is an integer type (int32), so we can directly check its value.
		for _, char := range txt {
			if char >= 'a' && char <= 'z' && char%2 == 0 {
				count++
			}
		}
		return count
	}
}

func main() {
	// Example 1: split on spaces
	result1 := SplitWords("Hello world!")
	fmt.Printf("SplitWords(\"Hello world!\") -> %v\n", result1)

	// Example 2: split on commas
	result2 := SplitWords("Hello,world!")
	fmt.Printf("SplitWords(\"Hello,world!\") -> %v\n", result2)

	// Example 3: count letters
	result3 := SplitWords("abcdef")
	fmt.Printf("SplitWords(\"abcdef\") -> %v\n", result3)

	// Additional test cases
	result4 := SplitWords("one,two,,three")
	fmt.Printf("SplitWords(\"one,two,,three\") -> %v\n", result4)

	result5 := SplitWords("no_special_chars")
	fmt.Printf("SplitWords(\"no_special_chars\") -> %v\n", result5)
}