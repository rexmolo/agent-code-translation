package main

import "fmt"

func StrangeSortList(lst []int) []int {
    res := []int{}
    switchVal := true

    // Make a copy to work with (since we'll be removing elements)
    working := make([]int, len(lst))
    copy(working, lst)

    for len(working) > 0 {
        var val int
        if switchVal {
            val = min(working)
        } else {
            val = max(working)
        }
        res = append(res, val)
        working = removeElement(working, val)
        switchVal = !switchVal
    }
    return res
}

func min(slice []int) int {
    minVal := slice[0]
    for _, v := range slice {
        if v < minVal {
            minVal = v
        }
    }
    return minVal
}

func max(slice []int) int {
    maxVal := slice[0]
    for _, v := range slice {
        if v > maxVal {
            maxVal = v
        }
    }
    return maxVal
}

func removeElement(slice []int, element int) []int {
    for i, v := range slice {
        if v == element {
            return append(slice[:i], slice[i+1:]...)
        }
    }
    return slice
}

func main() {
    // Test cases
    fmt.Println(StrangeSortList([]int{1, 2, 3, 4})) // [1 4 2 3]
    fmt.Println(StrangeSortList([]int{5, 5, 5, 5})) // [5 5 5 5]
    fmt.Println(StrangeSortList([]int{}))           // []
}