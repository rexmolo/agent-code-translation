package main

import (
	"strings"
)

func ParseMusic(musicString string) []int {
	// Map of note symbols to their beat durations
	noteMap := map[string]int{
		"o":  4,
		"o|": 2,
		".|": 1,
	}

	// Split the input string by spaces
	parts := strings.Split(musicString, " ")

	// Initialize result slice with capacity hint for efficiency
	result := make([]int, 0, len(parts))

	// Iterate through parts and map non-empty notes to their beat values
	for _, note := range parts {
		if note != "" {
			result = append(result, noteMap[note])
		}
	}

	return result
}