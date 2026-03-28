package main

import "math"

// StartsOneEnds returns the count of n-digit positive integers that start or end with 1.
func StartsOneEnds(n int) int {
    if n == 1 {
        return 1
    }
    return 18 * int(math.Pow10(n-2))
}

func main() {
    // Example usage - can be tested with various inputs
    println(StartsOneEnds(1)) // Output: 1
    println(StartsOneEnds(2)) // Output: 18
    println(StartsOneEnds(3)) // Output: 180
}