package main

import "sort"

func SortEven(l []int) []int {
    // Extract elements at even indices (0, 2, 4, ...)
    evens := make([]int, 0, (len(l)+1)/2)
    for i := 0; i < len(l); i += 2 {
        evens = append(evens, l[i])
    }

    // Extract elements at odd indices (1, 3, 5, ...)
    odds := make([]int, 0, len(l)/2)
    for i := 1; i < len(l); i += 2 {
        odds = append(odds, l[i])
    }

    // Sort the even-indexed elements
    sort.Ints(evens)

    // Interleave sorted evens with odds
    result := make([]int, 0, len(l))
    for i := 0; i < len(odds); i++ {
        result = append(result, evens[i], odds[i])
    }

    // If there's one more even than odd, append the last even element
    if len(evens) > len(odds) {
        result = append(result, evens[len(evens)-1])
    }

    return result
}
