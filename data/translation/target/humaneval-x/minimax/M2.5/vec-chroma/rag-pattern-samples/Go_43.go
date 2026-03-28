package main

import "fmt"

func PairsSumToZero(l []int) bool {
	for i := 0; i < len(l); i++ {
		for j := i + 1; j < len(l); j++ {
			if l[i]+l[j] == 0 {
				return true
			}
		}
	}
	return false
}

func main() {
	// Test cases
	fmt.Println(PairsSumToZero([]int{1, 3, 5, 0}))          // false
	fmt.Println(PairsSumToZero([]int{1, 3, -2, 1}))        // false
	fmt.Println(PairsSumToZero([]int{1, 2, 3, 7}))          // false
	fmt.Println(PairsSumToZero([]int{2, 4, -5, 3, 5, 7}))  // true
	fmt.Println(PairsSumToZero([]int{1}))                   // false
}
