package main

import (
	"fmt"
	"os"
	"strconv"
)

// Fib returns the n-th Fibonacci number.
func Fib(n int) int {
	if n <= 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <number>")
		// Default to a known example if no argument is provided.
		fmt.Println("\nRunning example with n=10:")
		n := 10
		fmt.Printf("Fib(%d) = %d\n", n, Fib(n))
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Invalid number provided '%s'\n", os.Args[1])
		os.Exit(1)
	}
	fmt.Println(Fib(n))
}