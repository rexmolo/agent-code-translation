package main

import (
    "math"
)

func PrimeFib(n int) int {
    // Helper function to check if a number is prime
    isPrime := func(p int) bool {
        if p < 2 {
            return false
        }
        // Check divisibility from 2 up to min(sqrt(p)+1, p-1)
        limit := int(math.Sqrt(float64(p))) + 1
        upperBound := limit
        if p-1 < limit {
            upperBound = p - 1
        }
        for k := 2; k < upperBound; k++ {
            if p%k == 0 {
                return false
            }
        }
        return true
    }

    // Initialize Fibonacci sequence with first two values
    f := []int{0, 1}

    for {
        // Generate next Fibonacci number
        next := f[len(f)-1] + f[len(f)-2]
        f = append(f, next)

        // Check if it's prime
        if isPrime(next) {
            n--
        }

        // Return when we've found n prime Fibonacci numbers
        if n == 0 {
            return next
        }
    }
}
