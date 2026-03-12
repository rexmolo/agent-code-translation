import (
    "strconv"
    "strings"
)

func Simplify(x, n string) bool {
    // Parse first fraction x
    parts1 := strings.Split(x, "/")
    a, err := strconv.Atoi(parts1[0])
    if err != nil {
        return false
    }
    b, err := strconv.Atoi(parts1[1])
    if err != nil {
        return false
    }

    // Parse second fraction n
    parts2 := strings.Split(n, "/")
    c, err := strconv.Atoi(parts2[0])
    if err != nil {
        return false
    }
    d, err := strconv.Atoi(parts2[1])
    if err != nil {
        return false
    }

    // Multiply fractions: (a/b) * (c/d) = (a*c)/(b*d)
    numerator := a * c
    denom := b * d

    // Check if the result is a whole number (numerator divisible by denominator)
    return numerator%denom == 0
}