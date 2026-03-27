package main

import (
    "math"
    "strconv"
    "strings"
)

func ClosestInteger(value string) int {
    // Handle empty string case
    if len(value) == 0 {
        return 0
    }

    // Remove trailing zeros after decimal point
    if strings.Contains(value, ".") {
        value = strings.TrimRight(value, "0")
        // Also remove trailing decimal point if it becomes just "."
        value = strings.TrimRight(value, ".")
    }

    num, err := strconv.ParseFloat(value, 64)
    if err != nil {
        return 0
    }

    // Check if value ends with ".5" - round away from zero
    if strings.HasSuffix(value, ".5") {
        if num > 0 {
            return int(math.Ceil(num))
        } else {
            return int(math.Floor(num))
        }
    }

    // For other values, use standard rounding
    return int(math.Round(num))
}

func main() {
    // Test cases can be added here if needed
}