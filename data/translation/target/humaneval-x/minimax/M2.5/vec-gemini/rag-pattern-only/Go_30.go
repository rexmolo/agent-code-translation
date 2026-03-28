package main

import "fmt"

func GetPositive(l []int) []int {
    result := make([]int, 0, len(l))
    for _, v := range l {
        if v > 0 {
            result = append(result, v)
        }
    }
    return result
}

func main() {
    // Test cases
    fmt.Println(GetPositive([]int{-1, 2, -4, 5, 6}))
    fmt.Println(GetPositive([]int{5, 3, -5, 2, -3, 3, 9, 0, 123, 1, -10}))
}
