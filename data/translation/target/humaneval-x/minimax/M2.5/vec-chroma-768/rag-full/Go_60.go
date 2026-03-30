package main

import "fmt"

func SumToN(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum
}

func main() {
	fmt.Println(SumToN(30))
	fmt.Println(SumToN(100))
	fmt.Println(SumToN(5))
	fmt.Println(SumToN(10))
	fmt.Println(SumToN(1))
}