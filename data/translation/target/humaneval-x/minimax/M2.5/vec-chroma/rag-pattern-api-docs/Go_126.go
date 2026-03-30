package main

import "sort"

func IsSorted(lst []int) bool {
    // Count occurrences of each number
    count := make(map[int]int)
    for _, v := range lst {
        count[v]++
    }

    // Check if any number appears more than 2 times
    for _, v := range lst {
        if count[v] > 2 {
            return false
        }
    }

    // Check if the list is sorted in ascending order
    // Using sort.IntsAreSorted for idiomatic Go
    return sort.IntsAreSorted(lst)
}

// Test function to verify the implementation
func main() {
    tests := []struct {
        lst      []int
        expected bool
    }{
        {[]int{5}, true},
        {[]int{1, 2, 3, 4, 5}, true},
        {[]int{1, 3, 2, 4, 5}, false},
        {[]int{1, 2, 3, 4, 5, 6}, true},
        {[]int{1, 2, 3, 4, 5, 6, 7}, true},
        {[]int{1, 3, 2, 4, 5, 6, 7}, false},
        {[]int{1, 2, 2, 3, 3, 4}, true},
        {[]int{1, 2, 2, 2, 3, 4}, false},
    }

    for _, test := range tests {
        result := IsSorted(test.lst)
        if result != test.expected {
            println("FAIL:", result, test.expected)
        } else {
            println("PASS:", result)
        }
    }
}