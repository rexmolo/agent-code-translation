package main

import (
    "fmt"
    "slices"
)

func LargestSmallestIntegers(lst []int) [2]interface{} {
    var negatives []int
    var positives []int

    // Filter negative and positive integers
    for _, x := range lst {
        if x < 0 {
            negatives = append(negatives, x)
        } else if x > 0 {
            positives = append(positives, x)
        }
    }

    var largestNegative interface{} = nil
    var smallestPositive interface{} = nil

    // Get largest of negative (max, since -1 > -5)
    if len(negatives) > 0 {
        largestNegative = slices.Max(negatives)
    }

    // Get smallest of positive (min, since 1 < 5)
    if len(positives) > 0 {
        smallestPositive = slices.Min(positives)
    }

    return [2]interface{}{largestNegative, smallestPositive}
}

func main() {
    // Test cases
    fmt.Println(LargestSmallestIntegers([]int{2, 4, 1, 3, 5, 7})) // [None 1]
    fmt.Println(LargestSmallestIntegers([]int{}))                // [None None]
    fmt.Println(LargestSmallestIntegers([]int{0}))               // [None None]
}