package main

import (
	"fmt"
	"math"
)

func Multiply(a, b int) int {
	return int(math.Abs(float64(a%10))) * int(math.Abs(float64(b%10)))
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(Multiply(a, b))
}
