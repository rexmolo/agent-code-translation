package main

func GetMaxTriples(n int) int {
    // Create array A where A[i] = i*i - i + 1 for i from 1 to n
    // Using 0-based indexing in Go, we adjust: value = (i+1)^2 - (i+1) + 1
    A := make([]int, n)
    for i := 0; i < n; i++ {
        A[i] = (i+1)*(i+1) - (i + 1) + 1
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
    // Example usage
    result := GetMaxTriples(5)
    println(result) // Expected output: 1
}