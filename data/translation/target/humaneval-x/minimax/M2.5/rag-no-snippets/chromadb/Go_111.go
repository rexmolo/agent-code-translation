package main

import (
    "strings"
)

func Histogram(test string) map[rune]int {
    // Handle empty string case
    if test == "" {
        return map[rune]int{}
    }

    // Split the string by space
    list1 := strings.Split(test, " ")

    // Count occurrences of each element
    counts := make(map[string]int)
    for _, i := range list1 {
        if i == "" {
            continue
        }
        counts[i]++
    }

    // Find the maximum count
    t := 0
    for _, count := range counts {
        if count > t {
            t = count
        }
    }

    // If no valid letters (all empty), return empty map
    if t == 0 {
        return map[rune]int{}
    }

    // Build result map with letters that have the maximum count
    result := make(map[rune]int)
    for letter, count := range counts {
        if count == t {
            // Convert single character string to rune
            for _, r := range letter {
                result[r] = count
                break // Only take the first rune since it's a single character
            }
        }
    }

    return result
}