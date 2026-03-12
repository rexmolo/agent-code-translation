package main

import (
	"fmt"
	"strings"
)

func SplitWords(txt string) interface{} {
	if strings.Contains(txt, " ") {
		return strings.Fields(txt)
	} else if strings.Contains(txt, ",") {
		replaced := strings.ReplaceAll(txt, ",", " ")
		return strings.Fields(replaced)
	} else {
		count := 0
		for _, c := range txt {
			if c >= 'a' && c <= 'z' && (int(c)-int('a'))%2 == 0 {
				count++
			}
		}
		return count
	}
}

func main() {
	// Test cases
	result1 := SplitWords("Hello world!")
	fmt.Println(result1)

	result2 := SplitWords("Hello,world!")
	fmt.Println(result2)

	result3 := SplitWords("abcdef")
	fmt.Println(result3)
}
