package main

import "slices"

func Pluck(arr []int) []int {
    if len(arr) == 0 {
        return []int{}
    }

    minVal := -1
    minIdx := -1

    for i, v := range arr {
        if v%2 == 0 {
            if minVal == -1 || v < minVal {
                minVal = v
                minIdx = i
            }
        }
    }

    if minVal == -1 {
        return []int{}
    }

    return []int{minVal, minIdx}
}

// Alternative implementation using slices package functions:
/*
func Pluck(arr []int) []int {
    if len(arr) == 0 {
        return []int{}
    }

    var evens []int
    for _, v := range arr {
        if v%2 == 0 {
            evens = append(evens, v)
        }
    }

    if len(evens) == 0 {
        return []int{}
    }

    minVal := slices.Min(evens)
    minIdx := slices.Index(arr, minVal)

    return []int{minVal, minIdx}
}
*/
