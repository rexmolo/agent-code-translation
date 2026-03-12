package main

import (
    "fmt"
    "math"
)

func StartsOneEnds(n int) int {
    if n == 1 {
        return 1
    }
    return 18 * int(math.Pow(10, float64(n-2)))
}

func main() {
    fmt.Println(StartsOneEnds(1)) // 1
    fmt.Println(StartsOneEnds(2)) // 18
    fmt.Println(StartsOneEnds(3)) // 180
}