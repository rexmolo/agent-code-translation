package main

import (
	"fmt"
	"strconv"
)

func DoAlgebra(operator []string, operand []int) int {
	// Build the expression by combining operators and operands
	expression := strconv.Itoa(operand[0])
	for i, oprt := range operator {
		expression += oprt + strconv.Itoa(operand[i+1])
	}
	
	// Evaluate the expression (no eval() in Go, so implement manually)
	result := evaluate(expression)
	return result
}

func evaluate(expr string) int {
	// Parse the expression into tokens (numbers and operators)
	tokens := parseExpression(expr)
	
	// First pass: handle exponentiation (**) - right to left (right associative)
	tokens = evaluateExponentiation(tokens)
	
	// Second pass: handle multiplication (*) and floor division (//)
	tokens = evaluateMulDiv(tokens)
	
	// Third pass: handle addition (+) and subtraction (-)
	tokens = evaluateAddSub(tokens)
	
	// Return the final result (should be a single number)
	return tokens[0]
}

func parseExpression(expr string) []string {
	var tokens []string
	var currentNum string
	
	for _, ch := range expr {
		if (ch >= '0' && ch <= '9') || ch == '-' {
			// Check if this is a negative number (only at start or after operator)
			if len(tokens) > 0 && len(currentNum) == 0 && ch == '-' {
				tokens = append(tokens, string(ch))
				continue
			}
			currentNum += string(ch)
		} else {
			// It's an operator
			if len(currentNum) > 0 {
				tokens = append(tokens, currentNum)
				currentNum = ""
			}
			tokens = append(tokens, string(ch))
		}
	}
	
	// Add the last number
	if len(currentNum) > 0 {
		tokens = append(tokens, currentNum)
	}
	
	return tokens
}

func evaluateExponentiation(tokens []string) []string {
	var result []string
	i := 0
	
	for i < len(tokens) {
		if i+2 < len(tokens) && tokens[i+1] == "**" {
			base, _ := strconv.Atoi(tokens[i])
			exp, _ := strconv.Atoi(tokens[i+2])
			val := 1
			for j := 0; j < exp; j++ {
				val *= base
			}
			result = append(result, strconv.Itoa(val))
			i += 3
		} else {
			result = append(result, tokens[i])
			i++
		}
	}
	
	return result
}

func evaluateMulDiv(tokens []string) []string {
	var result []string
	i := 0
	
	for i < len(tokens) {
		if i+2 < len(tokens) && (tokens[i+1] == "*" || tokens[i+1] == "//") {
			left, _ := strconv.Atoi(tokens[i])
			right, _ := strconv.Atoi(tokens[i+2])
			var val int
			if tokens[i+1] == "*" {
				val = left * right
			} else {
				val = left / right // floor division in Go for positive ints
			}
			result = append(result, strconv.Itoa(val))
			i += 3
		} else {
			result = append(result, tokens[i])
			i++
		}
	}
	
	return result
}

func evaluateAddSub(tokens []string) []string {
	var result []string
	i := 0
	
	for i < len(tokens) {
		if i+2 < len(tokens) && (tokens[i+1] == "+" || tokens[i+1] == "-") {
			left, _ := strconv.Atoi(tokens[i])
			right, _ := strconv.Atoi(tokens[i+2])
			var val int
			if tokens[i+1] == "+" {
				val = left + right
			} else {
				val = left - right
			}
			result = append(result, strconv.Itoa(val))
			i += 3
		} else {
			result = append(result, tokens[i])
			i++
		}
	}
	
	return result
}

func main() {
	// Test with the example
	operator := []string{"+", "*", "-"}
	operand := []int{2, 3, 4, 5}
	result := DoAlgebra(operator, operand)
	fmt.Println(result) // Expected: 9
}