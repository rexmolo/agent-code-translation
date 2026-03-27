package main

import "fmt"

func Fibfib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}

	// Use iterative approach for efficiency (O(n) vs exponential time with recursion)
	a, b, c := 0, 0, 1 // fibfib(0), fibfib(1), fibfib(2)
	for i := 3; i <= n; i++ {
		a, b, c = b, c, a+b+c
	}
	return c
}

func main() {
	fmt.Println(Fibfib(1)) // 0
	fmt.Println(Fibfib(5)) // 4
	fmt.Println(Fibfib(8)) // 24
}
