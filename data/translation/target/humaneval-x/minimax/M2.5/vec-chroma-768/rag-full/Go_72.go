package main

import "fmt"

func WillItFly(q []int, w int) bool {
    // Check if sum of elements exceeds maximum weight
    sum := 0
    for _, v := range q {
        sum += v
    }
    if sum > w {
        return false
    }

    // Check if the list is balanced (palindromic)
    i, j := 0, len(q)-1
    for i < j {
        if q[i] != q[j] {
            return false
        }
        i++
        j--
    }

    return true
}

func main() {
    // Test examples
    fmt.Println(WillItFly([]int{1, 2}, 5))   // false
    fmt.Println(WillItFly([]int{3, 2, 3}, 1)) // false
    fmt.Println(WillItFly([]int{3, 2, 3}, 9)) // true
    fmt.Println(WillItFly([]int{3}, 5))       // true
}