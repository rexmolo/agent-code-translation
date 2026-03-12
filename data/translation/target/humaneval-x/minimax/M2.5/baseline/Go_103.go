package main

import (
    "fmt"
    "math"
)

func RoundedAvg(n, m int) interface{} {
    if m < n {
        return -1
    }

    sum := 0
    for i := n; i <= m; i++ {
        sum += i
    }

    avg := float64(sum) / float64(m-n+1)
    rounded := int(math.Round(avg))

    return "0b" + fmt.Sprintf("%b", rounded)
}

func main() {
    fmt.Println(RoundedAvg(1, 5))
    fmt.Println(RoundedAvg(7, 5))
    fmt.Println(RoundedAvg(10, 20))
    fmt.Println(RoundedAvg(20, 33))
}