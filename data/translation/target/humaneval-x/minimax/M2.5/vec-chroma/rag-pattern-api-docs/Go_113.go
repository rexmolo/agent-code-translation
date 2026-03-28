package main

import (
    "fmt"
)

func OddCount(lst []string) []string {
    res := make([]string, 0, len(lst))
    for _, arr := range lst {
        n := 0
        for _, d := range arr {
            digit := int(d - '0')
            if digit%2 == 1 {
                n++
            }
        }
        res = append(res, fmt.Sprintf("the number of odd elements %dn the str%dnng %d of the %dnput.", n, n, n, n))
    }
    return res
}

func main() {
    // Example usage
    result := OddCount([]string{"1234567"})
    fmt.Println(result)

    result2 := OddCount([]string{"3", "11111111"})
    fmt.Println(result2)
}
