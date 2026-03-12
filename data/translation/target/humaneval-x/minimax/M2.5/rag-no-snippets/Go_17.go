package main

import (
	"fmt"
	"strings"
)

func ParseMusic(musicString string) []int {
	noteMap := map[string]int{
		"o":  4,
		"o|": 2,
		".|": 1,
	}

	parts := strings.Split(musicString, " ")
	var result []int

	for _, x := range parts {
		if x != "" {
			result = append(result, noteMap[x])
		}
	}

	return result
}

func main() {
	// Example usage
	input := "o o| .| o| o| .| .| .| .| o o"
	result := ParseMusic(input)
	fmt.Println(result)
}
