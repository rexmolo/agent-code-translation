package main

import "strings"

func FixSpaces(text string) string {
	var newText strings.Builder
	i := 0
	start := 0
	end := 0
	for i < len(text) {
		if text[i] == ' ' {
			end++
		} else {
			spaceCount := end - start
			if spaceCount > 2 {
				newText.WriteString("-")
			} else if spaceCount > 0 {
				newText.WriteString(strings.Repeat("_", spaceCount))
			}
			newText.WriteByte(text[i])
			start = i + 1
			end = i + 1
		}
		i++
	}

	spaceCount := end - start
	if spaceCount > 2 {
		newText.WriteString("-")
	} else if spaceCount > 0 {
		newText.WriteString("_")
	}

	return newText.String()
}
