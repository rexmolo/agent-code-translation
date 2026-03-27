package main

import (
	"fmt"
	"sort"
)

func GetOddCollatz(n int) []int {
	var oddCollatz []int

	if n%2 != 0 {
		oddCollatz = []int{n}
	} else {
		oddCollatz = []int{}
	}

	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}

		if n%2 != 0 {
			oddCollatz = append(oddCollatz, n)
		}
	}

	sort.Ints(oddCollatz)
	return oddCollatz
}

func main() {
	fmt.Println(GetOddCollatz(5))
}
