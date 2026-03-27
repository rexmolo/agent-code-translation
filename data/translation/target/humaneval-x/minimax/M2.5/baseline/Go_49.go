package main

import "fmt"

func Modp(n int, p int) int {
    ret := 1
    for i := 0; i < n; i++ {
        ret = (2 * ret) % p
    }
    return ret
}

func main() {
    // Test cases from the Python docstring
    fmt.Println(Modp(3, 5))   // Expected: 3
    fmt.Println(Modp(1101, 101)) // Expected: 2
    fmt.Println(Modp(0, 101))  // Expected: 1
    fmt.Println(Modp(3, 11))   // Expected: 8
    fmt.Println(Modp(100, 101)) // Expected: 1
}
