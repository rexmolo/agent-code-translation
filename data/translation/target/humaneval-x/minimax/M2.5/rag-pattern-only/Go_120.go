package main

import (
    "sort"
)

func Maximum(arr []int, k int) []int {
    if k == 0 {
        return []int{}
    }

    // Make a copy to avoid modifying the original slice
    copyArr := make([]int, len(arr))
    copy(copyArr, arr)

    // Sort the copy in ascending order
    sort.Ints(copyArr)

    // Get the k largest elements (they will be at the end after sorting)
    ans := copyArr[len(copyArr)-k:]

    // Return a copy to avoid returning a slice that shares the underlying array
    result := make([]int, len(ans))
    copy(result, ans)

    return result
}
