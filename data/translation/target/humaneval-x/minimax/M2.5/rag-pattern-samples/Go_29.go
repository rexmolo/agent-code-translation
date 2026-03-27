package main

import "strings"

func FilterByPrefix(strings []string, prefix string) []string {
    var result []string
    for _, s := range strings {
        if strings.HasPrefix(s, prefix) {
            result = append(result, s)
        }
    }
    return result
}