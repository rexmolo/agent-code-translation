package main

import (
	"strings"
)

// parseParenGroup calculates the max depth for a single parenthesis group.
// This is equivalent to the inner `parse_paren_group` function in the Python code.
func parseParenGroup(s string) int {
	depth := 0
	maxDepth := 0
	for _, c := range s {
		if c == '(' {
			depth++
			if depth > maxDepth {
				maxDepth = depth
			}
		} else if c == ')' {
			depth--
		}
	}
	return maxDepth
}

// ParseNestedParens translates the Python function parse_nested_parens.
// Input to this function is a string represented multiple groups for nested parentheses separated by spaces.
// For each of the group, output the deepest level of nesting of parentheses.
// E.g. (()()) has maximum two levels of nesting while ((())) has three.
func ParseNestedParens(paren_string string) []int {
	// strings.Fields is the idiomatic Go way to split a string by whitespace,
	// and it correctly handles multiple spaces between groups, which is equivalent to
	// the Python `[... for x in paren_string.split(' ') if x]`.
	groups := strings.Fields(paren_string)

	// Create a slice to hold the results. Pre-allocating capacity is a
	// common Go optimization.
	results := make([]int, 0, len(groups))

	for _, group := range groups {
		// Calculate the max depth for the current group and append it to the results.
		results = append(results, parseParenGroup(group))
	}

	return results
}
