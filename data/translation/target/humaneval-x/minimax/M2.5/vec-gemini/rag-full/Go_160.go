package main

import (
	"fmt"
	"strconv"
)

func DoAlgebra(operator []string, operand []int) int {
	// Convert operands to strings for building expression
	// Then evaluate respecting operator precedence
	// Operators: ** (highest), * // (middle), + - (lowest)
	
	// First, handle exponentiation (right to left in Python, but we'll evaluate left to right for simplicity)
	// Actually need to handle precedence properly
	
	// Build initial values array
	result := operand
	
	// First pass: handle exponentiation (**)
	var newResult []int
	var newOps []string
	
	i := 0
	for i < len(operator) {
		if operator[i] == "**" {
			// Apply exponentiation to last element and current operand
			lastIdx := len(newResult) - 1
			base := newResult[lastIdx]
			exp := result[i+1]
			// Simple power calculation
			p := 1
			for j := 0; j < exp; j++ {
				p *= base
			}
			newResult[lastIdx] = p
			i++
		} else {
			newOps = append(newOps, operator[i])
			newResult = append(newResult, result[i+1])
			i++
		}
	}
	
	// Second pass: handle * and //
	finalResult := []int{newResult[0]}
	finalOps := []string{}
	
	for j := 0; j < len(newOps); j++ {
		if newOps[j] == "*" {
			lastIdx := len(finalResult) - 1
			finalResult[lastIdx] = finalResult[lastIdx] * newResult[j+1]
		} else if newOps[j] == "//" {
			lastIdx := len(finalResult) - 1
			finalResult[lastIdx] = finalResult[lastIdx] / newResult[j+1]
		} else {
			finalOps = append(finalOps, newOps[j])
			finalResult = append(finalResult, newResult[j+1])
		}
	}
	
	// Third pass: handle + and -
	answer := finalResult[0]
	for k := 0; k < len(finalOps); k++ {
		if finalOps[k] == "+" {
			answer = answer + finalResult[k+1]
		} else if finalOps[k] == "-" {
			answer = answer - finalResult[k+1]
		}
	}
	
	return answer
}

// Alternative simpler version: left-to-right evaluation
func DoAlgebraSimple(operator []string, operand []int) int {
	result := operand[0]
	for i, op := range operator {
		switch op {
		case "+":
			result = result + operand[i+1]
		case "-":
			result = result - operand[i+1]
		case "*":
			result = result * operand[i+1]
		case "//":
			result = result / operand[i+1]
		case "**":
			// Simple power
			p := 1
			for j := 0; j < operand[i+1]; j++ {
				p *= result
			}
			result = p
		}
	}
	return result
}

func main() {
	// Test with example: operator=['+', '*', '-'], operand=[2, 3, 4, 5]
	// Expected: 2 + 3 * 4 - 5 = 2 + 12 - 5 = 9
	operator := []string{"+", "*", "-"}
	operand := []int{2, 3, 4, 5}
	
	result := DoAlgebra(operator, operand)
	fmt.Println(result) // Should print 9
}
