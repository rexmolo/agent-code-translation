package main

import "strings"

func FixSpaces(text string) string {
	newText := ""
	start := 0
	end := 0

	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			end++
		} else {
			if end-start > 2 {
				newText += "-" + string(text[i])
			} else if end-start > 0 {
				newText += strings.Repeat("_", end-start) + string(text[i])
			} else {
				newText += string(text[i])
			}
			start = i + 1
			end = i + 1
		}
	}

	if end-start > 2 {
		newText += "-"
	} else if end-start > 0 {
		newText += "_"
	}

	return newText
}