package main

import (
	"strconv"
)

func DoAlgebra(operator []string, operand []int) int {
	// Build initial values array
	values := make([]interface{}, len(operand))
	for i, v := range operand {
		values[i] = v
	}
	
	// Operators to process
	ops := make([]string, len(operator))
	copy(ops, operator)
	
	// First pass: handle ** (exponentiation) - right to left
	// Find ** from right to left
	for i := len(ops) - 1; i >= 0; i-- {
		if ops[i] == "**" {
			left := values[i].(int)
			right := values[i+1].(int)
			result := 1
			for j := 0; j < right; j++ {
				result *= left
			}
			values[i] = result
			// Remove the used operator and right operand
			copy(ops[i:], ops[i+1:])
			ops = ops[:len(ops)-1]
			copy(values[i+1:], values[i+2:])
			values = values[:len(values)-1]
		}
	}
	
	// Second pass: handle * and // - left to right
	newOps := []string{}
	newValues := []interface{}{values[0]}
	
	for i := 0; i < len(ops); i++ {
		val := values[i+1].(int)
		if ops[i] == "*" {
			last := newValues[len(newValues)-1].(int)
			newValues[len(newValues)-1] = last * val
		} else if ops[i] == "//" {
			last := newValues[len(newValues)-1].(int)
			newValues[len(newValues)-1] = last / val
		} else {
			newOps = append(newOps, ops[i])
			newValues = append(newValues, val)
		}
	}
	
	ops = newOps
	values = newValues
	
	// Third pass: handle + and - - left to right
	result := values[0].(int)
	for i := 0; i < len(ops); i++ {
		val := values[i+1].(int)
		if ops[i] == "+" {
			result += val
		} else if ops[i] == "-" {
			result -= val
		}
	}
	
	return result
}

func main() {
	// Test the function
	operator := []string{"+", "*", "-"}
	operand := []int{2, 3, 4, 5}
	result := DoAlgebra(operator, operand)
	println(result)
}