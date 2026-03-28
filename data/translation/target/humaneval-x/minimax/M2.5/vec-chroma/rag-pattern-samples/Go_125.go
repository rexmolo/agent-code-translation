package main

import (
	"fmt"
	"strings"
)

func SplitWords(txt string) interface{} {
	if strings.Contains(txt, " ") {
		return strings.Fields(txt)
	} else if strings.Contains(txt, ",") {
		return strings.Split(strings.ReplaceAll(txt, ",", " "), " ")
	} else {
		count := 0
		for _, c := range txt {
			if c >= 'a' && c <= 'z' {
				pos := int(c - 'a')
				if pos%2 == 0 {
					count++
				}
			}
		}
		return count
	}
}

func main() {
	// Test cases
	fmt.Println(SplitWords("Hello world!"))   // [Hello world!]
	fmt.Println(SplitWords("Hello,world!"))   // [Hello world!]
	fmt.Println(SplitWords("abcdef"))         // 3
}