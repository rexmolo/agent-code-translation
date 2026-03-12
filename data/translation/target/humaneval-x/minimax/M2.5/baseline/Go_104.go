package main

import (
    "fmt"
    "sort"
)

func UniqueDigits(x []int) []int {
    var result []int
    
    for _, num := range x {
        if allDigitsOdd(num) {
            result = append(result, num)
        }
    }
    
    sort.Ints(result)
    return result
}

func allDigitsOdd(n int) bool {
    if n == 0 {
        return false // 0 is an even digit
    }
    
    for n > 0 {
        digit := n % 10
        if digit%2 == 0 {
            return false
        }
        n /= 10
    }
    return true
}

func main() {
    // Test cases
    fmt.Println(UniqueDigits([]int{15, 33, 1422, 1}))
    fmt.Println(UniqueDigits([]int{152, 323, 1422, 10}))
}