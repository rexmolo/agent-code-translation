package main

import "fmt"

func IsSimplePower(x int, n int) bool {
    if n == 1 {
        return x == 1
    }
    power := 1
    for power < x {
        power = power * n
    }
    return power == x
}

func main() {
    // Test examples from the docstring
    fmt.Println(IsSimplePower(1, 4)) // true
    fmt.Println(IsSimplePower(2, 2)) // true
    fmt.Println(IsSimplePower(8, 2)) // true
    fmt.Println(IsSimplePower(3, 2)) // false
    fmt.Println(IsSimplePower(3, 1)) // false
    fmt.Println(IsSimplePower(5, 3)) // false
}