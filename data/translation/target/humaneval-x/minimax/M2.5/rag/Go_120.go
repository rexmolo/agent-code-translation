package main

import "sort"

func Maximum(arr []int, k int) []int {
    if k == 0 {
        return []int{}
    }

    // Make a copy to avoid modifying the original slice
    copied := make([]int, len(arr))
    copy(copied, arr)

    // Sort in ascending order
    sort.Ints(copied)

    // Return the last k elements (k largest numbers)
    return copied[len(copied)-k:]
}

func main() {
    // Example tests
    arr1 := []int{-3, -4, 5}
    k1 := 3
    result1 := Maximum(arr1, k1)
    // Expected: [-4, -3, 5]

    arr2 := []int{4, -4, 4}
    k2 := 2
    result2 := Maximum(arr2, k2)
    // Expected: [4, 4]

    arr3 := []int{-3, 2, 1, 2, -1, -2, 1}
    k3 := 1
    result3 := Maximum(arr3, k3)
    // Expected: [2]

    // Test k=0
    result4 := Maximum(arr1, 0)
    // Expected: []

    // Print results (for demonstration)
    // Using a simple print to verify correctness
    _ = result1
    _ = result2
    _ = result3
    _ = result4
}