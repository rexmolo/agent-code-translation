package main

import "fmt"

func CarRaceCollision(n int) int {
	return n * n
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(CarRaceCollision(n))
}