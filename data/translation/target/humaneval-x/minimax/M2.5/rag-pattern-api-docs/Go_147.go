package main

import "fmt"

func GetMaxTriples(n int) int {
	// Create array A where A[i] = i*i - i + 1 for i = 1 to n
	// Using 1-based indexing in formula as per Python code
	A := make([]int, n)
	for i := 1; i <= n; i++ {
		A[i-1] = i*i - i + 1
	}

	count := 0
	// Find all triples (i, j, k) where i < j < k
	// and (A[i] + A[j] + A[k]) % 3 == 0
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
	// Test case from the example
	n := 5
	result := GetMaxTriples(n)
	fmt.Println(result)
}
