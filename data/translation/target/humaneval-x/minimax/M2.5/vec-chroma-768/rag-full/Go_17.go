package main

import (
	"fmt"
	"strings"
)

func ParseMusic(music_string string) []int {
	noteMap := map[string]int{
		"o":  4,
		"o|": 2,
		".|": 1,
	}

	var result []int
	for _, note := range strings.Split(music_string, " ") {
		if note != "" {
			result = append(result, noteMap[note])
		}
	}

	return result
}

func main() {
	// Test with the example from the docstring
	result := ParseMusic("o o| .| o| o| .| .| .| .| o o")
	fmt.Println(result)
}