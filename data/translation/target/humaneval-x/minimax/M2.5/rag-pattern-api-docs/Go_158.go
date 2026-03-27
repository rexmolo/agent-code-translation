package main

import (
    "fmt"
    "sort"
)

func FindMax(words []string) string {
    if len(words) == 0 {
        return ""
    }
    
    // Sort by: 1) descending number of unique characters, 2) ascending lexicographical order
    sort.Slice(words, func(i, j int) bool {
        countI := uniqueCharCount(words[i])
        countJ := uniqueCharCount(words[j])
        
        // If unique counts differ, sort by count descending (more unique chars first)
        if countI != countJ {
            return countI > countJ
        }
        
        // If counts are equal, sort lexicographically ascending (first in alphabetical order)
        return words[i] < words[j]
    })
    
    return words[0]
}

func uniqueCharCount(s string) int {
    unique := make(map[rune]struct{})
    for _, r := range s {
        unique[r] = struct{}{}
    }
    return len(unique)
}

func main() {
    // Test cases
    fmt.Println(FindMax([]string{"name", "of", "string"})) // "string"
    fmt.Println(FindMax([]string{"name", "enam", "game"})) // "enam"
    fmt.Println(FindMax([]string{"aaaaaaa", "bb", "cc"})) // "aaaaaaa"
}