package main

import "fmt"

func FlipCase(str string) string {
    result := make([]rune, len(str))
    for i, c := range str {
        if c >= 'A' && c <= 'Z' {
            result[i] = c + 32 // Convert uppercase to lowercase
        } else if c >= 'a' && c <= 'z' {
            result[i] = c - 32 // Convert lowercase to uppercase
        } else {
            result[i] = c
        }
    }
    return string(result)
}

func main() {
    fmt.Println(FlipCase("Hello"))
}
