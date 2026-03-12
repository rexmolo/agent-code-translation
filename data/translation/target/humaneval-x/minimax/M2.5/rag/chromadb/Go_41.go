package main

import "fmt"

func CarRaceCollision(n int) int {
	return n * n
}

func main() {
	// Example usage - the function calculates n^2 collisions
	// n cars going one direction vs n cars going the opposite direction
	// results in n*n possible collision pairs
	fmt.Println(CarRaceCollision(3)) // Output: 9
}
