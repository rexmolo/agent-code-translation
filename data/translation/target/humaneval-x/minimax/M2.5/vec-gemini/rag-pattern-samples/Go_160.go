package main

import (
	"fmt"
	"strconv"
)

func DoAlgebra(operator []string, operand []int) int {
	// Build the expression string
	expression := fmt.Sprintf("%d", operand[0])
	for i, oprt := range operator {
		expression += oprt + fmt.Sprintf("%d", operand[i+1])
	}

	// Tokenize and evaluate the expression with proper operator precedence
	tokens := tokenize(expression)
	rpn := shuntingYard(tokens)
	return evaluateRPN(rpn)
}

// tokenize splits the expression into numbers and operators
func tokenize(expr string) []string {
	var tokens []string
	var num string

	for i := 0; i < len(expr); i++ {
		c := string(expr[i])
		if c >= "0" && c <= "9" {
			num += c
		} else {
			if num != "" {
				tokens = append(tokens, num)
				num = ""
			}
			// Handle multi-character operators (** or //)
			if i+1 < len(expr) && (c == "*" || c == "/") && string(expr[i+1]) == c {
				tokens = append(tokens, c+c)
				i++
			} else {
				tokens = append(tokens, c)
			}
		}
	}
	if num != "" {
		tokens = append(tokens, num)
	}
	return tokens
}

// precedence returns the precedence level of an operator
func precedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "//":
		return 2
	case "**":
		return 3
	default:
		return 0
	}
}

// isRightAssociative returns true for right-associative operators
func isRightAssociative(op string) bool {
	return op == "**"
}

// shuntingYard converts infix notation to Reverse Polish Notation (RPN)
func shuntingYard(tokens []string) []string {
	var output []string
	var operators []string

	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" || token == "//" || token == "**" {
			for len(operators) > 0 {
				lastOp := operators[len(operators)-1]
				if precedence(lastOp) > precedence(token) ||
					(precedence(lastOp) == precedence(token) && !isRightAssociative(token)) {
					output = append(output, operators[len(operators)-1])
					operators = operators[:len(operators)-1]
				} else {
					break
				}
			}
			operators = append(operators, token)
		} else {
			output = append(output, token)
		}
	}

	for len(operators) > 0 {
		output = append(output, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}

	return output
}

// evaluateRPN evaluates a RPN expression
func evaluateRPN(rpn []string) int {
	var stack []int

	for _, token := range rpn {
		if token == "+" || token == "-" || token == "*" || token == "/" || token == "//" || token == "**" {
			// Pop two operands
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result int
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/", "//":
				result = a / b // integer division in Go
			case "**":
				result = pow(a, b)
			}
			stack = append(stack, result)
		} else {
			n, _ := strconv.Atoi(token)
			stack = append(stack, n)
		}
	}

	return stack[0]
}

// pow handles exponentiation (including 0^0 which we define as 1)
func pow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

// For testing - uncomment main to test manually
// func main() {
// 	operator := []string{"+", "*", "-"}
// 	operand := []int{2, 3, 4, 5}
// 	result := DoAlgebra(operator, operand)
// 	fmt.Println(result) // Output: 9
// }
