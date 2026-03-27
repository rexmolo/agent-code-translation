package main

import "fmt"

func Derivative(xs []int) []int {
    if len(xs) <= 1 {
        return []int{}
    }
    result := make([]int, len(xs)-1)
    for i := 1; i < len(xs); i++ {
        result[i-1] = i * xs[i]
    }
    return result
}

func main() {
    fmt.Println(Derivative([]int{3, 1, 2, 4, 5})) // Output: [1 4 12 20]
    fmt.Println(Derivative([]int{1, 2, 3}))      // Output: [2 6]
}
