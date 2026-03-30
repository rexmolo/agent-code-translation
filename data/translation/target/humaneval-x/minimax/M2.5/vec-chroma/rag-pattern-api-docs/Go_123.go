package main

import (
    "fmt"
    "sort"
)

func GetOddCollatz(n int) []int {
    var oddCollatz []int

    // If n is odd, start with n in the list
    if n%2 != 0 {
        oddCollatz = append(oddCollatz, n)
    }

    // Generate Collatz sequence and collect odd numbers
    for n > 1 {
        if n%2 == 0 {
            n = n / 2
        } else {
            n = n*3 + 1
        }

        // If new n is odd, append it
        if n%2 != 0 {
            oddCollatz = append(oddCollatz, n)
        }
    }

    // Sort in increasing order
    sort.Slice(oddCollatz, func(i, j int) bool {
        return oddCollatz[i] < oddCollatz[j]
    })

    return oddCollatz
}

func main() {
    // Test example
    result := GetOddCollatz(5)
    fmt.Println(result)
}