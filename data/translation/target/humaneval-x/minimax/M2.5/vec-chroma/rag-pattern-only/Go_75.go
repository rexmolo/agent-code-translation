package main

import "fmt"

func isPrime(n int) bool {
    for j := 2; j < n; j++ {
        if n%j == 0 {
            return false
        }
    }
    return true
}

func IsMultiplyPrime(a int) bool {
    for i := 2; i <= 100; i++ {
        if !isPrime(i) {
            continue
        }
        for j := 2; j <= 100; j++ {
            if !isPrime(j) {
                continue
            }
            for k := 2; k <= 100; k++ {
                if !isPrime(k) {
                    continue
                }
                if i*j*k == a {
                    return true
                }
            }
        }
    }
    return false
}

func main() {
    fmt.Println(IsMultiplyPrime(30)) // Output: true
}