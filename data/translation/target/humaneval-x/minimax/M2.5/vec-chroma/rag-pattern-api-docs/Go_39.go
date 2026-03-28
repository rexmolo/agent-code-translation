package main

import (
    "math"
)

func PrimeFib(n int) int {
    isPrime := func(p int) bool {
        if p < 2 {
            return false
        }
        // Python's range(2, min(int(math.sqrt(p)) + 1, p - 1)) means:
        // k goes from 2 to min(sqrt(p)+1, p-1) inclusive
        // So we need k <= min(sqrt(p)+1, p-1)
        // Since p-1 > sqrt(p)+1 for p > 2, we can just use sqrt(p)+1
        sqrtP := int(math.Sqrt(float64(p)))
        for k := 2; k <= sqrtP; k++ {
            if p%k == 0 {
                return false
            }
        }
        return true
    }

    // Initialize Fibonacci sequence
    f := []int{0, 1}

    for {
        // Append next Fibonacci number
        next := f[len(f)-1] + f[len(f)-2]
        f = append(f, next)

        if isPrime(next) {
            n--
        }
        if n == 0 {
            return next
        }
    }
}
