package main

import (
    "fmt"
    "sort"
)

func GetOddCollatz(n int) []int {
    var oddCollatz []int

    // Initialize with n if it's odd
    if n%2 == 1 {
        oddCollatz = append(oddCollatz, n)
    }

    // Generate Collatz sequence until reaching 1
    for n > 1 {
        if n%2 == 0 {
            n = n / 2
        } else {
            n = n*3 + 1
        }

        // Collect odd numbers in the sequence
        if n%2 == 1 {
            oddCollatz = append(oddCollatz, n)
        }
    }

    // Sort in increasing order
    sort.Ints(oddCollatz)
    return oddCollatz
}

func main() {
    fmt.Println(GetOddCollatz(5)) // Output: [1 5]
}