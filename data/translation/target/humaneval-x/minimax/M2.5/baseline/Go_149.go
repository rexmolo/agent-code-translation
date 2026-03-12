package main

import (
    "sort"
)

func SortedListSum(lst []string) []string {
    // First sort alphabetically (in-place)
    sort.Strings(lst)

    // Filter to keep only strings with even length
    var new_lst []string
    for _, s := range lst {
        if len(s)%2 == 0 {
            new_lst = append(new_lst, s)
        }
    }

    // Sort by length, then alphabetically for ties
    sort.Slice(new_lst, func(i, j int) bool {
        if len(new_lst[i]) != len(new_lst[j]) {
            return len(new_lst[i]) < len(new_lst[j])
        }
        return new_lst[i] < new_lst[j]
    })

    return new_lst
}
