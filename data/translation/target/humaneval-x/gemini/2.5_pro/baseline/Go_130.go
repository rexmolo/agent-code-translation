package main

import (
	"fmt"
	"os"
	"strconv"
)

// Tri calculates the first n + 1 terms of a special sequence based on the provided Python code.
// The sequence is defined as:
// tri(0) = 1
// tri(1) = 3
// tri(n) = 1 + n / 2, if n is even.
// tri(n) = tri(n - 1) + tri(n - 2) + (n + 3) / 2, if n is odd.
func Tri(n int) []float64 {
	if n == 0 {
		return []float64{1.0}
	}

	// The Python code implicitly assumes n >= 0. For n=1, it returns [1, 3].
	// We pre-allocate the slice of size n+1 for efficiency.
	myTri := make([]float64, n+1)

	myTri[0] = 1.0
	myTri[1] = 3.0

	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			// Even case: Use float64 division to match Python's behavior.
			myTri[i] = float64(i)/2.0 + 1.0
		} else {
			// Odd case: Follows the logic from the Python code.
			myTri[i] = myTri[i-1] + myTri[i-2] + (float64(i)+3.0)/2.0
		}
	}

	return myTri
}

// main function to make the code runnable and test the Tri function.
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run <filename>.go <n>")
		fmt.Println("Please provide a non-negative integer 'n'.")
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 0 {
		fmt.Println("Error: Please provide a valid non-negative integer.")
		os.Exit(1)
	}

	result := Tri(n)
	fmt.Printf("%v\n", result)
}
