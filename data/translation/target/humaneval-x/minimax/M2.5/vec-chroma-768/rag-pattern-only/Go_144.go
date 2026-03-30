package main

import (
    "strconv"
    "strings"
)

func Simplify(x, n string) bool {
    // Parse first fraction x
    parts1 := strings.Split(x, "/")
    a, _ := strconv.Atoi(parts1[0])
    b, _ := strconv.Atoi(parts1[1])

    // Parse second fraction n
    parts2 := strings.Split(n, "/")
    c, _ := strconv.Atoi(parts2[0])
    d, _ := strconv.Atoi(parts2[1])

    // Multiply fractions: (a/b) * (c/d) = (a*c) / (b*d)
    numerator := a * c
    denominator := b * d

    // Check if the result is a whole number (no remainder)
    return numerator%denominator == 0
}