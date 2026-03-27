package main

import "fmt"

func Fib4(n int) int {
    // Base cases
    if n == 0 {
        return 0
    }
    if n == 1 {
        return 0
    }
    if n == 2 {
        return 2
    }
    if n == 3 {
        return 0
    }

    // Use sliding window approach with 4 variables
    // a, b, c, d represent fib4(n-4), fib4(n-3), fib4(n-2), fib4(n-1)
    a, b, c, d := 0, 0, 2, 0

    // Iteratively compute fib4(n)
    for i := 4; i <= n; i++ {
        next := a + b + c + d
        a = b
        b = c
        c = d
        d = next
    }

    return d
}

func main() {
    // Test cases
    fmt.Println(Fib4(5)) // 4
    fmt.Println(Fib4(6)) // 8
    fmt.Println(Fib4(7)) // 14
}