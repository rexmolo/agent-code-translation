package main

import (
	"fmt"
)

func StartsOneEnds(n int) int {
	if n == 1 {
		return 1
	}
	result := 18
	for i := 0; i < n-2; i++ {
		result *= 10
	}
	return result
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(StartsOneEnds(n))
}
