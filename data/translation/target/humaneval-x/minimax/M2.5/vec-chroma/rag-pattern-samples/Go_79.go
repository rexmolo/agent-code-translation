package main

import (
	"fmt"
	"strconv"
)

func DecimalToBinary(decimal int) string {
	if decimal == 0 {
		return "db0db"
	}
	return "db" + strconv.FormatInt(int64(decimal), 2) + "db"
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(DecimalToBinary(n))
}
