package main

import (
    "fmt"
    "sort"
)

func SortEven(l []int) []int {
    // Extract elements at even indices (0, 2, 4, ...)
    var evens []int
    for i := 0; i < len(l); i += 2 {
        evens = append(evens, l[i])
    }

    // Extract elements at odd indices (1, 3, 5, ...)
    var odds []int
    for i := 1; i < len(l); i += 2 {
        odds = append(odds, l[i])
    }

    // Sort the even-indexed elements
    sort.Ints(evens)

    // Interleave sorted evens with original odds
    var ans []int
    for i := 0; i < len(evens); i++ {
        ans = append(ans, evens[i])
        if i < len(odds) {
            ans = append(ans, odds[i])
        }
    }

    return ans
}

func main() {
    // Example usage
    fmt.Println(SortEven([]int{1, 2, 3}))         // [1 2 3]
    fmt.Println(SortEven([]int{5, 6, 3, 4}))      // [3 6 5 4]
}