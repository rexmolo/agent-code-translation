package main

import (
    "fmt"
    "sort"
)

func SortedListSum(lst []string) []string {
    // Filter out strings with odd lengths
    var new_lst []string
    for _, s := range lst {
        if len(s)%2 == 0 {
            new_lst = append(new_lst, s)
        }
    }

    // Sort by length first, then alphabetically for ties
    sort.Slice(new_lst, func(i, j int) bool {
        if len(new_lst[i]) != len(new_lst[j]) {
            return len(new_lst[i]) < len(new_lst[j])
        }
        return new_lst[i] < new_lst[j]
    })

    return new_lst
}

func main() {
    // Test cases from the examples
    fmt.Println(SortedListSum([]string{"aa", "a", "aaa"}))           // Output: [aa]
    fmt.Println(SortedListSum([]string{"ab", "a", "aaa", "cd"})) // Output: [ab cd]
}