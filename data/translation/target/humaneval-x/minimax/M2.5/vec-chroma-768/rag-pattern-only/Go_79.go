package main

import (
	"fmt"
	"strconv"
)

func DecimalToBinary(decimal int) string {
	return "db" + strconv.FormatInt(int64(decimal), 2) + "db"
}

func main() {
	fmt.Println(DecimalToBinary(15))
	fmt.Println(DecimalToBinary(32))
}
