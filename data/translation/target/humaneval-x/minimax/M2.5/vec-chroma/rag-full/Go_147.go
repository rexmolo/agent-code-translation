package main

// GetMaxTriples returns the number of triples (a[i], a[j], a[k]) where i < j < k
// and a[i] + a[j] + a[k] is a multiple of 3.
// a[i] = i*i - i + 1 for i from 1 to n.
func GetMaxTriples(n int) int {
	// Create array A where A[i] = i*i - i + 1 for i from 1 to n
	A := make([]int, n)
	for i := 0; i < n; i++ {
		val := i + 1
		A[i] = val*val - val + 1
	}

	// Count triples where sum is divisible by 3
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