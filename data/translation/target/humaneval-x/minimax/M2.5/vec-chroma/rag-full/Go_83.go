package main

import (
	"fmt"
	"math"
)

func StartsOneEnds(n int) int {
	if n == 1 {
		return 1
	}
	return 18 * int(math.Pow10(n-2))
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(StartsOneEnds(n))
}
