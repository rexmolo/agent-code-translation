package main

import (
	"strconv"
)

func DecimalToBinary(decimal int) string {
	return "db" + strconv.FormatInt(int64(decimal), 2) + "db"
}