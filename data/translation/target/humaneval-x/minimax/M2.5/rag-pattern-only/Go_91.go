package main

import (
    "regexp"
    "strings"
)

func IsBored(S string) int {
    re := regexp.MustCompile(`[.?!]\s*`)
    sentences := re.Split(S, -1)

    count := 0
    for _, sentence := range sentences {
        sentence = strings.TrimSpace(sentence)
        if len(sentence) >= 2 && sentence[0:2] == "I " {
            count++
        }
    }
    return count
}

func main() {
    // Test cases
    println(IsBored("Hello world"))
    println(IsBored("The sky is blue. The sun is shining. I love this weather"))
}