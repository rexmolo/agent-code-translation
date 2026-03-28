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
    // Test with the example from docstring
    result := ParseMusic("o o| .| o| o| .| .| .| .| o o")
    fmt.Println(result)
    // Output: [4 2 1 2 2 1 1 1 1 4 4]
}