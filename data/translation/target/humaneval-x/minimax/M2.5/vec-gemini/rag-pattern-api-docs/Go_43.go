package main

import "fmt"

func PairsSumToZero(l []int) bool {
    for i, l1 := range l {
        for j := i + 1; j < len(l); j++ {
            if l1+l[j] == 0 {
                return true
            }
        }
    }
    return false
}

func main() {
    // Test cases from docstring
    fmt.Println(PairsSumToZero([]int{1, 3, 5, 0}))    // Expected: false
    fmt.Println(PairsSumToZero([]int{1, 3, -2, 1}))  // Expected: false
    fmt.Println(PairsSumToZero([]int{1, 2, 3, 7}))   // Expected: false
    fmt.Println(PairsSumToZero([]int{2, 4, -5, 3, 5, 7})) // Expected: true
    fmt.Println(PairsSumToZero([]int{1}))            // Expected: false
}