package main

import (
	"strings"
	"unicode"
)

func SplitWords(txt string) interface{} {
	if strings.Contains(txt, " ") {
		return strings.Split(txt, " ")
	} else if strings.Contains(txt, ",") {
		return strings.Split(strings.ReplaceAll(txt, ",", " "), " ")
	} else {
		count := 0
		for _, r := range txt {
			if unicode.IsLower(r) {
				pos := int(r) - int('a')
				if pos%2 == 0 {
					count++
				}
			}
		}
		return count
	}
}

func main() {
	// Example usage
	result1 := SplitWords("Hello world!")
	println(result1)
	result2 := SplitWords("Hello,world!")
	println(result2)
	result3 := SplitWords("abcdef")
	println(result3)
}
