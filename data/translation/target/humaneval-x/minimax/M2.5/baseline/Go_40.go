package main

import "fmt"

func TriplesSumToZero(l []int) bool {
	n := len(l)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if l[i]+l[j]+l[k] == 0 {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	// Test cases
	fmt.Println(TriplesSumToZero([]int{1, 3, 5, 0}))
	fmt.Println(TriplesSumToZero([]int{1, 3, -2, 1}))
	fmt.Println(TriplesSumToZero([]int{1, 2, 3, 7}))
	fmt.Println(TriplesSumToZero([]int{2, 4, -5, 3, 9, 7}))
	fmt.Println(TriplesSumToZero([]int{1}))
}