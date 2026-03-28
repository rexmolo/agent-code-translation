package main

import (
    "fmt"
    "sort"
)

func SortArray(array []int) []int {
    // Handle empty array case
    if len(array) == 0 {
        return []int{}
    }
    
    // Create a copy to not modify the original array
    result := make([]int, len(array))
    copy(result, array)
    
    // Determine sort direction based on sum of first and last elements
    // If sum is even: descending order
    // If sum is odd: ascending order
    descending := (array[0]+array[-1])%2 == 0
    
    if descending {
        sort.Sort(sort.Reverse(sort.IntSlice(result)))
    } else {
        sort.Ints(result)
    }
    
    return result
}

func main() {
    // Test cases from the examples
    fmt.Println(SortArray([]int{}))                    // []
    fmt.Println(SortArray([]int{5}))                  // [5]
    fmt.Println(SortArray([]int{2, 4, 3, 0, 1, 5}))   // [0 1 2 3 4 5]
    fmt.Println(SortArray([]int{2, 4, 3, 0, 1, 5, 6})) // [6 5 4 3 2 1 0]
}
