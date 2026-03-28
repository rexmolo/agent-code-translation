package main

import (
    "strings"
)

func FileNameCheck(fileName string) string {
    // Valid extensions
    validSuffixes := []string{"txt", "exe", "dll"}

    // Split by '.'
    parts := strings.Split(fileName, ".")

    // Must have exactly 2 parts (before and after dot)
    if len(parts) != 2 {
        return "No"
    }

    // Check extension is valid
    validExt := false
    for _, s := range validSuffixes {
        if parts[1] == s {
            validExt = true
            break
        }
    }
    if !validExt {
        return "No"
    }

    // Name before dot must not be empty
    if len(parts[0]) == 0 {
        return "No"
    }

    // First character must be a letter (a-z or A-Z)
    firstChar := parts[0][0]
    if !((firstChar >= 'a' && firstChar <= 'z') || (firstChar >= 'A' && firstChar <= 'Z')) {
        return "No"
    }

    // Count digits in the name (must not exceed 3)
    digitCount := 0
    for _, c := range parts[0] {
        if c >= '0' && c <= '9' {
            digitCount++
        }
    }
    if digitCount > 3 {
        return "No"
    }

    return "Yes"
}

func main() {
    // Test cases
    println(FileNameCheck("example.txt"))   // => Yes
    println(FileNameCheck("1example.dll"))  // => No
}
