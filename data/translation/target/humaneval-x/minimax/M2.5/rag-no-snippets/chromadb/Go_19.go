package main

import (
    "sort"
    "strings"
)

func SortNumbers(numbers string) string {
    valueMap := map[string]int{
        "zero":  0,
        "one":   1,
        "two":   2,
        "three": 3,
        "four":  4,
        "five":  5,
        "six":   6,
        "seven": 7,
        "eight": 8,
        "nine":  9,
    }

    // Split by space and filter out empty strings
    parts := strings.Split(numbers, " ")
    var words []string
    for _, w := range parts {
        if w != "" {
            words = append(words, w)
        }
    }

    // Sort by numeric value using the map
    sort.Slice(words, func(i, j int) bool {
        return valueMap[words[i]] < valueMap[words[j]]
    })

    return strings.Join(words, " ")
}
