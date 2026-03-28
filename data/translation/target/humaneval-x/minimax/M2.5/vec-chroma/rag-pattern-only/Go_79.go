package main

import (
	"fmt"
)

func DecimalToBinary(decimal int) string {
	return "db" + fmt.Sprintf("%b", decimal) + "db"
}

func main() {
	// Example usage
	fmt.Println(DecimalToBinary(15))  // Output: db1111db
	fmt.Println(DecimalToBinary(32))  // Output: db100000db
}
