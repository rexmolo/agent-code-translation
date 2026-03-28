package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func Longest(strings []string) interface{} {
    if len(strings) == 0 {
        return nil
    }

    // Find the maximum length
    maxLen := 0
    for _, s := range strings {
        if len(s) > maxLen {
            maxLen = len(s)
        }
    }

    // Return the first string with the maximum length
    for _, s := range strings {
        if len(s) == maxLen {
            return s
        }
    }

    return nil
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter strings (comma-separated): ")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)

    if input == "" {
        // Test case: empty list
        result := Longest([]string{})
        if result == nil {
            fmt.Println("Result: nil")
        } else {
            fmt.Printf("Result: %s\n", result)
        }
        return
    }

    strs := strings.Split(input, ",")
    // Trim whitespace from each string
    for i := range strs {
        strs[i] = strings.TrimSpace(strs[i])
    }

    result := Longest(strs)
    if result == nil {
        fmt.Println("Result: nil")
    } else {
        fmt.Printf("Result: %s\n", result)
    }
}
