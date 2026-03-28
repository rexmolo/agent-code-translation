package main

import (
	"fmt"
	"strconv"
	"strings"
)

func DoAlgebra(operator []string, operand []int) int {
	// Build expression string
	var sb strings.Builder
	sb.WriteString(strconv.Itoa(operand[0]))
	for i, oprt := range operator {
		sb.WriteString(oprt)
		sb.WriteString(strconv.Itoa(operand[i+1]))
	}
	expr := sb.String()

	// Tokenize the expression
	tokens := tokenize(expr)

	// Evaluate with proper operator precedence
	result := evaluate(tokens)

	return result
}

func tokenize(expr string) []string {
	var tokens []string
	var num strings.Builder

	for i := 0; i < len(expr); i++ {
		ch := expr[i]
		// Check for two-character operators
		if i+1 < len(expr) && (expr[i:i+2] == "**" || expr[i:i+2] == "//") {
			if num.Len() > 0 {
				tokens = append(tokens, num.String())
				num.Reset()
			}
			tokens = append(tokens, expr[i:i+2])
			i++ // skip next character
		} else if ch == '+' || ch == '-' || ch == '*' || ch == '/' {
			if num.Len() > 0 {
				tokens = append(tokens, num.String())
				num.Reset()
			}
			tokens = append(tokens, string(ch))
		} else if ch >= '0' && ch <= '9' {
			num.WriteByte(ch)
		}
	}
	if num.Len() > 0 {
		tokens = append(tokens, num.String())
	}

	return tokens
}

func evaluate(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}

	// Parse first number
	pos := 0
	num, _ := strconv.Atoi(tokens[pos])
	pos++

	// Evaluate expression with proper precedence
	// Pass 1: Handle ** (highest precedence)
	newTokens := []string{strconv.Itoa(num)}
	for pos < len(tokens) {
		op := tokens[pos]
		val := tokens[pos+1]
		pos += 2

		if op == "**" {
			// Process ** immediately with previous value
			lastIdx := len(newTokens) - 1
			left, _ := strconv.Atoi(newTokens[lastIdx])
			right, _ := strconv.Atoi(val)
			power := 1
			for i := 0; i < right; i++ {
				power *= left
			}
			newTokens[lastIdx] = strconv.Itoa(power)
		} else {
			newTokens = append(newTokens, op, val)
		}
	}
	tokens = newTokens

	// Pass 2: Handle * and // (medium precedence)
	pos = 0
	num, _ = strconv.Atoi(tokens[pos])
	pos++
	newTokens = []string{strconv.Itoa(num)}

	for pos < len(tokens) {
		op := tokens[pos]
		val := tokens[pos+1]
		pos += 2

		if op == "*" {
			lastIdx := len(newTokens) - 1
			left, _ := strconv.Atoi(newTokens[lastIdx])
			right, _ := strconv.Atoi(val)
			newTokens[lastIdx] = strconv.Itoa(left * right)
		} else if op == "//" {
			lastIdx := len(newTokens) - 1
			left, _ := strconv.Atoi(newTokens[lastIdx])
			right, _ := strconv.Atoi(val)
			newTokens[lastIdx] = strconv.Itoa(left / right)
		} else {
			newTokens = append(newTokens, op, val)
		}
	}
	tokens = newTokens

	// Pass 3: Handle + and - (lowest precedence, left to right)
	pos = 0
	result, _ := strconv.Atoi(tokens[pos])
	pos++

	for pos < len(tokens) {
		op := tokens[pos]
		val, _ := strconv.Atoi(tokens[pos+1])
		pos += 2

		if op == "+" {
			result += val
		} else if op == "-" {
			result -= val
		}
	}

	return result
}

func main() {
	// Test case from the example
	operator := []string{"+", "*", "-"}
	operand := []int{2, 3, 4, 5}
	result := DoAlgebra(operator, operand)
	fmt.Printf("Result: %d\n", result) // Expected: 9

	// Additional test cases
	// Test exponentiation
	operator2 := []string{"**", "+"}
	operand2 := []int{2, 3, 2}
	result2 := DoAlgebra(operator2, operand2)
	fmt.Printf("2**3+2 = %d (expected: 10)\n", result2)

	// Test floor division
	operator3 := []string{"/", "//", "+"}
	operand3 := []int{10, 2, 3, 1}
	result3 := DoAlgebra(operator3, operand3)
	fmt.Printf("10/2//3+1 = %d (expected: 2)\n", result3)
}