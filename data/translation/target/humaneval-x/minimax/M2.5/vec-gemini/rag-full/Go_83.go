package main

import "fmt"

func StartsOneEnds(n int) int {
	if n == 1 {
		return 1
	}

	// Calculate 10^(n-2)
	power := 1
	for i := 0; i < n-2; i++ {
		power *= 10
	}

	return 18 * power
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(StartsOneEnds(n))
}