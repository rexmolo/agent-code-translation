package main

import "fmt"

func MakeAPile(n int) []int {
    result := make([]int, n)
    for i := 0; i < n; i++ {
        result[i] = n + 2*i
    }
    return result
}

func main() {
    // Example usage
    fmt.Println(MakeAPile(3)) // Output: [3 5 7]
    fmt.Println(MakeAPile(4)) // Output: [4 6 8 10]
    fmt.Println(MakeAPile(5)) // Output: [5 7 9 11 13]
}
