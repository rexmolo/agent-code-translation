package main

import "fmt"

func DecimalToBinary(decimal int) string {
    return "db" + fmt.Sprintf("%b", decimal) + "db"
}

func main() {
    // Test cases to verify the function
    fmt.Println(DecimalToBinary(15)) // Output: db1111db
    fmt.Println(DecimalToBinary(32)) // Output: db100000db
}