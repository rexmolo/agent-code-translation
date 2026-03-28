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

    parts := strings.Split(numbers, " ")

    // Filter out empty strings (equivalent to Python's [x for x in numbers.split(' ') if x])
    var words []string
    for _, x := range parts {
        if x != "" {
            words = append(words, x)
        }
    }

    // Sort using custom comparator based on value map
    sort.Slice(words, func(i, j int) int {
        return valueMap[words[i]] - valueMap[words[j]]
    })

    return strings.Join(words, " ")
}
