package main

import (
	"fmt"
	"strings"
)

func ParseMusic(musicString string) []int {
	// Map note symbols to their beat values
	noteMap := map[string]int{
		"o":  4, // whole note
		"o|": 2, // half note
		".|": 1, // quarter note
	}

	// Parse the music string and convert to beats
	var result []int
	parts := strings.Split(musicString, " ")
	for _, note := range parts {
		if note != "" {
			// The 'if x' in Python filters out empty strings from split
			// strings.Split can produce empty elements when multiple spaces exist
			result = append(result, noteMap[note])
		}
	}

	return result
}

func main() {
	// Test the function with the example from the docstring
	result := ParseMusic("o o| .| o| o| .| .| .| .| o o")
	fmt.Println(result)
	// Output: [4 2 1 2 2 1 1 1 1 4 4]
}
