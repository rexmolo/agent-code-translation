package main

import "sort"

func GetOddCollatz(n int) []int {
    var oddCollatz []int

    // Handle initial number - if n is odd, include it
    if n%2 != 0 {
        oddCollatz = append(oddCollatz, n)
    }

    // Generate Collatz sequence until n reaches 1
    for n > 1 {
        if n%2 == 0 {
            n = n / 2
        } else {
            n = n*3 + 1
        }

        // If the new number is odd, append it to the list
        if n%2 != 0 {
            oddCollatz = append(oddCollatz, n)
        }
    }

    // Sort the result in increasing order
    sort.Slice(oddCollatz, func(i, j int) bool {
        return oddCollatz[i] < oddCollatz[j]
    })

    return oddCollatz
}