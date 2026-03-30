package main

// Fibfib computes the n-th element of the FibFib number sequence.
// The sequence is defined as:
//   - fibfib(0) == 0
//   - fibfib(1) == 0
//   - fibfib(2) == 1
//   - fibfib(n) == fibfib(n-1) + fibfib(n-2) + fibfib(n-3)
func Fibfib(n int) int {
    if n == 0 {
        return 0
    }
    if n == 1 {
        return 0
    }
    if n == 2 {
        return 1
    }

    // Use iterative approach with dynamic programming for efficiency
    // a = fibfib(n-3), b = fibfib(n-2), c = fibfib(n-1)
    a, b, c := 0, 0, 1
    for i := 3; i <= n; i++ {
        a, b, c = b, c, a+b+c
    }
    return c
}
