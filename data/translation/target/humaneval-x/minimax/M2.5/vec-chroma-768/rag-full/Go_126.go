package main

func IsSorted(lst []int) bool {
    // Count occurrences of each number using a map
    countDigit := make(map[int]int)
    for _, v := range lst {
        countDigit[v]++
    }
    
    // Check if any number appears more than 2 times
    for _, v := range lst {
        if countDigit[v] > 2 {
            return false
        }
    }
    
    // Check if sorted in ascending (non-decreasing) order
    for i := 1; i < len(lst); i++ {
        if lst[i-1] > lst[i] {
            return false
        }
    }
    
    return true
}
