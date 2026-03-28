package main

import (
    "strconv"
    "strings"
)

func FruitDistribution(s string, n int) int {
    parts := strings.Split(s, " ")
    var nums []int
    
    for _, part := range parts {
        // Try to convert each part to an integer
        // strconv.Atoi only succeeds for valid integer strings,
        // which matches Python's isdigit() behavior
        if num, err := strconv.Atoi(part); err == nil {
            nums = append(nums, num)
        }
    }
    
    sum := 0
    for _, num := range nums {
        sum += num
    }
    
    return n - sum
}