package main

import "fmt"

func GetMaxTriples(n int) int {
	A := make([]int, n)
	for i := 1; i <= n; i++ {
		A[i-1] = i*i - i + 1
	}

	count := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if (A[i]+A[j]+A[k])%3 == 0 {
					count++
				}
			}
		}
	}

	return count
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(GetMaxTriples(n))
}
