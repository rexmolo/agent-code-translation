package main

import (
    "fmt"
    "slices"
)

func StrangeSortList(lst []int) []int {
    // Create a copy to avoid modifying the original slice
    working := make([]int, len(lst))
    copy(working, lst)

    res := make([]int, 0, len(lst))
    switchFlag := true // true = take min, false = take max

    for len(working) > 0 {
        var val int
        if switchFlag {
            val = slices.Min(working)
        } else {
            val = slices.Max(working)
        }
        res = append(res, val)

        // Remove the first occurrence of val
        idx := slices.Index(working, val)
        working = slices.Delete(working, idx, idx+1)

        switchFlag = !switchFlag
    }

    return res
}

func main() {
    // Test cases
    fmt.Println(StrangeSortList([]int{1, 2, 3, 4})) // [1 4 2 3]
    fmt.Println(StrangeSortList([]int{5, 5, 5, 5})) // [5 5 5 5]
    fmt.Println(StrangeSortList([]int{}))           // []
}
