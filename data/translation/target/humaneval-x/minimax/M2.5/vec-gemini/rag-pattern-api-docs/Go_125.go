package main

import (
    "strings"
)

func SplitWords(txt string) interface{} {
    if strings.Contains(txt, " ") {
        return strings.Split(txt, " ")
    } else if strings.Contains(txt, ",") {
        return strings.Split(strings.ReplaceAll(txt, ",", " "), " ")
    } else {
        count := 0
        for _, r := range txt {
            if r >= 'a' && r <= 'z' && (int(r)-int('a'))%2 == 0 {
                count++
            }
        }
        return count
    }
}

func main() {
    // Test cases
    result1 := SplitWords("Hello world!")
    println(result1)
    
    result2 := SplitWords("Hello,world!")
    println(result2)
    
    result3 := SplitWords("abcdef")
    println(result3)
}