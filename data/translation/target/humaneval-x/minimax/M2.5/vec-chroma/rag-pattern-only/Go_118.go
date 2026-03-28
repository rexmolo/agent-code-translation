package main

import "fmt"

func GetClosestVowel(word string) string {
    if len(word) < 3 {
        return ""
    }

    vowels := map[rune]bool{
        'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
        'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
    }

    for i := len(word) - 2; i > 0; i-- {
        ch := rune(word[i])
        if vowels[ch] {
            left := rune(word[i-1])
            right := rune(word[i+1])
            if !vowels[left] && !vowels[right] {
                return string(ch)
            }
        }
    }
    return ""
}

func main() {
    fmt.Println(GetClosestVowel("yogurt"))  // Expected: "u"
    fmt.Println(GetClosestVowel("FULL"))    // Expected: "U"
    fmt.Println(GetClosestVowel("quick"))   // Expected: ""
    fmt.Println(GetClosestVowel("ab"))      // Expected: ""
}
