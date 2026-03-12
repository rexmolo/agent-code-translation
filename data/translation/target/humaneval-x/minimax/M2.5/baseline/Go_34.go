package main

import (
    "fmt"
    "sort"
)

func Unique(l []int) []int {
    // Use a map to track seen elements (simulating a set)
    seen := make(map[int]struct{})
    result := []int{}

    // Iterate through the slice and add only unique elements
    for _, v := range l {
        if _, exists := seen[v]; !exists {
            seen[v] = struct{}{}
            result = append(result, v)
        }
    }

    // Sort the result in ascending order
    sort.Ints(result)

    return result
}

func main() {
    // Example usage
    l := []int{5, 3, 5, 2, 3, 3, 9, 0, 123}
    fmt.Println(Unique(l))
}
