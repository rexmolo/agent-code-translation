package main

import (
    "fmt"
    "sort"
)

func StrangeSortList(lst []int) []int {
    if len(lst) == 0 {
        return []int{}
    }
    
    // Sort the slice first for efficient min/max access
    sorted := make([]int, len(lst))
    copy(sorted, lst)
    sort.Ints(sorted)
    
    res := make([]int, 0, len(lst))
    left, right := 0, len(sorted)-1
    
    // Alternate between smallest (left) and largest (right)
    for left <= right {
        res = append(res, sorted[left])
        left++
        if left <= right {
            res = append(res, sorted[right])
            right--
        }
    }
    
    return res
}

func main() {
    // Test cases
    fmt.Println(StrangeSortList([]int{}))           // []
    fmt.Println(StrangeSortList([]int{1, 2, 3, 4})) // [1 4 2 3]
    fmt.Println(StrangeSortList([]int{5, 5, 5, 5})) // [5 5 5 5]
}