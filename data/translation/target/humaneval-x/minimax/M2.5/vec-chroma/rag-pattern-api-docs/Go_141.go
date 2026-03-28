package main

import (
    "strings"
    "unicode"
)

func FileNameCheck(file_name string) string {
    validSuffixes := map[string]bool{
        "txt": true,
        "exe": true,
        "dll": true,
    }
    
    parts := strings.Split(file_name, ".")
    
    // There should be exactly one dot (2 parts)
    if len(parts) != 2 {
        return "No"
    }
    
    // The substring after the dot should be one of: txt, exe, dll
    if !validSuffixes[parts[1]] {
        return "No"
    }
    
    // The substring before the dot should not be empty
    if len(parts[0]) == 0 {
        return "No"
    }
    
    // The substring before the dot should start with a letter
    firstChar := rune(parts[0][0])
    if !unicode.IsLetter(firstChar) {
        return "No"
    }
    
    // Count digits in the name (before the dot)
    digitCount := 0
    for _, c := range parts[0] {
        if unicode.IsDigit(rune(c)) {
            digitCount++
        }
    }
    
    // There should not be more than three digits
    if digitCount > 3 {
        return "No"
    }
    
    return "Yes"
}
