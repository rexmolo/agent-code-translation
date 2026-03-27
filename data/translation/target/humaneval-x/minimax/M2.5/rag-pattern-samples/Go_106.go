package main

import "fmt"

func F(n int) []int {
    ret := make([]int, 0, n)
    for i := 1; i <= n; i++ {
        var x int
        if i%2 == 0 {
            // even: factorial of i
            x = 1
            for j := 1; j <= i; j++ {
                x *= j
            }
        } else {
            // odd: sum of 1 to i
            x = 0
            for j := 1; j <= i; j++ {
                x += j
            }
        }
        ret = append(ret, x)
    }
    return ret
}

func main() {
    // Example usage
    result := F(5)
    fmt.Println(result)
}
