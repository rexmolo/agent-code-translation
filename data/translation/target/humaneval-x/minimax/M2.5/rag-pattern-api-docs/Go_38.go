package main

import (
    "strings"
)

func encodeCyclic(s string) string {
    // Calculate number of groups (equivalent to (len(s) + 2) // 3)
    n := len(s)
    numGroups := (n + 2) / 3

    // Split string into groups of 3
    var groups []string
    for i := 0; i < numGroups; i++ {
        start := 3 * i
        end := start + 3
        if end > n {
            end = n
        }
        groups = append(groups, s[start:end])
    }

    // Cycle elements in each group. Unless group has fewer than 3 elements.
    var result []string
    for _, group := range groups {
        if len(group) == 3 {
            // group[1:] + group[0] in Python becomes group[1:] + group[0:1] in Go
            result = append(result, group[1:]+group[0:1])
        } else {
            result = append(result, group)
        }
    }

    return strings.Join(result, "")
}

func DecodeCyclic(s string) string {
    // Decoding is applying encode twice
    return encodeCyclic(encodeCyclic(s))
}
