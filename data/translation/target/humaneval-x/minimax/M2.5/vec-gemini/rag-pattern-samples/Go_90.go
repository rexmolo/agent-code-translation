package main

import (
    "sort"
)

func NextSmallest(lst []int) interface{} {
    // Remove duplicates using a map
    unique := make(map[int]bool)
    for _, v := range lst {
        unique[v] = true
    }
    
    // Convert map keys to slice
    keys := make([]int, 0, len(unique))
    for k := range unique {
        keys = append(keys, k)
    }
    
    // Sort the slice
    sort.Ints(keys)
    
    // Return the 2nd smallest element or nil if there aren't enough unique elements
    if len(keys) < 2 {
        return nil
    }
    return keys[1]
}
