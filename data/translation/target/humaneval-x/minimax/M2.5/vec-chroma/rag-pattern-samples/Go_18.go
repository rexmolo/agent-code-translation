package main

import "fmt"

func HowManyTimes(str string, substring string) int {
    times := 0
    s := []rune(str)
    sub := []rune(substring)
    subLen := len(sub)

    // Edge case: if substring is longer than string, return 0
    if subLen > len(s) {
        return 0
    }

    for i := 0; i <= len(s)-subLen; i++ {
        // Extract substring at position i and compare
        if string(s[i:i+subLen]) == string(sub) {
            times++
        }
    }

    return times
}

func main() {
    // Test cases from the docstring
    fmt.Println(HowManyTimes("", "a"))   // 0
    fmt.Println(HowManyTimes("aaa", "a")) // 3
    fmt.Println(HowManyTimes("aaaa", "aa")) // 3
}