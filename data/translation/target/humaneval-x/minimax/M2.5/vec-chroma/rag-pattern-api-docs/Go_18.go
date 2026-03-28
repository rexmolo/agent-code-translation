package main

import "fmt"

func HowManyTimes(str string, substring string) int {
    times := 0

    for i := 0; i <= len(str)-len(substring); i++ {
        if str[i:i+len(substring)] == substring {
            times++
        }
    }

    return times
}

func main() {
    // Test cases
    fmt.Println(HowManyTimes("", "a"))      // Expected: 0
    fmt.Println(HowManyTimes("aaa", "a"))   // Expected: 3
    fmt.Println(HowManyTimes("aaaa", "aa")) // Expected: 3
}
