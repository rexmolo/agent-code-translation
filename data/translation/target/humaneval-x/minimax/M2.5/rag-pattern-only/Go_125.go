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
		for _, c := range txt {
			if c >= 'a' && c <= 'z' && int(c)%2 == 0 {
				count++
			}
		}
		return count
	}
}