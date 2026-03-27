package main

func GetMaxTriples(n int) int {
	// Create array A where A[i] = i*i - i + 1 for i from 1 to n
	A := make([]int, n)
	for i := 1; i <= n; i++ {
		A[i-1] = i*i - i + 1
	}

	// Count triples (i, j, k) where i < j < k and (A[i] + A[j] + A[k]) % 3 == 0
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