package main

import "strings"

func ParseMusic(musicString string) []int {
	noteMap := map[string]int{
		"o":  4,
		"o|": 2,
		".|": 1,
	}

	var result []int
	for _, note := range strings.Split(musicString, " ") {
		if note != "" {
			result = append(result, noteMap[note])
		}
	}
	return result
}
