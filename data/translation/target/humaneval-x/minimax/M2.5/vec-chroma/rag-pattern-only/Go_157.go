package main

import "fmt"

func RightAngleTriangle(a, b, c int) bool {
    return a*a == b*b+c*c || b*b == a*a+c*c || c*c == a*a+b*b
}

func main() {
    fmt.Println(RightAngleTriangle(3, 4, 5)) // true
    fmt.Println(RightAngleTriangle(1, 2, 3)) // false
}