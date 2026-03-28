package main

import "fmt"

func IsMultiplyPrime(a int) bool {
    // Helper function to check if a number is prime
    isPrime := func(n int) bool {
        if n < 2 {
            return false
        }
        for j := 2; j < n; j++ {
            if n%j == 0 {
                return false
            }
        }
        return true
    }

    // Try all combinations of three prime numbers
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
    // Test case from the example: 30 = 2 * 3 * 5
    fmt.Println(IsMultiplyPrime(30)) // should print true
}
