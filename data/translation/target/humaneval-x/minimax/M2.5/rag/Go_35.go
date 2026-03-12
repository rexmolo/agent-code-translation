package main

import "fmt"

func MaxElement(l []int) int {
	m := l[0]
	for _, e := range l {
		if e > m {
			m = e
		}
	}
	return m
}

func main() {
	// Test cases
	fmt.Println(MaxElement([]int{1, 2, 3}))                           // 3
	fmt.Println(MaxElement([]int{5, 3, -5, 2, -3, 3, 9, 0, 123, 1, -10})) // 123
}
