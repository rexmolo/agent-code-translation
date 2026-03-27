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
	notes := strings.Split(music_string, " ")

	for _, note := range notes {
		if note == "" {
			continue
		}
		if beats, ok := noteMap[note]; ok {
			result = append(result, beats)
		}
	}

	return result
}

func main() {
	result := ParseMusic("o o| .| o| o| .| .| .| .| o o")
	fmt.Println(result)
}