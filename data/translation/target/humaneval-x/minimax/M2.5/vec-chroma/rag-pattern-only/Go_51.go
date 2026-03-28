package main

import "fmt"

func RemoveVowels(text string) string {
    vowels := map[rune]bool{
        'a': true, 'A': true,
        'e': true, 'E': true,
        'i': true, 'I': true,
        'o': true, 'O': true,
        'u': true, 'U': true,
    }

    var result []rune
    for _, r := range text {
        if !vowels[r] {
            result = append(result, r)
        }
    }
    return string(result)
}

func main() {
    // Test cases from docstring
    fmt.Println(RemoveVowels(""))
    fmt.Println("abcdef\nghijklm")
    fmt.Println(RemoveVowels("abcdef\nghijklm"))
    fmt.Println(RemoveVowels("abcdef"))
    fmt.Println(RemoveVowels("aaaaa"))
    fmt.Println(RemoveVowels("aaBAA"))
    fmt.Println(RemoveVowels("zbcd"))
}