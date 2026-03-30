package main

import "fmt"

func StartsOneEnds(n int) int {
	if n == 1 {
		return 1
	}
	// Calculate 10^(n-2)
	pow := 1
	for i := 0; i < n-2; i++ {
		pow *= 10
	}
	return 18 * pow
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(StartsOneEnds(n))
}
