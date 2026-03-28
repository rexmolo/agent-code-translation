package main

import (
    "fmt"
    "strings"
)

func CountDistinctCharacters(str string) int {
    // Convert string to lowercase
    lower := strings.ToLower(str)
    
    // Use map to simulate a set for distinct characters
    distinct := make(map[rune]struct{})
    for _, r := range lower {
        distinct[r] = struct{}{}
    }
    
    return len(distinct)
}

func main() {
    // Test cases
    fmt.Println(CountDistinctCharacters("xyzXYZ")) // Output: 3
    fmt.Println(CountDistinctCharacters("Jerry")) // Output: 4
}