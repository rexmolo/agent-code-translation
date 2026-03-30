package main

import "fmt"

func IsNested(s string) bool {
    var opening []int
    var closing []int
    
    // Collect indices of opening and closing brackets
    for i := 0; i < len(s); i++ {
        if s[i] == '[' {
            opening = append(opening, i)
        } else {
            closing = append(closing, i)
        }
    }
    
    // Reverse closing bracket indices to match from outside in
    for i, j := 0, len(closing)-1; i < j; i, j = i+1, j-1 {
        closing[i], closing[j] = closing[j], closing[i]
    }
    
    cnt := 0
    j := 0
    for _, idx := range opening {
        if j < len(closing) && idx < closing[j] {
            cnt++
            j++
        }
    }
    
    return cnt >= 2
}

func main() {
    testCases := []string{
        "[[]]",
        "[]]]]]]][[[[[]",
        "[][]",
        "[]",
        "[[][]]",
        "[[]][[",
    }
    
    for _, tc := range testCases {
        fmt.Printf("IsNested('%s') = %v\n", tc, IsNested(tc))
    }
}
