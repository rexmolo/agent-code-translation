package main

import "fmt"

func SumToN(n int) int {
	return n * (n + 1) / 2
}

func main() {
	fmt.Println(SumToN(30))
	fmt.Println(SumToN(100))
	fmt.Println(SumToN(5))
	fmt.Println(SumToN(10))
	fmt.Println(SumToN(1))
}
