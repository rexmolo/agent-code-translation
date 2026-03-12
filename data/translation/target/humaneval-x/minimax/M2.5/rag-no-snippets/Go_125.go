package main

import (
	"strings"
)

func SplitWords(txt string) interface{} {
	if strings.Contains(txt, " ") {
		return strings.Fields(txt)
	} else if strings.Contains(txt, ",") {
		return strings.Fields(strings.ReplaceAll(txt, ",", " "))
	} else {
		count := 0
		for _, r := range txt {
			if r >= 'a' && r <= 'z' && int(r)%2 == 0 {
				count++
			}
		}
		return count
	}
}

func main() {
	// Example usage - these can be tested manually
	_ = SplitWords("Hello world!")
	_ = SplitWords("Hello,world!")
	_ = SplitWords("abcdef")
}