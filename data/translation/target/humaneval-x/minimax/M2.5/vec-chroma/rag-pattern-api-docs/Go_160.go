package main

import (
	"fmt"
	"strconv"
	"strings"
)

func DoAlgebra(operator []string, operand []int) int {
	if len(operand) < 2 || len(operator) != len(operand)-1 {
		return 0
	}

	// Build the expression string
	expression := strconv.Itoa(operand[0])
	for i, oprt := range operator {
		expression += oprt + strconv.Itoa(operand[i+1])
	}

	// Evaluate the expression with proper operator precedence
	return evaluate(expression)
}

// evaluate computes the expression respecting operator precedence
func evaluate(expr string) int {
	// First, handle parentheses (highest precedence)
	if strings.Contains(expr, "(") {
		return evaluate(parseParentheses(expr))
	}

	// Handle exponentiation ** (highest precedence)
	for strings.Contains(expr, "**") {
		expr = evaluateOperator(expr, "**")
	}

	// Handle *, // (next precedence)
	for strings.Contains(expr, "*") || strings.Contains(expr, "//") {
		expr = evaluateOperator(expr, "*", "//")
	}

	// Handle +, - (lowest precedence)
	for strings.Contains(expr, "+") || (strings.Contains(expr, "-") && !strings.HasPrefix(expr, "-")) {
		expr = evaluateOperator(expr, "+", "-")
	}

	result, _ := strconv.Atoi(expr)
	return result
}

// parseParentheses handles parentheses in the expression
func parseParentheses(expr string) string {
	for i := 0; i < len(expr); i++ {
		if expr[i] == '(' {
			start := i
			for j := i + 1; j < len(expr); j++ {
				if expr[j] == '(' {
					i = j
				} else if expr[j] == ')' {
					// Found matching closing parenthesis
					innerResult := evaluate(expr[start+1 : j])
					return expr[:start] + strconv.Itoa(innerResult) + expr[j+1:]
				}
			}
		}
	}
	return expr
}

// evaluateOperator evaluates a specific operator at highest precedence left-to-right
func evaluateOperator(expr string, ops ...string) string {
	// Tokenize the expression
	tokens := tokenize(expr)

	for i := 1; i < len(tokens)-1; i += 2 {
		for _, op := range ops {
			if tokens[i] == op {
				// Evaluate this operation
				left, _ := strconv.Atoi(tokens[i-1])
				right, _ := strconv.Atoi(tokens[i+1])
				var result int
				switch op {
				case "**":
					result = 1
					for k := 0; k < right; k++ {
						result *= left
					}
				case "*":
					result = left * right
				case "//":
					result = left / right
				case "+":
					result = left + right
				case "-":
					result = left - right
				}
				tokens[i-1] = strconv.Itoa(result)
				// Remove the evaluated operator and right operand
				tokens = append(tokens[:i], tokens[i+2:]...)
				return evaluateOperator(strings.Join(tokens, ""), ops...)
			}
		}
	}
	return strings.Join(tokens, "")
}

// tokenize splits an expression into numbers and operators
func tokenize(expr string) []string {
	var tokens []string
	var currentNum strings.Builder

	for i := 0; i < len(expr); i++ {
		ch := expr[i]
		if (ch >= '0' && ch <= '9') || ch == '-' {
			if ch == '-' && currentNum.Len() > 0 {
				// Negative number
				currentNum.WriteByte(ch)
			} else if currentNum.Len() == 0 && ch == '-' {
				// Negative number at start
				currentNum.WriteByte(ch)
			} else {
				currentNum.WriteByte(ch)
			}
		} else {
			if currentNum.Len() > 0 {
				tokens = append(tokens, currentNum.String())
				currentNum.Reset()
			}
			if ch != ' ' {
				tokens = append(tokens, string(ch))
			}
		}
	}
	if currentNum.Len() > 0 {
		tokens = append(tokens, currentNum.String())
	}

	return tokens
}

func main() {
	// Test the function
	operator := []string{"+", "*", "-"}
	operand := []int{2, 3, 4, 5}
	result := DoAlgebra(operator, operand)
	fmt.Printf("Result: %d\n", result)
}