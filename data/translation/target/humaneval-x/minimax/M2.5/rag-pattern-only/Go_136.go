package main

import (
    "fmt"
)

func LargestSmallestIntegers(lst []int) [2]interface{} {
    // Use sentinel values to track the largest negative and smallest positive
    // Initialize with values that will be replaced when found
    maxNegative := 1  // Start with value that will be replaced by any negative
    minPositive := -1 // Start with value that will be replaced by any positive
    
    hasNegative := false
    hasPositive := false
    
    for _, n := range lst {
        if n < 0 {
            hasNegative = true
            // Among negatives, the "largest" is the one closest to zero (highest value)
            if n > maxNegative {
                maxNegative = n
            }
        }
        if n > 0 {
            hasPositive = true
            // The smallest positive is the minimum positive value
            if n < minPositive {
                minPositive = n
            }
        }
    }
    
    var result [2]interface{}
    
    // Assign results - nil if not found
    if hasNegative {
        result[0] = maxNegative
    } else {
        result[0] = nil
    }
    
    if hasPositive {
        result[1] = minPositive
    } else {
        result[1] = nil
    }
    
    return result
}

func main() {
    // Test cases
    fmt.Println(LargestSmallestIntegers([]int{2, 4, 1, 3, 5, 7})) // [nil 1]
    fmt.Println(LargestSmallestIntegers([]int{}))                // [nil nil]
    fmt.Println(LargestSmallestIntegers([]int{0}))               // [nil nil]
}