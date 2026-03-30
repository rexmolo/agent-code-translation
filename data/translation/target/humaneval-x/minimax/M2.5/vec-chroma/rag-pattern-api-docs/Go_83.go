package main

import (
	"fmt"
	"math"
)

func StartsOneEnds(n int) int {
	if n == 1 {
		return 1
	}
	// 18 * 10^(n-2)
	return 18 * int(math.Pow10(n-2))
}

func main() {
	fmt.Println(StartsOneEnds(1)) // 1
	fmt.Println(StartsOneEnds(2)) // 18
	fmt.Println(StartsOneEnds(3)) // 180
}
