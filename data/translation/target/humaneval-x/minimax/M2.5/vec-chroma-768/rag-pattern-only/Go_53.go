package main

import "fmt"

func Add(x int, y int) int {
	return x + y
}

func main() {
	result := Add(2, 3)
	fmt.Println("add(2, 3) =", result)

	result = Add(5, 7)
	fmt.Println("add(5, 7) =", result)
}
