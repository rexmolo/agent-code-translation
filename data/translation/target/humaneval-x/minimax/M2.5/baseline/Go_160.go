package main

import (
	"errors"
	"strconv"
	"strings"
)

func DoAlgebra(operator []string, operand []int) int {
	// Build the expression string
	expression := strconv.Itoa(operand[0])
	for i, oprt := range operator {
		expression += oprt + strconv.Itoa(operand[i+1])
	}
	
	// Parse and evaluate the expression
	result, _ := evaluate(expression)
	return result
}

// evaluate parses and evaluates a mathematical expression
// Supports: +, -, *, //, **
func evaluate(expr string) (int, error) {
	// Remove whitespace
	expr = strings.ReplaceAll(expr, " ", "")
	
	// First pass: handle addition and subtraction (lowest precedence)
	// We split by + and - but need to be careful with negative numbers
	var terms []string
	var ops []string
	
	// Simple tokenization: split by operators while keeping them
	currentNum := ""
	for i := 0; i < len(expr); i++ {
		char := string(expr[i])
		if char == "+" || char == "-" {
			if currentNum != "" {
				terms = append(terms, currentNum)
				ops = append(ops, char)
				currentNum = ""
			} else if char == "-" && len(terms) == 0 {
				// Handle negative number at start
				currentNum = "-"
			}
		} else {
			currentNum += char
		}
	}
	if currentNum != "" {
		terms = append(terms, currentNum)
	}
	
	// Evaluate + and - from left to right
	if len(terms) == 0 {
		return 0, errors.New("empty expression")
	}
	
	result, _ := evaluateMulDiv(terms[0])
	for i, op := range ops {
		right, _ := evaluateMulDiv(terms[i+1])
		if op == "+" {
			result += right
		} else {
			result -= right
		}
	}
	
	return result, nil
}

// evaluateMulDiv handles *, // (medium precedence)
func evaluateMulDiv(term string) (int, error) {
	// Handle exponentiation first by splitting on **
	if strings.Contains(term, "**") {
		parts := strings.Split(term, "**")
		result, _ := evaluateExp(parts[0])
		for i := 1; i < len(parts); i++ {
			right, _ := evaluateExp(parts[i])
			result = pow(result, right)
		}
		return result, nil
	}
	
	// Split by * and //
	var factors []string
	var mulOps []string
	
	currentFactor := ""
	i := 0
	for i < len(term) {
		// Check for //
		if i+1 < len(term) && term[i:i+2] == "//" {
			if currentFactor != "" {
				factors = append(factors, currentFactor)
				mulOps = append(mulOps, "//")
				currentFactor = ""
			}
			i += 2
			continue
		}
		
		if string(term[i]) == "*" {
			if currentFactor != "" {
				factors = append(factors, currentFactor)
				mulOps = append(mulOps, "*")
				currentFactor = ""
			}
		} else {
			currentFactor += string(term[i])
		}
		i++
	}
	if currentFactor != "" {
		factors = append(factors, currentFactor)
	}
	
	if len(factors) == 0 {
		return 0, errors.New("empty factor")
	}
	
	// Evaluate * and // from left to right
	result, _ := strconv.Atoi(factors[0])
	for j, op := range mulOps {
		right, _ := strconv.Atoi(factors[j+1])
		if op == "*" {
			result *= right
		} else {
			result /= right
		}
	}
	
	return result, nil
}

// evaluateExp handles ** (highest precedence)
func evaluateExp(factor string) (int, error) {
	return strconv.Atoi(factor)
}

// pow handles exponentiation
func pow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

// For testing
func main() {
	// Test case from the example
	operator := []string{"+", "*", "-"}
	operand := []int{2, 3, 4, 5}
	result := DoAlgebra(operator, operand)
	println("Result:", result) // Should be 9
}
