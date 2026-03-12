package main

import "fmt"

// Add adds two numbers x and y
func Add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(Add(2, 3))
	fmt.Println(Add(5, 7))
}
