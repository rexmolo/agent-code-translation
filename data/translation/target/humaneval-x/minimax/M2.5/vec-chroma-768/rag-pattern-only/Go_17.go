package main

import (
	"strings"
)

func ParseMusic(musicString string) []int {
	noteMap := map[string]int{
		"o":  4,
		"o|": 2,
		".|": 1,
	}

	notes := strings.Split(musicString, " ")

	var result []int
	for _, note := range notes {
		if note == "" {
			continue
		}
		result = append(result, noteMap[note])
	}

	return result
}