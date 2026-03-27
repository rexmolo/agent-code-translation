package main

import (
    "fmt"
    "sort"
)

func ByLength(arr []int) []string {
    digitNames := map[int]string{
        1: "One",
        2: "Two",
        3: "Three",
        4: "Four",
        5: "Five",
        6: "Six",
        7: "Seven",
        8: "Eight",
        9: "Nine",
    }

    // Sort in reverse order (largest to smallest)
    sort.Slice(arr, func(i, j int) bool {
        return arr[i] > arr[j]
    })

    var result []string
    for _, v := range arr {
        // Use comma-ok idiom to check if key exists in map
        if name, ok := digitNames[v]; ok {
            result = append(result, name)
        }
    }

    return result
}

func main() {
    // Test cases
    fmt.Println(ByLength([]int{2, 1, 1, 4, 5, 8, 2, 3}))
    fmt.Println(ByLength([]int{}))
    fmt.Println(ByLength([]int{1, -1, 55}))
}