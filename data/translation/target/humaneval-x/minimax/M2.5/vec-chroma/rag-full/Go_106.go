package main

func F(n int) []int {
    ret := make([]int, 0)
    for i := 1; i <= n; i++ {
        if i%2 == 0 {
            // factorial of i
            x := 1
            for j := 1; j <= i; j++ {
                x *= j
            }
            ret = append(ret, x)
        } else {
            // sum of numbers from 1 to i
            x := 0
            for j := 1; j <= i; j++ {
                x += j
            }
            ret = append(ret, x)
        }
    }
    return ret
}