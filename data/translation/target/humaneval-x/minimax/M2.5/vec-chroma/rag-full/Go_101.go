package main

import (
    "strings"
)

func WordsString(s string) []string {
    if s == "" {
        return []string{}
    }
    
    // Replace commas with spaces
    s = strings.ReplaceAll(s, ",", " ")
    
    // Split by whitespace (handles multiple spaces and removes empty strings)
    return strings.Fields(s)
}
