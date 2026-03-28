package main

import (
	"strings"
	"unicode"
)

func SplitWords(txt string) interface{} {
	if strings.Contains(txt, " ") {
		return strings.Fields(txt)
	} else if strings.Contains(txt, ",") {
		return strings.Fields(strings.ReplaceAll(txt, ",", " "))
	} else {
		count := 0
		for _, c := range txt {
			if unicode.IsLower(c) && int(c)%2 == 0 {
				count++
			}
		}
		return count
	}
}