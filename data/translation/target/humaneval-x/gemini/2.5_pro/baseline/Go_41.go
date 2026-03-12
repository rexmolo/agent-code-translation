package main

import "fmt"

// CarRaceCollision imagines a road that's a perfectly straight infinitely long line.
// n cars are driving left to right; simultaneously, a different set of n cars
// are driving right to left. The two sets of cars start out being very far from
// each other. All cars move in the same speed. Two cars are said to collide
// when a car that's moving left to right hits a car that's moving right to left.
// However, the cars are infinitely sturdy and strong; as a result, they continue moving
// in their trajectory as if they did not collide.
// This function outputs the number of such collisions.
func CarRaceCollision(n int) int {
	// Each of the n cars moving left-to-right will collide with each of the n cars
	// moving right-to-left. Therefore, the total number of collisions is n * n.
	return n * n
}

// main is the entry point of the program and demonstrates CarRaceCollision.
func main() {
	// Example 1
	n1 := 2
	fmt.Printf("For n = %d, collisions: %d\n", n1, CarRaceCollision(n1))

	// Example 2
	n2 := 5
	fmt.Printf("For n = %d, collisions: %d\n", n2, CarRaceCollision(n2))
}
