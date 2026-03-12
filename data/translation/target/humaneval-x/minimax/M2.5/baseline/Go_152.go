package main

import (
    "fmt"
)

func Compare(game, guess []int) []int {
    result := make([]int, len(game))
    for i := range game {
        diff := game[i] - guess[i]
        if diff < 0 {
            diff = -diff
        }
        result[i] = diff
    }
    return result
}

func main() {
    // Test cases
    fmt.Println(Compare([]int{1, 2, 3, 4, 5, 1}, []int{1, 2, 3, 4, 2, -2}))
    fmt.Println(Compare([]int{0, 5, 0, 0, 0, 4}, []int{4, 1, 1, 0, 0, -2}))
}
