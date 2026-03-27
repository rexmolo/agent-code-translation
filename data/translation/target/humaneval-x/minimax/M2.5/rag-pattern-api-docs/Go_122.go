package main

import (
    "strconv"
)

func AddElements(arr []int, k int) int {
    sum := 0
    for i := 0; i < k; i++ {
        // Check if the element has at most 2 digits
        if len(strconv.Itoa(arr[i])) <= 2 {
            sum += arr[i]
        }
    }
    return sum
}

func main() {
    // Example usage
    arr := []int{111, 21, 3, 4000, 5, 6, 7, 8, 9}
    k := 4
    result := AddElements(arr, k)
    println(result) // Output: 24
}