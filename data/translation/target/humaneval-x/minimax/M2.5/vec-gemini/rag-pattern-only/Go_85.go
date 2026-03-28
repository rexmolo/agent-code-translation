package main

import "fmt"

func Add(lst []int) int {
    sum := 0
    for i := 1; i < len(lst); i += 2 {
        if lst[i]%2 == 0 {
            sum += lst[i]
        }
    }
    return sum
}

func main() {
    // Test examples
    fmt.Println(Add([]int{4, 2, 6, 7})) // Expected: 2
}
