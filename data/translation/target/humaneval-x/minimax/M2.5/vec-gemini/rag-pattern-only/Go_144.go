package main

import (
    "fmt"
    "strconv"
    "strings"
)

func Simplify(x, n string) bool {
    // Split the first fraction x
    partsX := strings.Split(x, "/")
    a := partsX[0]
    b := partsX[1]

    // Split the second fraction n
    partsN := strings.Split(n, "/")
    c := partsN[0]
    d := partsN[1]

    // Convert to integers
    aNum, _ := strconv.Atoi(a)
    bNum, _ := strconv.Atoi(b)
    cNum, _ := strconv.Atoi(c)
    dNum, _ := strconv.Atoi(d)

    // Multiply numerators: a * c
    numerator := aNum * cNum
    // Multiply denominators: b * d
    denom := bNum * dNum

    // Check if the result is a whole number by checking if numerator is divisible by denominator
    return numerator%denom == 0
}

func main() {
    // Test cases
    fmt.Println(Simplify("1/5", "5/1"))   // true (1/5 * 5/1 = 5/5 = 1)
    fmt.Println(Simplify("1/6", "2/1"))   // false (1/6 * 2/1 = 2/6 = 1/3)
    fmt.Println(Simplify("7/10", "10/2")) // false (7/10 * 10/2 = 70/20 = 7/2 = 3.5)
}
