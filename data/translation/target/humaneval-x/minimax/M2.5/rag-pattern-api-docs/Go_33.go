package main

import "slices"

func SortThird(l []int) []int {
    // Create a copy of the original slice
    result := make([]int, len(l))
    copy(result, l)

    // Extract elements at indices divisible by 3 (0, 3, 6, 9, ...)
    var thirdIndices []int
    for i := 0; i < len(result); i += 3 {
        thirdIndices = append(thirdIndices, result[i])
    }

    // Sort the extracted elements
    slices.Sort(thirdIndices)

    // Put them back in their original positions
    idx := 0
    for i := 0; i < len(result); i += 3 {
        result[i] = thirdIndices[idx]
        idx++
    }

    return result
}
