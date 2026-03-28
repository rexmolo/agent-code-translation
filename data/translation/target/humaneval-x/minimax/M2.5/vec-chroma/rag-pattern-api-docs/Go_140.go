package main

import "fmt"

func FixSpaces(text string) string {
	newText := ""
	i := 0
	start, end := 0, 0

	for i < len(text) {
		if text[i] == ' ' {
			end++
		} else {
			spaceCount := end - start
			if spaceCount > 2 {
				newText += "-" + string(text[i])
			} else if spaceCount > 0 {
				for j := 0; j < spaceCount; j++ {
					newText += "_"
				}
				newText += string(text[i])
			} else {
				newText += string(text[i])
			}
			start = i + 1
			end = i + 1
		}
		i++
	}

	// Handle trailing spaces after the loop
	spaceCount := end - start
	if spaceCount > 2 {
		newText += "-"
	} else if spaceCount > 0 {
		for j := 0; j < spaceCount; j++ {
			newText += "_"
		}
	}

	return newText
}

func main() {
	// Test cases
	testCases := []string{
		"Example",
		"Example 1",
		" Example 2",
		" Example   3",
	}

	for _, tc := range testCases {
		result := FixSpaces(tc)
		fmt.Printf("FixSpaces(%q) = %q\n", tc, result)
	}
}
