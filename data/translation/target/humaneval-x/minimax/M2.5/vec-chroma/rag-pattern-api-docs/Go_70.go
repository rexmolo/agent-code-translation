package main

import (
    "fmt"
    "slices"
)

func StrangeSortList(lst []int) []int {
    res := make([]int, 0, len(lst))
    switchOn := true

    for len(lst) > 0 {
        var val int
        if switchOn {
            val = slices.Min(lst)
        } else {
            val = slices.Max(lst)
        }
        res = append(res, val)

        // Remove the first occurrence of val from lst
        idx := slices.Index(lst, val)
        lst = slices.Delete(lst, idx, idx+1)

        switchOn = !switchOn
    }
    return res
}

func main() {
    // Test cases
    fmt.Println(StrangeSortList([]int{1, 2, 3, 4}))      // [1 4 2 3]
    fmt.Println(StrangeSortList([]int{5, 5, 5, 5}))      // [5 5 5 5]
    fmt.Println(StrangeSortList([]int{}))                // []
}