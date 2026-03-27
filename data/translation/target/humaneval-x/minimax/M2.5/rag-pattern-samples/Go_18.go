package main

import "fmt"

func HowManyTimes(str string, substring string) int {
    if len(substring) == 0 {
        return 0
    }

    times := 0
    strRunes := []rune(str)
    subRunes := []rune(substring)
    subLen := len(subRunes)

    for i := 0; i <= len(strRunes)-subLen; i++ {
        if string(strRunes[i:i+subLen]) == substring {
            times++
        }
    }

    return times
}

func main() {
    // Test cases
    fmt.Println(HowManyTimes("", "a"))    // 0
    fmt.Println(HowManyTimes("aaa", "a"))  // 3
    fmt.Println(HowManyTimes("aaaa", "aa")) // 3
}