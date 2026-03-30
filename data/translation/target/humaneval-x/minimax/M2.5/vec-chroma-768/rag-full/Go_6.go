package main

import (
	"strings"
)

func ParseNestedParens(paren_string string) []int {
	groups := strings.Split(paren_string, " ")
	
	var result []int
	
	for _, group := range groups {
		if group == "" {
			continue
		}
		
		depth := 0
		maxDepth := 0
		for _, c := range group {
			if c == '(' {
				depth++
				if depth > maxDepth {
					maxDepth = depth
				}
			} else {
				depth--
			}
		}
		result = append(result, maxDepth)
	}
	
	return result
}