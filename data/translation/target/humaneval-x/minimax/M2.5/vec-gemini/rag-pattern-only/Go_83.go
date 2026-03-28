package main

import "fmt"

func StartsOneEnds(n int) int {
    if n == 1 {
        return 1
    }
    
    // Calculate 10^(n-2)
    power := 1
    for i := 0; i < n-2; i++ {
        power *= 10
    }
    
    return 18 * power
}

func main() {
    // Test cases
    fmt.Println(StartsOneEnds(1)) // 1
    fmt.Println(StartsOneEnds(2)) // 18
    fmt.Println(StartsOneEnds(3)) // 180
}