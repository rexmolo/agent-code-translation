package main

import (
    "fmt"
    "math"
)

func RoundedAvg(n, m int) interface{} {
    if n > m {
        return -1
    }
    
    summation := 0
    for i := n; i <= m; i++ {
        summation += i
    }
    
    count := m - n + 1
    avg := float64(summation) / float64(count)
    rounded := int(math.Round(avg))
    
    return "0b" + fmt.Sprintf("%b", rounded)
}
