package main

import (
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
		for _, ch := range txt {
			if ch >= 'a' && ch <= 'z' {
				// ord('a') = 0, ord('b') = 1, ... ord('z') = 25
				// odd positions in alphabet: a(0), c(2), e(4), ... have ord % 2 == 0
				if int(ch-'a')%2 == 0 {
					count++
				}
			}
		}
		return count
	}
}