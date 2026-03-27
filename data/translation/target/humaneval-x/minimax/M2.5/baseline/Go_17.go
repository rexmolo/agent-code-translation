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
    for _, token := range strings.Split(music_string, " ") {
        if token == "" {
            continue
        }
        result = append(result, noteMap[token])
    }

    return result
}

func main() {
    // Test
    result := ParseMusic("o o| .| o| o| .| .| .| .| o o")
    fmt.Println(result)
}