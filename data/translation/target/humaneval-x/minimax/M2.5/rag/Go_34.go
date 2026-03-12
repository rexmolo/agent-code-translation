package main

import "sort"

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
