package main

import (
    "bufio"
    "fmt"
    "os"
)

func Intersection(interval1 [2]int, interval2 [2]int) string {
    // Helper function to check if a number is prime
    isPrime := func(num int) bool {
        if num == 0 || num == 1 {
            return false
        }
        if num == 2 {
            return true
        }
        for i := 2; i < num; i++ {
            if num%i == 0 {
                return false
            }
        }
        return true
    }

    l := max(interval1[0], interval2[0])
    r := min(interval1[1], interval2[1])
    length := r - l

    if length > 0 && isPrime(length) {
        return "YES"
    }
    return "NO"
}

func main() {
    // Read input
    reader := bufio.NewReader(os.Stdin)
    
    var n int
    fmt.Fscan(reader, &n)
    
    for i := 0; i < n; i++ {
        var a1, b1, a2, b2 int
        fmt.Fscan(reader, &a1, &b1, &a2, &b2)
        result := Intersection([2]int{a1, b1}, [2]int{a2, b2})
        fmt.Println(result)
    }
}
