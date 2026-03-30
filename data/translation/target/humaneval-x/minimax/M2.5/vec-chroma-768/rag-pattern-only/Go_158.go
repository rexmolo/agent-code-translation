package main

import (
    "cmp"
    "slices"
)

func FindMax(words []string) string {
    if len(words) == 0 {
        return ""
    }

    // Helper function to count unique characters in a string
    uniqueCount := func(s string) int {
        seen := make(map[rune]bool)
        for _, c := range s {
            seen[c] = true
        }
        return len(seen)
    }

    // Sort by: descending unique characters, ascending lexicographical order
    slices.SortFunc(words, func(a, b string) int {
        countA := uniqueCount(a)
        countB := uniqueCount(b)

        // If different unique counts, higher count comes first (descending)
        if countA != countB {
            return cmp.Compare(countB, countA)
        }
        // Same unique count, return lexicographically smaller (ascending)
        return cmp.Compare(a, b)
    })

    return words[0]
}

func main() {
    // Example usage
    words1 := []string{"name", "of", "string"}
    words2 := []string{"name", "enam", "game"}
    words3 := []string{"aaaaaaa", "bb", "cc"}

    println(FindMax(words1)) // "string"
    println(FindMax(words2)) // "enam"
    println(FindMax(words3)) // "aaaaaaa"
}