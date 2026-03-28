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
	return Fibfib(n-1) + Fibfib(n-2) + Fibfib(n-3)
}

func main() {
	// Test cases
	fmt.Println(Fibfib(1)) // 0
	fmt.Println(Fibfib(5)) // 4
	fmt.Println(Fibfib(8)) // 24
}
