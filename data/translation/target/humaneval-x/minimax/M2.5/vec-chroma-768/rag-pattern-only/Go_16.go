package main

import (
    "fmt"
    "strings"
)

func CountDistinctCharacters(str string) int {
    seen := make(map[rune]struct{})
    for _, r := range strings.ToLower(str) {
        seen[r] = struct{}{}
    }
    return len(seen)
}

func main() {
    fmt.Println(CountDistinctCharacters("xyzXYZ"))
    fmt.Println(CountDistinctCharacters("Jerry"))
}