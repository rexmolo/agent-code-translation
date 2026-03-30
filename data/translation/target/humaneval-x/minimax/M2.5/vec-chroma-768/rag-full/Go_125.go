package main

import (
	"fmt"
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
			if unicode.IsLower(r) && int(r-'a')%2 == 0 {
				count++
			}
		}
		return count
	}
}

func main() {
	// Test cases from examples
	fmt.Println(SplitWords("Hello world!"))
	fmt.Println(SplitWords("Hello,world!"))
	fmt.Println(SplitWords("abcdef"))
}