package main

import "fmt"

func DecimalToBinary(decimal int) string {
    return "db" + fmt.Sprintf("%b", decimal) + "db"
}

func main() {
    // Example usage
    fmt.Println(DecimalToBinary(15))   // db1111db
    fmt.Println(DecimalToBinary(32))   // db100000db
}
