package main

import (
    "fmt"
    "sort"
)

func Unique(l []int) []int {
    seen := make(map[int]bool)
    result := make([]int, 0)

    for _, v := range l {
        if !seen[v] {
            seen[v] = true
            result = append(result, v)
        }
    }

    sort.Ints(result)
    return result
}

func main() {
    // Example usage
    fmt.Println(Unique([]int{5, 3, 5, 2, 3, 3, 9, 0, 123}))
}
