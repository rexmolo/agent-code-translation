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
	result := make([]int, 0)

	for _, part := range parts {
		if part != "" {
			if beats, ok := noteMap[part]; ok {
				result = append(result, beats)
			}
		}
	}

	return result
}

func main() {
	// Test the function
	result := ParseMusic("o o| .| o| o| .| .| .| .| o o")
	fmt.Println(result)
}