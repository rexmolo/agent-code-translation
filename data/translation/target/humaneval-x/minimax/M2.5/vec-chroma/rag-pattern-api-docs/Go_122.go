package main

import (
    "fmt"
    "strconv"
)

func AddElements(arr []int, k int) int {
    sum := 0
    for i := 0; i < k; i++ {
        s := strconv.Itoa(arr[i])
        if len(s) <= 2 {
            sum += arr[i]
        }
    }
    return sum
}

func main() {
    // Example test case from the docstring
    arr := []int{111, 21, 3, 4000, 5, 6, 7, 8, 9}
    k := 4
    result := AddElements(arr, k)
    fmt.Println(result) // Expected output: 24
}
