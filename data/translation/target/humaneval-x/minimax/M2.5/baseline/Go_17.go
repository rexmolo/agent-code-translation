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

	parts := strings.Split(musicString, " ")

	var result []int
	for _, part := range parts {
		if part == "" {
			continue
		}
		if val, ok := noteMap[part]; ok {
			result = append(result, val)
		}
	}

	result
}