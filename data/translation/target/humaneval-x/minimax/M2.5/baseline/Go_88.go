package main

import (
    "sort"
)

func SortArray(array []int) []int {
    if len(array) == 0 {
        return []int{}
    }

    // Create a copy to avoid modifying the original array
    result := make([]int, len(array))
    copy(result, array)

    // Check if sum of first and last element is even
    isEven := (array[0]+array[len(array)-1])%2 == 0

    if isEven {
        // Sort in descending order
        sort.Ints(result)
        // Reverse the array
        for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
            result[i], result[j] = result[j], result[i]
        }
    } else {
        // Sort in ascending order
        sort.Ints(result)
    }

    return result
}
