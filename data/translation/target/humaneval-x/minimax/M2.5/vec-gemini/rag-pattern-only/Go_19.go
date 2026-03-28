package main

import (
    "fmt"
    "slices"
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

    // Split string by spaces and filter empty entries (equivalent to Python's [x for x in numbers.split(' ') if x])
    parts := strings.Fields(numbers)

    // Sort using anonymous function as key (equivalent to Python's key=lambda x: value_map[x])
    slices.SortFunc(parts, func(a, b string) int {
        return valueMap[a] - valueMap[b]
    })

    // Join back into space-delimited string (equivalent to ' '.join(...))
    return strings.Join(parts, " ")
}

func main() {
    fmt.Println(SortNumbers("three one five"))
}
