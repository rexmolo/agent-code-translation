package main

import "strings"

func ParseMusic(music_string string) []int {
    noteMap := map[string]int{
        "o":  4,
        "o|": 2,
        ".|": 1,
    }

    parts := strings.Fields(music_string)
    result := make([]int, 0, len(parts))

    for _, part := range parts {
        result = append(result, noteMap[part])
    }

    return result
}
