package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strings"
)

func IsBored(S string) int {
    // Split by '.', '?', or '!' followed by optional whitespace
    re := regexp.MustCompile(`[.?!]\s*`)
    sentences := re.Split(S, -1)

    count := 0
    for _, sentence := range sentences {
        if strings.HasPrefix(sentence, "I ") {
            count++
        }
    }
    return count
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        fmt.Println(IsBored(scanner.Text()))
    }
}
