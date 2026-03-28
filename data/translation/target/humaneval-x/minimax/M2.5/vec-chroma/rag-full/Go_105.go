package main

import (
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

    // Create a copy to avoid modifying the original slice
    sortedArr := make([]int, len(arr))
    copy(sortedArr, arr)

    // Sort in reverse order (descending)
    sort.Slice(sortedArr, func(i, j int) bool {
        return sortedArr[i] > sortedArr[j]
    })

    // Map valid digits to their names, ignoring others
    var result []string
    for _, v := range sortedArr {
        if name, ok := digitNames[v]; ok {
            result = append(result, name)
        }
    }

    return result
}