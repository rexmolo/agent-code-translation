package main

import (
	"fmt"
	"strconv"
)

func DoAlgebra(operator []string, operand []int) int {
	// Build the expression string
	expression := strconv.Itoa(operand[0])
	for i, oprt := range operator {
		expression += oprt + strconv.Itoa(operand[i+1])
	}

	// Parse and evaluate the expression with proper precedence
	fmt.Sscanf(expression, "%d", &result)
	// Actually evaluate properly
	tokens := tokenize(expression)
	rpn := shuntingYard(tokens)
	return evaluateRPN(rpn)
}

var result int

func tokenize(expr string) []string {
	var tokens []string
	var currentNum string

	for i := 0; i < len(expr); i++ {
		ch := expr[i]
		if (ch >= '0' && ch <= '9') {
			currentNum += string(ch)
		} else {
			if currentNum != "" {
				tokens = append(tokens, currentNum)
				currentNum = ""
			}
			if ch != ' ' {
				tokens = append(tokens, string(ch))
			}
		}
	}
	if currentNum != "" {
		tokens = append(tokens, currentNum)
	}

	return tokens
}

func precedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	case "//", "**":
		return 3
	}
	return 0
}

func isRightAssociative(op string) bool {
	return op == "**" || op == "//"
}

func shuntingYard(tokens []string) []string {
	var output []string
	var operators []string

	for _, token := range tokens {
		if isNumber(token) {
			output = append(output, token)
		} else {
			for len(operators) > 0 && precedence(operators[len(operators)-1]) >= precedence(token) {
				if isRightAssociative(token) && precedence(operators[len(operators)-1]) == precedence(token) {
					break
				}
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, token)
		}
	}

	for len(operators) > 0 {
		output = append(output, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}

	return output
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func evaluateRPN(tokens []string) int {
	var stack []int

	for _, token := range tokens {
		if isNumber(token) {
			val, _ := strconv.Atoi(token)
			stack = append(stack, val)
		} else {
			if len(stack) < 2 {
				continue
			}
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
			case "/":
				result = a / b
			case "//":
				result = a / b
			case "**":
				result = 1
				for i := 0; i < b; i++ {
					result *= a
				}
			}
			stack = append(stack, result)
		}
	}

	if len(stack) > 0 {
		return stack[0]
	}
	return 0
}

func main() {
	operator := []string{"+", "*", "-"}
	operand := []int{2, 3, 4, 5}
	result := DoAlgebra(operator, operand)
	fmt.Println(result)
}