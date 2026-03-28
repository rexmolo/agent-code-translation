package main

import "fmt"

func RoundedAvg(n, m int) interface{} {
    if n > m {
        return -1
    }
    
    summation := 0
    for i := n; i <= m; i++ {
        summation += i
    }
    
    avg := float64(summation) / float64(m-n+1)
    
    // Python's round() uses "round half to even" banking rule
    // When exactly .5, round to nearest even integer
    rounded := int(avg + 0.5)
    fractional := avg - float64(int(avg))
    if fractional == 0.5 && rounded%2 == 0 {
        rounded--
    }
    
    return fmt.Sprintf("0b%d", rounded)
}

func main() {
    fmt.Println(RoundedAvg(1, 5))
    fmt.Println(RoundedAvg(7, 5))
    fmt.Println(RoundedAvg(10, 20))
    fmt.Println(RoundedAvg(20, 33))
}