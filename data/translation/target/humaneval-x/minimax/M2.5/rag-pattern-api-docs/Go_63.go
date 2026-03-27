package main

import "fmt"

func Fibfib(n int) int {
    if n == 0 {
        return 0
    }
    if n == 1 {
        return 0
    }
    if n == 2 {
        return 1
    }

    // Use iterative approach instead of recursive to avoid exponential time complexity
    // a = fibfib(n-3), b = fibfib(n-2), c = fibfib(n-1)
    a, b, c := 0, 0, 1
    for i := 3; i <= n; i++ {
        a, b, c = b, c, a+b+c
    }
    return c
}

func main() {
    fmt.Println(Fibfib(1)) // 0
    fmt.Println(Fibfib(5)) // 4
    fmt.Println(Fibfib(8)) // 24
}